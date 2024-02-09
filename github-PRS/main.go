package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/go-github/v58/github"
)

func main() {
	ctx := context.Background()
	client := github.NewClient(nil)
	owner := "owner"
	repo := "repo"
	// calculate the time 12 hours ago
	twelveHoursAgo := time.Now().Add(-12 * time.Hour)
	prs, _, err := client.PullRequests.List(ctx, owner, repo, &github.PullRequestListOptions{
		State:     "all",
		Sort:      "updated",
		Direction: "desc",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pull Requests updated in the last 12 hours:")
	for _, pr := range prs {
		prTime := pr.UpdatedAt.Time
		if prTime.After(twelveHoursAgo) {
			fmt.Printf("- %s (#%d) by %s at %s\n", *pr.Title, *pr.Number, *pr.User.Login, prTime.Format(time.RFC3339))
		}

	}

}
