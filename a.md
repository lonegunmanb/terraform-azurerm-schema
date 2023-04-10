Given the following go main function:

```go
func main() {
err := os.RemoveAll("generated")
if err != nil {
log.Fatalf("Failed to remove 'generated' folder: %v", err)
}

azureVersion, err := schema.RefreshAzureRMSchema()
if err != nil {
log.Fatalf("Failed to refresh AzureRM schema: %v", err)
}

tagExists, err := checkGitHubTag(azureVersion)
if err != nil {
log.Fatalf("Failed to check GitHub tag: %v", err)
}

if !tagExists {
tag := fmt.Sprintf("v%s", azureVersion.String())
commitMsg := fmt.Sprintf("update schema to version %s", azureVersion.String())

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
fmt.Println(obj)

tagRef := plumbing.ReferenceName("refs/tags/" + tag)
_, err = repo.CreateTag(tag, obj.Hash, &git.CreateTagOptions{
Tagger:  &obj.Author,
Message: commitMsg,
})
if err != nil {
log.Fatalf("Failed to create a new tag: %v", err)
}

remoteURL := ""
if runtime.GOOS == "windows" {
remoteURL, err = convertToHttpsUrl(repo)
if err != nil {
log.Fatalf("Failed to convert remote URL to HTTPS: %v", err)
}
}
err = repo.Push(&git.PushOptions{
RemoteName: "origin",
Auth:       &githttp.TokenAuth{Token: os.Getenv("GITHUB_TOKEN")},
RefSpecs:   []config.RefSpec{config.RefSpec(tagRef + ":" + tagRef)},
RemoteURL:  remoteURL,
})
if err != nil {
log.Fatalf("Failed to push tag: %v", err)
}
}
}
```

Sometimes the tag is already existed in the local machine, in that case, recreate the tag then continue to push the tag to the remote repository.

