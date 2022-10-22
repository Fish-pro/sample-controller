package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"

	clientSet "github.com/Fish-pro/sample-controller/pkg/generated/clientset/versioned"
	"github.com/Fish-pro/sample-controller/pkg/generated/informers/externalversions"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		log.Fatalln(err)
	}
	clientset, err := clientSet.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	}
	list, err := clientset.SamplecontrollerV1alpha1().Ats().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
	}
	for _, bar := range list.Items {
		fmt.Println(bar.Name)
	}
	factory := externalversions.NewSharedInformerFactory(clientset, 0)
	factory.Crd().V1().Bars().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    nil,
		UpdateFunc: nil,
		DeleteFunc: nil,
	})
	// todo
}
