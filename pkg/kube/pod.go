package kube

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"

	pb "github.com/patnaikshekhar/kubernetescitool/interface"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	kciNamespace = "kci"
	gitImage     = "alpine/git"
)

// CreatePod creates a kubernetes job
func CreatePod(clientset *kubernetes.Clientset, request *pb.BuildRequest) (
	string, error) {

	log.Printf("Create Pod - Started %+v", request)

	// Create generic mount
	mounts := []corev1.VolumeMount{
		corev1.VolumeMount{
			MountPath: "/workspace",
			Name:      "main",
		},
	}

	// Add first container to clone repo
	containers := []corev1.Container{
		corev1.Container{
			Image:        gitImage,
			Name:         fmt.Sprintf("step-1"),
			Args:         []string{"clone", request.Repository, "/workspace"},
			VolumeMounts: mounts,
		},
	}

	for index, step := range request.Steps {
		buildStep := corev1.Container{
			Image:        step.Image,
			Name:         fmt.Sprintf("step-%s", strconv.Itoa(index+2)),
			Args:         step.Args,
			VolumeMounts: mounts,
			WorkingDir:   "/workspace",
		}

		if step.Image == "docker" {

			buildStep.VolumeMounts = append(buildStep.VolumeMounts,
				corev1.VolumeMount{
					MountPath: "/var/run",
					Name:      "docker-sock",
				},
			)
		}

		containers = append(containers, buildStep)
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:    kciNamespace,
			GenerateName: "build-",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				corev1.Container{
					Image:        "alpine",
					Name:         "completed",
					Args:         []string{"echo", "completed"},
					WorkingDir:   "/workspace",
					VolumeMounts: mounts,
				},
			},
			InitContainers: containers,
			Volumes: []corev1.Volume{
				corev1.Volume{
					Name: "main",
					VolumeSource: corev1.VolumeSource{
						EmptyDir: &corev1.EmptyDirVolumeSource{},
					},
				},
				corev1.Volume{
					Name: "docker-sock",
					VolumeSource: corev1.VolumeSource{
						HostPath: &corev1.HostPathVolumeSource{
							Path: "/var/run",
						},
					},
				},
			},
			RestartPolicy: "Never",
		},
	}

	newPod, err := clientset.CoreV1().Pods(kciNamespace).Create(pod)
	if err != nil {
		return "", err
	}

	log.Printf("Create Pod - Completed %+v", newPod)

	return newPod.Name, nil
}

// GetLogs streams logs from the build pod when it is deployed
func GetLogs(clientset *kubernetes.Clientset,
	podName string, stream pb.Kci_BuildServer) error {

	// Get a list of steps
	pod, err := clientset.CoreV1().Pods(kciNamespace).Get(podName,
		metav1.GetOptions{})
	if err != nil {
		return err
	}

	noOfSteps := int32(len(pod.Spec.InitContainers))
	currentStep := int32(1)

	for {
		req := clientset.CoreV1().Pods(kciNamespace).GetLogs(podName,
			&corev1.PodLogOptions{
				Follow:    true,
				Container: fmt.Sprintf("step-%d", currentStep),
			},
		)

		readCloser, err := req.Stream()
		if err != nil {
			if strings.Contains(err.Error(), "PodInitializing") ||
				strings.Contains(err.Error(), "is not available") {
				time.Sleep(time.Second * 1)
				continue
			}
			return err
		}

		defer readCloser.Close()
		result, err := ioutil.ReadAll(readCloser)
		if err != nil {
			return err
		}

		err = stream.Send(&pb.BuildResponse{
			Status: "In Progress",
			Data:   string(result),
			Step:   currentStep,
		})
		if err != nil {
			return err
		}

		// Check if container failed
		pod, err := clientset.CoreV1().Pods(kciNamespace).Get(
			podName, metav1.GetOptions{})

		if pod.Status.InitContainerStatuses[currentStep-1].State.Terminated != nil {
			if pod.Status.InitContainerStatuses[currentStep-1].
				State.Terminated.Reason == "Error" {
				return fmt.Errorf("Build failed")
			}
		}

		currentStep++
		if currentStep > noOfSteps {
			break
		}
	}

	return nil
}
