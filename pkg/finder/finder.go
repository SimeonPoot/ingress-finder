// This package contains all the configurations to the logic of finding/processing an object.
package finder

import (
	"fmt"
	"k8s-ingress-finder/pkg/client"

	"k8s.io/klog"
)

type Finder struct {
	KClient              client.KubernetesAPI
	IngressLabelSelector string
}

func (f *Finder) Run() {
	klog.Info("Starting finder.Run()")

	result := f.KClient.SelectIngresses()
	for _, v := range result.Items {
		if v.Spec.IngressClassName != nil {
			fmt.Printf("ingress found: %s/%s, class %s\n", v.Namespace, v.Name, *v.Spec.IngressClassName)
		}
		if v.Spec.IngressClassName == nil {
			fmt.Printf("ingress found: %s/%s, WITHOUT class \n", v.Namespace, v.Name)
		}
	}
	// fmt.Println(result)

}
