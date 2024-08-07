package main

import (
	"context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pods, err := clientSet.CoreV1().Pods(v1.NamespaceDefault).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, pod := range pods.Items {
		println(pod.Name)
	}

	deploys, err := clientSet.AppsV1().Deployments("kube-system").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, deploy := range deploys.Items {
		println(deploy.Name)
	}
}
