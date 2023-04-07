package schema_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/gruntwork-io/terratest/modules/random"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/azurermSchema/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RefreshAzureRMSchema(t *testing.T) {
	tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("temp%d.json", random.Random(0, 1000)))
	defer func() {
		_ = os.Remove(tmpFile)
	}()
	err := schema.RefreshAzureRMSchema(tmpFile)
	require.NoError(t, err)
	s, err := schema.LoadSchema(tmpFile)
	require.NoError(t, err)
	assert.Contains(t, s.ResourceSchemas, "azurerm_kubernetes_cluster")
}

func Test_ExtractAzureRMSchema(t *testing.T) {
	schema, err := schema.ExtractAzureRMProviderSchema()
	require.NoError(t, err)
	assert.Contains(t, schema.ResourceSchemas, "azurerm_kubernetes_cluster")
}

func Test_SaveSchema(t *testing.T) {
	s := &schema.ProviderSchema{ProviderSchema: &tfjson.ProviderSchema{}}

	tmpFile := filepath.Join(os.TempDir(), fmt.Sprintf("temp%d.json", random.Random(0, 1000)))
	defer func() {
		_ = os.Remove(tmpFile)
	}()
	err := schema.Save(s, tmpFile)
	require.NoError(t, err)
	defer func() {
		_ = os.Remove(tmpFile)
	}()
	if _, err := os.Stat(tmpFile); os.IsNotExist(err) {
		assert.Failf(t, "file does not exist: %s", err.Error())
	}
}

func Test_Save_OverwriteExistingFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp(os.TempDir(), "tmp.json")
	require.NoError(t, err)
	defer func() {
		_ = os.Remove(tmpFile.Name())
	}()

	// Write some data to the file
	_, err = tmpFile.WriteString("existing data")
	require.NoError(t, err)

	// Extract AzureRM provider schema and save it to the file
	s, err := schema.LoadSchema("./azure.json")
	require.NoError(t, err)
	err = schema.Save(&schema.ProviderSchema{ProviderSchema: s}, tmpFile.Name())
	require.NoError(t, err)

	// Read the file contents and check that they match the expected JSON
	bytes, err := os.ReadFile(tmpFile.Name())
	require.NoError(t, err)
	expectedJSON, err := json.Marshal(schema.ProviderSchema{ProviderSchema: s})
	require.NoError(t, err)
	require.JSONEq(t, string(expectedJSON), string(bytes))
}
