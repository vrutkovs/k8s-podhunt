package main

import (
	"fmt"
	"math/rand"
	"time"

	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
  killOptions = 3
  etcdRegExp = ".*etcd.*"
)


func inClusterLogin() (*k8s.Clientset, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	return k8s.NewForConfig(config)
}

func killRandomPod(c *k8s.Clientset) (string, error) {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Find random pod
	pods, err := c.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if pods.Items == nil || len(pods.Items) == 0 {
		return "", fmt.Errorf("No pods fetched")
	}
	randomPod := pods.Items[rand.Intn(len(pods.Items))]

	// Kill kill kill
	if err := c.CoreV1().Pods(randomPod.Namespace).Delete(randomPod.Name, &metav1.DeleteOptions{}); err != nil {
		return "", err
	}
	return fmt.Sprintf("Killed pod %s in namespace %s", randomPod.Name, randomPod.Namespace), nil
}

func killRandomDeployment(c *k8s.Clientset) (string, error) {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Find random pod
	deployments, err := c.AppsV1().Deployments("").List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if deployments.Items == nil || len(deployments.Items) == 0 {
		return "", fmt.Errorf("No deployments fetched")
	}
	randomDeployment := deployments.Items[rand.Intn(len(deployments.Items))]

	// Kill kill kill
	if err := c.AppsV1().Deployments(randomDeployment.Namespace).Delete(randomDeployment.Name, &metav1.DeleteOptions{}); err != nil {
		return "", err
	}
	return fmt.Sprintf("Killed deployment %s in namespace %s", randomDeployment.Name, randomDeployment.Namespace), nil
}

func killRandomStatefulSet(c *k8s.Clientset) (string, error) {
	// Seed random
	rand.Seed(time.Now().Unix())

	// Find random pod
	statefulSets, err := c.AppsV1().StatefulSets("").List(metav1.ListOptions{})
	if err != nil {
		return "", err
	}
	if statefulSets.Items == nil || len(statefulSets.Items) == 0 {
		return "", fmt.Errorf("No Stateful Sets fetched")
	}
	randomStatefulSet := statefulSets.Items[rand.Intn(len(statefulSets.Items))]

	// Kill kill kill
	if err := c.AppsV1().StatefulSets(randomStatefulSet.Namespace).Delete(randomStatefulSet.Name, &metav1.DeleteOptions{}); err != nil {
		return "", err
	}
	return fmt.Sprintf("Killed statefulset %s in namespace %s", randomStatefulSet.Name, randomStatefulSet.Namespace), nil
}
