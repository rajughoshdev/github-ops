package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/v58/github"
)

func main() {
	orgUsername := "rajughoshdev"

	repoNames := []string{"github-ops"}

	for _, repoName := range repoNames {
		fmt.Printf("Recent Pull Requests for %s/%s in the last 12 hours:\n", orgUsername, repoName)
		fetchRecentPRs(orgUsername, repoName)
		fmt.Println()
	}
}

func fetchRecentPRs(owner, repo string) {
	ctx := context.Background()

	client := github.NewClient(nil)

	// Calculate the time 12 hours ago
	twelveHoursAgo := time.Now().Add(-12 * time.Hour)

	// List pull requests for the repository
	prs, _, err := client.PullRequests.List(ctx, owner, repo, &github.PullRequestListOptions{
		State:     "all", // "all" includes open and closed pull requests
		Sort:      "updated",
		Direction: "desc",
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, pr := range prs {
		prTime := pr.UpdatedAt.Time

		// Check if the PR was updated in the last 12 hours
		if prTime.After(twelveHoursAgo) {
			prURL := fmt.Sprintf("https://github.com/%s/%s/pull/%d", owner, repo, *pr.Number)
			fmt.Printf("- [%s](%s) by %s at %s\n", *pr.Title, prURL, *pr.User.Login, prTime.Format(time.RFC3339))
		}
	}
}
