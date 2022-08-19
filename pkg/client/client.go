// This package will contain all configuration to talk to the k8s api-server
package client

import (
	"context"
	"fmt"

	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog"
)

type KubernetesAPI struct {
	Client kubernetes.Interface
}

func (k *KubernetesAPI) SelectIngresses() *v1.IngressList {

	klog.V(5).Info("Getting a list of ingresses")
	// fmt.Println("### Getting a List of Ingresses ###")
	result, err := k.Client.NetworkingV1().Ingresses("").List(context.TODO(), metav1.ListOptions{
		// LabelSelector: labelSelector,
	})
	if err != nil {
		panic(fmt.Errorf("failed to get a list of Ingresses: %v", err))
	}

	klog.V(5).Info("Retrieved a list of ingresses")
	return result

}
