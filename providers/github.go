package providers

import (
	"context"
	"encoding/json"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"kubewheel/apps"
	"net/url"
	"os"
	"strings"
)

func GetGithubClient() *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func GetConfigFromRepo(repo, branch string) (*apps.KubeWheelConfig, error) {
	if branch == "" {
		branch = "master"
	}
	name, owner, _ := parseRepoUrl(repo)
	client := GetGithubClient()
	fileContents, _, _, err := client.Repositories.GetContents(context.Background(), owner, name, "kubewheel.json", &github.RepositoryContentGetOptions{Ref: branch})
	if err != nil {
		panic(err)
	}
	kubewheelConfig := apps.KubeWheelConfig{}
	contents, err := fileContents.GetContent()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(contents), &apps.KubeWheelConfig{})
	if err != nil {
		return nil, err
	}
	kubewheelConfig.Clean()
	return &kubewheelConfig, nil
}

func parseRepoUrl(repoUrl string) (name string, owner string, err error) {
	components, err := url.Parse(repoUrl)
	if err != nil {
		return
	}
	path := strings.Split(strings.Trim(components.Path, "/"), "/")
	name = strings.Replace(path[1], ".git", "", 1)
	owner = path[0]
	return
}
