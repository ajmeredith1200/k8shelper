package main

import (
	"fmt"
	"k8shelper/modules/cluster"
	"k8shelper/modules/logging"
)

func podLogs() {
	err := logging.GetPodLogs("kube-system", "etcd-minikube")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Pod logs retrieved successfully.")
}

func clusterInfo() {
	clusterInfo, err := cluster.GetClusterInfo()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func main() {
	clusterInfo()
}
