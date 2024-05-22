package readme

import (
	"testing"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/http"
	"net/http/httptest"
	"os"
)

// TestRead verifies the Read function can correctly read content from a file.
func TestRead(t *testing.T) {
	// Setup a temporary file
	filePath := "./test_readme.md"
	content := "This is a test content."
	os.WriteFile(filePath, []byte(content), 0644)
	defer os.Remove(filePath)

	// Execute the test
	got, err := Read(filePath)
	if err != nil {
		t.Errorf("Read() error = %v, wantErr %v", err, false)
	}
	if got != content {
		t.Errorf("Read() got = %v, want %v", got, content)
	}
}

// TestUpdate verifies the Update function can correctly write content to a file.
func TestUpdate(t *testing.T) {
	// Setup a temporary file
	filePath := "./test_update.md"
	initialContent := "Initial content."
	updatedContent := "Updated content."
	os.WriteFile(filePath, []byte(initialContent), 0644)
	defer os.Remove(filePath)

	// Execute the test
	err := Update(filePath, updatedContent)
	if err != nil {
		t.Errorf("Update() error = %v, wantErr %v", err, false)
	}

	// Verify the file content
	content, _ := os.ReadFile(filePath)
	if string(content) != updatedContent {
		t.Errorf("Update() got = %v, want %v", string(content), updatedContent)
	}
}

// TestUpdateGitHubRepoFile verifies the UpdateGitHubRepoFile function can interact with the GitHub API.
func TestUpdateGitHubRepoFile(t *testing.T) {
	// Setup a mock GitHub server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Setup GitHub client with mock server
	ctx := oauth2.NoContext
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: "fake-token"})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	client.BaseURL = server.URL

	// Execute the test
	fileContents := []byte("Test content for GitHub file update.")
	repo := "test-repo"
	org := "test-org"
	filePath := "test_file.md"
	status := UpdateGitHubRepoFile(client, fileContents, repo, org, filePath)
	if status != "200 OK" {
		t.Errorf("UpdateGitHubRepoFile() status = %v, want %v", status, "200 OK")
	}
}
