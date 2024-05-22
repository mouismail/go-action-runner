package cmd

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
)

func TestExecute(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetArgs([]string{"--help"})
	err := rootCmd.Execute()
	if err != nil {
		t.Errorf("Execute() with --help flag failed, %v", err)
	}
	if !bytes.Contains(buf.Bytes(), []byte("Migrator is Repositories Migration Stats Tool.")) {
		t.Errorf("Execute() with --help flag did not print expected help message")
	}
}

func TestExecuteWithInvalidArgs(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetArgs([]string{"--invalid"})
	err := rootCmd.Execute()
	if err == nil {
		t.Errorf("Execute() with invalid flag did not fail")
	}
}
