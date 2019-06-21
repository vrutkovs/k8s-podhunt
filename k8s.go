package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
  etcdRegExp = ".*etcd.*"
)


func inClusterLogin() (*k8s.Clientset, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// Seed random
	rand.Seed(time.Now().Unix())

	// creates the clientset
	return k8s.NewForConfig(config)
}

func getAvailableNamespaces(c *k8s.Clientset) ([]string, error) {
	nms, err := c.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	namespaces := make([]string, len(nms.Items))
	for _, n := range nms.Items {
		namespaces = append(namespaces, n.Name)
	}
	return namespaces, nil
}

func killRandomPod(c *k8s.Clientset) (string, error) {
	log.Println("Killing random pod")

	namespaces, err := getAvailableNamespaces(c)
	if err != nil {
		return "", err
	}
	randomNamespace := namespaces[rand.Intn(len(namespaces))]

	// Find random pod
	pods, err := c.CoreV1().Pods(randomNamespace).List(metav1.ListOptions{})
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
	message := fmt.Sprintf("Killed pod %s in namespace %s", randomPod.Name, randomNamespace)
	log.Println(message)
	return message, nil
}

func killRandomDeployment(c *k8s.Clientset) (string, error) {
	log.Println("Killing random Deployment")

	namespaces, err := getAvailableNamespaces(c)
	if err != nil {
		return "", err
	}
	randomNamespace := namespaces[rand.Intn(len(namespaces))]

	// Find random pod
	deployments, err := c.AppsV1().Deployments(randomNamespace).List(metav1.ListOptions{})
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
	message := fmt.Sprintf("Killed deployment %s in namespace %s", randomDeployment.Name, randomNamespace)
	log.Println(message)
	return message, nil
}

func killRandomStatefulSet(c *k8s.Clientset) (string, error) {
	log.Println("Killing random StatefulSet")

	namespaces, err := getAvailableNamespaces(c)
	if err != nil {
		return "", err
	}
	randomNamespace := namespaces[rand.Intn(len(namespaces))]

	// Find random pod
	statefulSets, err := c.AppsV1().StatefulSets(randomNamespace).List(metav1.ListOptions{})
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
	message := fmt.Sprintf("Killed statefulset %s in namespace %s", randomStatefulSet.Name, namespaces)
	log.Println(message)
	return message, nil
}
