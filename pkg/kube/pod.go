package kube

import (
	"fmt"
	"log"
	"strconv"

	pb "github.com/patnaikshekhar/kubernetescitool/interface"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	kciNamespace = "kci"
	gitImage     = "alpine/gitbuild-jsw68"
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
