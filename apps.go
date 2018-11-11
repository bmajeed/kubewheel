package main

import (
	"k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type App struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func GetApps() (*corev1.NamespaceList, error) {
	clientset, err := getKubeClientset()
	if err != nil {
		return nil, err
	}
	return clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
}

func (a App) Create() error {
	clientset, err := getKubeClientset()
	if err != nil {
		return err
	}
	_, err = clientset.CoreV1().Namespaces().Create(&corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: a.Name,
		},
	})
	return err
}

func (a App) Delete() error {
	clientset, err := getKubeClientset()
	if err != nil {
		return err
	}
	return clientset.CoreV1().Namespaces().Delete(a.Name, &metav1.DeleteOptions{})
}

func (a App) GetAppDeployments() (*v1.DeploymentList, error) {
	clientset, err := getKubeClientset()
	if err != nil {
		return nil, err
	}
	return clientset.AppsV1().Deployments(a.Name).List(metav1.ListOptions{})
}
