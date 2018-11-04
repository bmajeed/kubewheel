package main

import (
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getApps() (*corev1.NamespaceList, error) {
	clientset, err := getKubeClientset()
	if err != nil {
		return nil, err
	}
	return clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
}

func getApp(name string) (*v1.DeploymentList, error) {
	clientset, err := getKubeClientset()
	if err != nil {
		return nil, err
	}
	return clientset.AppsV1().Deployments(name).List(metav1.ListOptions{})
}