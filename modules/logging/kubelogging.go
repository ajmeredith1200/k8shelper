package logging

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"sync"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var clientSet *kubernetes.Clientset

var ctx = context.TODO()
var worker = "k8s_Logging" 

func init() {
	config, err := clientcmd.BuildConfigFromFlags("", "C:\\Users\\alexm\\.kube\\config")
	if err != nil {
		log.Fatal(err)
	}

	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPodLogs(namespace string, podName string) error {
	pod, err := clientSet.CoreV1().Pods(namespace).Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("error getting pod: %v", err)
	}

	var wg sync.WaitGroup
	functionList := []func(){}

	for _, container := range append(pod.Spec.InitContainers, pod.Spec.Containers...) {
		podLogOpts := v1.PodLogOptions{
			Follow:    true,
			TailLines: &[]int64{int64(100)}[0],
			Container: container.Name,
		}

		podLogs, err := clientSet.CoreV1().Pods(namespace).GetLogs(podName, &podLogOpts).Stream(ctx)
		if err != nil {
			return fmt.Errorf("error opening log stream: %v", err)
		}
		defer podLogs.Close()

		functionList = append(functionList, func() {
			defer wg.Done()
			reader := bufio.NewScanner(podLogs)
			for reader.Scan() {
				select {
				case <-ctx.Done():
					return
				default:
					line := reader.Text()
					fmt.Println(worker+"/"+podLogOpts.Container, line)
				}
			}
			log.Printf("INFO log EOF %s: %s/%s", reader.Err().Error(), worker, podLogOpts.Container)
		})
	}

	wg.Add(len(functionList))
	for _, f := range functionList {
		go f()
	}
	wg.Wait()

	return nil
}
