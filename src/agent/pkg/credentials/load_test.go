package credentials

import (
	"os"
	"testing"

	"gopkg.in/ini.v1"
)

// Mock the readFromSection function
func readFromSection(cfg *ini.File, section string, key string) (string, error) {
	return "mockValue", nil
}

func TestSet(t *testing.T) {
	// Set up the environment variable
	home := os.Getenv("HOME")
	os.Setenv("HOME", "/mock/home")

	// Create a mock credentials file
	mockCredentialsPath := "/mock/home/.gha/config"
	os.MkdirAll("/mock/home/.gha", os.ModePerm)
	mockFile, err := os.Create(mockCredentialsPath)
	if err != nil {
		t.Fatalf("Failed to create mock credentials file: %v", err)
	}
	defer os.RemoveAll("/mock/home")

	mockFile.WriteString(`
[credentials]
token = mockToken
owner = mockOwner
repo = mockRepo
`)
	mockFile.Close()

	// Create a Credentials instance
	c := &Credentials{}

	// Call the Set function
	err = c.Set()
	if err != nil {
		t.Fatalf("Set() returned an error: %v", err)
	}

	// Verify the values
	if c.token != "mockToken" {
		t.Errorf("Expected token to be 'mockToken', got '%s'", c.token)
	}
	if c.Owner != "mockOwner" {
		t.Errorf("Expected owner to be 'mockOwner', got '%s'", c.Owner)
	}
	if c.Repo != "mockRepo" {
		t.Errorf("Expected repo to be 'mockRepo', got '%s'", c.Repo)
	}

	// Restore the original HOME environment variable
	os.Setenv("HOME", home)
}
