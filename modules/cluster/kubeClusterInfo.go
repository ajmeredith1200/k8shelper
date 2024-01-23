package cluster

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClusterInfo() (string, error) {
	config, err := clientcmd.LoadFromFile("C:\\Users\\alexm\\.kube\\config")
	if err != nil {
		return "", err
	}

	configAuthInfo, found := config.AuthInfos[config.CurrentContext]
	if !found {
		return "", fmt.Errorf("authentication information not found for the current context")
	}

	configClusterInfo, found := config.Clusters[config.CurrentContext]
	if !found {
		return "", fmt.Errorf("cluster information not found for the current context")
	}

	authInfoName := configAuthInfo.Username
	if authInfoName == "" {
		authInfoName = configAuthInfo.Name
	}

	clusterInfo := fmt.Sprintf("Cluster: %s, AuthInfo: %s", configClusterInfo.Server, authInfoName)
	return clusterInfo, nil
}

//todo: add a function to get the current context
//todo: add a function to get cluster info
