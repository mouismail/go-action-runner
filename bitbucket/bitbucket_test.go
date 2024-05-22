package bitbucket

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetRepos(t *testing.T) {
	// Setup a mock HTTP server to simulate Bitbucket API responses
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"values":[{"name":"repo1","description":"Test repo 1","clone_url":"http://example.com/repo1","is_private":false},{"name":"repo2","description":"Test repo 2","clone_url":"http://example.com/repo2","is_private":true}]}`))
	}))
	defer server.Close()

	// Create a new Bitbucket client with the mock server URL
	client, err := NewClient(server.URL, "testuser", "testpass")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Execute the GetRepos function
	repos, err := client.GetRepos()
	if err != nil {
		t.Fatalf("GetRepos() error = %v, wantErr %v", err, false)
	}

	// Define the expected result
	want := []*Repo{
		{Name: "repo1", Description: "Test repo 1", CloneURL: "http://example.com/repo1", IsPrivate: false},
		{Name: "repo2", Description: "Test repo 2", CloneURL: "http://example.com/repo2", IsPrivate: true},
	}

	// Compare the actual result with the expected result
	if !reflect.DeepEqual(repos, want) {
		t.Errorf("GetRepos() got = %v, want %v", repos, want)
	}

	// Test error scenario
	server.Close() // Close the server to simulate a network error
	_, err = client.GetRepos()
	if err == nil {
		t.Errorf("GetRepos() error = %v, wantErr %v", err, true)
	}
}
