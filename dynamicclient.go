package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// 使用 DynamicClient.Resource()，指定要操作的资源对象，获取到该资源的 Client
	dynamicResourceClient := dynamicClient.Resource(schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	})

	// 先为该Client指定ns，然后调用 Client 的 Get() 方法，获取到该资源对象
	unstructured, err := dynamicResourceClient.
		Namespace("kube-system").
		Get(context.TODO(), "coredns", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	// 调用 runtime.DefaultUnstructuredConverter.FromUnstructured()，将 unstructured 反序列化成 Deployment 对象
	deploy := &appsv1.Deployment{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructured.UnstructuredContent(), deploy)
	if err != nil {
		panic(err)
	}

	// 打印 deploy 名称和命名空间
	fmt.Printf("deploy.Name: %s\ndeploy.namespace: %s\n", deploy.Name, deploy.Namespace)
}
