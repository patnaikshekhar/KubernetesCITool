package kube

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	kciSecretName = "kci-secret"
)

// CreateSecret creates a secret in the kubernetes cluster
func CreateSecret(

	clientset *kubernetes.Clientset, key string, value string) error {
	// Check if secret exists
	secret, err := clientset.CoreV1().
		Secrets(kciNamespace).Get(kciSecretName, metav1.GetOptions{})
	if err != nil {
		// If not then create
		secret = &corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      kciSecretName,
				Namespace: kciNamespace,
			},
		}
		_, err = clientset.CoreV1().Secrets(kciNamespace).Create(secret)
		if err != nil {
			return err
		}
	}

	// Else get existing secret and update
	secret.StringData = make(map[string]string)
	secret.StringData[key] = value

	_, err = clientset.CoreV1().Secrets(kciNamespace).Update(secret)
	if err != nil {
		return err
	}

	return nil
}
