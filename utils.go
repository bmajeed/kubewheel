package main

import (
	"github.com/atrox/homedir"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"math/rand"
	"path/filepath"
)

func getKubeClientset() (*kubernetes.Clientset, error) {
	config, err := getKubeConfig()

	if err != nil {
		return nil, err
	}

	return kubernetes.NewForConfig(config)
}

func getKubeConfig() (*rest.Config, error) {
	if !gin.IsDebugging() {
		return rest.InClusterConfig()
	}

	homeDir, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	context := ""
	if gin.Mode() != gin.ReleaseMode {
		context = "minikube"
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: filepath.Join(homeDir, ".kube", "config")},
		&clientcmd.ConfigOverrides{
			CurrentContext: context,
		}).ClientConfig()
}


func randomString(n int, onlyLower ...bool) string {
	var letter []rune
	if len(onlyLower) > 0 {
		letter = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	} else {
		letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
