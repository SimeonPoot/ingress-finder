package main

import (
	"flag"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	"k8s-ingress-finder/pkg/client"
	"k8s-ingress-finder/pkg/finder"
)

func main() {
	var (
		kubeconfig string
		namespace  string
		config     *rest.Config
		err        error
	)

	flag.StringVar(&kubeconfig, "kubeconfig", kubeconfig, "kubeconfig file")
	flag.StringVar(&namespace, "namespace", "default", "the namespace where the configMap is configured")
	// klog
	klog.InitFlags(nil)
	flag.Set("v", "3")
	flag.Set("alsologtostderr", "true")
	flag.Parse()
	// klog := klogr.New()

	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
	}

	if kubeconfig != "" {
		// creates the out-cluster config
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		// creates the in-cluster config
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		klog.Error(err, "error creating configuration")
		os.Exit(1)
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	finder := finder.Finder{
		KClient: client.KubernetesAPI{Client: clientset},
	}

	finder.Run()

}
