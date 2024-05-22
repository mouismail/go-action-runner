package gitlab

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetRepos(t *testing.T) {
	// Setup a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"id":1,"name":"repo1","description":"Test repo 1","url":"http://example.com/repo1"},{"id":2,"name":"repo2","description":"Test repo 2","url":"http://example.com/repo2"}]`))
	}))
	defer server.Close()

	tests := []struct {
		name    string
		projectID int
		want    []Repo
		wantErr bool
	}{
		{
			name: "Successfully retrieves repositories",
			projectID: 123,
			want: []Repo{
				{ID: 1, Name: "repo1", Description: "Test repo 1", URL: "http://example.com/repo1"},
				{ID: 2, Name: "repo2", Description: "Test repo 2", URL: "http://example.com/repo2"},
			},
			wantErr: false,
		},
		{
			name:    "Handles error when fetching repositories",
			projectID: 0,
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Override the GitLab API URL with the mock server URL
			oldAPIURL := apiURL
			apiURL = server.URL
			defer func() { apiURL = oldAPIURL }()

			got, err := GetRepos(tt.projectID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRepos() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRepos() got = %v, want %v", got, tt.want)
			}
		})
	}
}
