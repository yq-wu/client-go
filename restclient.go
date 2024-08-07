package main

import (
	"context"
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}
	config.GroupVersion = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs
	config.APIPath = "/api"

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// GET request
	pods := v1.PodList{}
	err = restClient.Get().Namespace("kube-system").Resource("pods").Do(context.TODO()).Into(&pods)
	if err != nil {
		panic(err)
	}
	for _, pod := range pods.Items {
		fmt.Println("GET Pod Name:", pod.Name)
	}

	// POST request
	newPod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "new-pod",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "nginx",
					Image: "nginx",
				},
			},
		},
	}
	result := &v1.Pod{}
	err = restClient.Post().Namespace("default").Resource("pods").Body(newPod).Do(context.TODO()).Into(result)
	if err != nil {
		panic(err)
	}
	fmt.Println("POST Pod Name:", result.Name)

	// PUT request
	updatedPod := result.DeepCopy()
	updatedPod.Labels = map[string]string{"updated": "true"}
	err = restClient.Put().Namespace("default").Resource("pods").Name(updatedPod.Name).Body(updatedPod).Do(context.TODO()).Into(result)
	if err != nil {
		panic(err)
	}
	fmt.Println("PUT Pod Name:", result.Name)

	// DELETE request
	err = restClient.Delete().Namespace("default").Resource("pods").Name(result.Name).Do(context.TODO()).Error()
	if err != nil {
		panic(err)
	}
	fmt.Println("DELETE Pod Name:", result.Name)
}
