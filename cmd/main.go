package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	githttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/v51/github"
	"github.com/hashicorp/go-version"
	"github.com/lonegunmanb/azurerm-provider-schema/schema"
	"golang.org/x/oauth2"
)

func main() {
	refreshFlag := flag.Bool("refresh", false, "Refresh AzureRM schema")
	commitFlag := flag.String("commit", "", "Commit schema with specified version")

	// Parse flags
	flag.Parse()

	if *refreshFlag {
		azureVersion := refreshAzureRMSchema()
		fmt.Println(azureVersion.String())
	} else if *commitFlag != "" {
		commitSchema(*commitFlag)
	} else {
		log.Fatal("No valid flag provided. Use -refresh or -commit=<schema_version>")
	}
}

func commitSchema(schemaVersion string) {
	pat := os.Getenv("GITHUB_TOKEN")
	if pat == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}
	var auth transport.AuthMethod = &githttp.BasicAuth{
		Username: "github-actions[bot]",
		Password: pat,
	}

	tag := fmt.Sprintf("v%s", schemaVersion)
	commitMsg := fmt.Sprintf("update schema to version %s", schemaVersion)

	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatalf("Failed to open the git repository: %v", err)
	}

	w, err := repo.Worktree()
	if err != nil {
		log.Fatalf("Failed to get the worktree: %v", err)
	}

	_, err = w.Add(".")
	if err != nil {
		log.Fatalf("Failed to add changes to the staging area: %v", err)
	}

	commit, err := w.Commit(commitMsg, &git.CommitOptions{
		Author: &object.Signature{
			Name:  "github-actions[bot]",
			Email: "github-actions[bot]@users.noreply.github.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatalf("Failed to commit changes: %w", err)
	}
	obj, err := repo.CommitObject(commit)
	if err != nil {
		log.Fatalf("Failed to get the commit object: %w", err)
	}
	remoteURL, err := convertToHttpsUrl(repo)
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth:       auth,
		Progress:   os.Stdout,
		RemoteURL:  remoteURL,
	})
	if err != nil {
		log.Fatalf("Failed to push changes: %v", err)
	}

	tagExists, err := checkGitHubTag(schemaVersion)
	if err != nil {
		log.Fatalf("Failed to check GitHub tag: %v", err)
	}

	if !tagExists {
		// Check if the local tag exists and delete it
		_, err = repo.Tag(tag)
		if err == nil {
			err = deleteLocalTag(repo, tag)
			if err != nil {
				log.Fatalf("Failed to delete existing local tag: %v", err)
			}
		}

		_, err = repo.CreateTag(tag, obj.Hash, &git.CreateTagOptions{
			Tagger:  &obj.Author,
			Message: commitMsg,
		})
		if err != nil {
			log.Fatalf("Failed to create a new tag: %v", err)
		}

		if err != nil {
			log.Fatalf("Failed to convert remote URL to HTTPS: %w", err)
		}
		tagRef := plumbing.ReferenceName("refs/tags/" + tag)
		pushOptions := &git.PushOptions{
			RemoteName: "origin",
			Auth:       auth,
			RefSpecs:   []config.RefSpec{config.RefSpec(tagRef + ":" + tagRef)},
			RemoteURL:  remoteURL,
			Progress:   os.Stdout,
		}
		err = repo.Push(pushOptions)
		if err != nil {
			log.Fatalf("Failed to push tag: %v", err)
		}
	}
}

func refreshAzureRMSchema() *version.Version {
	err := os.RemoveAll("./generated")
	if err != nil {
		log.Fatalf("Failed to remove 'generated' folder: %v", err)
	}
	azureVersion, err := schema.RefreshAzureRMSchema("generated")
	if err != nil {
		log.Fatalf("Failed to refresh AzureRM schema: %v", err)
	}
	return azureVersion
}

func convertToHttpsUrl(repo *git.Repository) (string, error) {
	remote, err := repo.Remote("origin")
	if err != nil {
		return "", err
	}
	remoteURL := remote.Config().URLs[0]

	if strings.HasPrefix(remoteURL, "git@") {
		httpsURL := strings.Replace(remoteURL, ":", "/", 1)
		httpsURL = strings.Replace(httpsURL, "git@", "https://", 1)
		return httpsURL, nil
	}
	return "", fmt.Errorf("remote URL is not in the expected format")
}

func addTokenToRemoteURL(remoteURL, token string) (string, error) {
	u, err := url.Parse(remoteURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse remote URL: %v", err)
	}

	u.User = url.UserPassword("git", token)
	return u.String(), nil
}

func deleteLocalTag(repo *git.Repository, tagName string) error {
	if err := repo.DeleteTag(tagName); err != nil {
		return fmt.Errorf("failed to delete local tag: %w", err)
	}
	return nil
}

func checkGitHubTag(schemaVersion string) (bool, error) {
	var tc *http.Client
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken != "" {
		// Replace "your-access-token" with your personal access token.
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: githubToken},
		)
		tc = oauth2.NewClient(context.Background(), ts)
	}
	client := github.NewClient(tc)
	options := &github.ListOptions{PerPage: 100}

	for {
		tags, resp, err := client.Repositories.ListTags(context.Background(), "lonegunmanb", "azurerm-provider-schema", options)
		if err != nil {
			return false, err
		}
		for _, tag := range tags {
			_, err := version.NewVersion(*tag.Name)
			if err != nil {
				continue
			}
			if *tag.Name == fmt.Sprintf("v%s", schemaVersion) || *tag.Name == schemaVersion {
				return true, nil
			}
		}
		if resp.NextPage == 0 {
			break
		}
		options.Page = resp.NextPage
	}

	return false, nil
}
