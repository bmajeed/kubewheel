package apps

import (
	"context"
	"github.com/atrox/homedir"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
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

func getGithubClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

