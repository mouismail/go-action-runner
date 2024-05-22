package cmd

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestStartCommand(t *testing.T) {
	var rootCmd = &cobra.Command{Use: "root"}
	var startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start command for testing",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("source", "s", "", "specify the repositories source, GitLab or BitBucket")
	startCmd.Flags().StringP("project", "p", "", "specify the project id for GitLab")
	startCmd.Flags().StringP("org", "o", "", "specify the GitHub organization")

	tests := []struct {
		name       string
		args       []string
		wantErr    bool
		errMessage string
	}{
		{
			name:       "No flags",
			args:       []string{"start"},
			wantErr:    true,
			errMessage: "required flag(s) \"org\", \"project\", \"source\" not set",
		},
		{
			name:       "All flags provided",
			args:       []string{"start", "--source=GitLab", "--project=123", "--org=testOrg"},
			wantErr:    false,
			errMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := executeCommand(rootCmd, tt.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestStartCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && err.Error() != tt.errMessage {
				t.Errorf("TestStartCommand() error message = %v, wantErrMessage %v", err.Error(), tt.errMessage)
			}
		})
	}
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, err = root.ExecuteC()
	return "", err
}
