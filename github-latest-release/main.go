package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v58/github"
	"github.com/joho/godotenv"
)

func getLatestTag(owner, repo string) (string, error) {

	client := github.NewClient(nil)

	ctx := context.Background()

	getRelease, _, err := client.Repositories.GetLatestRelease(ctx, owner, repo)

	if err != nil {
		return "", err
	}
	tagName := getRelease.TagName
	return string(*tagName), nil
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	getRepositories := os.Getenv("GITHUB_REPOSITORIES")
	repositories := strings.Split(getRepositories, ",")

	for _, repo := range repositories {
		splitOwnerRepo := strings.Split(strings.TrimSpace(repo), "/")

		if len(splitOwnerRepo) != 2 {
			log.Printf("Invalid owner/repo formate for %s\n", repo)
			continue
		}
		owner := splitOwnerRepo[0]
		repo := splitOwnerRepo[1]

		latestTag, err := getLatestTag(owner, repo)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The latest release of %s is: %s \n", repo, latestTag)
	}

}
