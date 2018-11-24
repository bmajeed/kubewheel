package apps

import (
	"github.com/atrox/homedir"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"path/filepath"
)

func GetKubeClientset() (*kubernetes.Clientset, error) {
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

	contextName := ""
	if gin.Mode() != gin.ReleaseMode {
		contextName = "minikube"
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: filepath.Join(homeDir, ".kube", "config")},
		&clientcmd.ConfigOverrides{
			CurrentContext: contextName,
		}).ClientConfig()
}
