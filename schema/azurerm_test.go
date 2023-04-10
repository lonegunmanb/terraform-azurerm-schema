package schema

import (
	"golang.org/x/mod/semver"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RefreshAzureRMSchema(t *testing.T) {
	v, err := RefreshAzureRMSchema("../generated")
	require.NoError(t, err)
	assert.NotNil(t, v)
	assert.True(t, semver.IsValid("v"+v.String()))
}

func Test_ExtractAzureRMSchema(t *testing.T) {
	schema, err := ExtractAzureRMProviderSchema()
	require.NoError(t, err)
	assert.Contains(t, schema.ResourceSchemas, "azurerm_kubernetes_cluster")
}

func TestSave(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "save_test")
	if err != nil {
		t.Fatalf("Unable to create temporary directory: %v", err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir) // Clean up after the test
	}()

	// Generate a file path with non-existing subfolders
	filePath := filepath.Join(tmpDir, "non_existing_folder", "test_file.txt")

	// Content to write to the file
	content := []byte("This is a test")

	// Call the save function to save the content to the file
	err = save(filePath, content)
	assert.NoError(t, err, "save function should not return an error")

	// Read the content back from the file and verify if it matches the original content
	readContent, err := os.ReadFile(filePath)
	assert.NoError(t, err, "Reading saved file should not return an error")
	assert.Equal(t, content, readContent, "Content read from the file should match the original content")
	content = []byte("This is another test")
	err = save(filePath, content)
	assert.NoError(t, err, "save function should not return an error")

	// Read the content back from the file and verify if it matches the original content
	readContent, err = os.ReadFile(filePath)
	assert.NoError(t, err, "Reading saved file should not return an error")
	assert.Equal(t, content, readContent, "Content read from the file should match the original content")
}
