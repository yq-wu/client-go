package main

import (
	"fmt"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 1、先创建一个客户端配置config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err.Error())
	}

	// 2、使用 discovery.NewDiscoveryClientForConfig()，创建一个 DiscoveryClient 对象
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// 3、使用 DiscoveryClient.ServerGroupsAndResources()，获取所有资源列表
	_, resourceLists, err := discoveryClient.ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}

	// 4、遍历资源列表，打印出资源组和资源名称
	for _, resource := range resourceLists {
		fmt.Printf("resource groupVersion: %s\n", resource.GroupVersion)
		for _, resource := range resource.APIResources {
			fmt.Printf("resource name: %s\n", resource.Name)
		}
		fmt.Println("--------------------------")
	}
}
