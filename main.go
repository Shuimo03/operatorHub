package main

import (
	"k8s.io/api/events/v1beta1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
	if err != nil {
		panic(err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	stopCh := make(chan struct{})
	defer close(stopCh)

	sharedInformer := informers.NewSharedInformerFactory(client, time.Minute)
	informer := sharedInformer.Core().V1().Events().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			event := v1beta1.Event{}
			log.Println(event.Name)
		},
	})
	informer.Run(stopCh)
}
