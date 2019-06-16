package main

import (
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
  killOptions = 3
  etcdRegExp = ".*etcd.*"
)


func inClusterLogin() (*k8s.Clientset, error) {
  // creates the in-cluster config
  config, err := rest.InClusterConfig()
  if err != nil {
    panic(err.Error())
  }
  // creates the clientset
  return k8s.NewForConfig(config)
}

func killRandomPod(c *k8s.Clientset) error {
  return nil
}
