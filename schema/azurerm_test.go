package schema_test

import (
	"encoding/json"
	"os"
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/azurermSchema/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RefreshAzureRMSchema(t *testing.T) {
	err := schema.RefreshAzureRMSchema()
	require.NoError(t, err)
	s, err := schema.LoadSchema()
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

	fileName := "./tmp.json"
	schema.Save(s, fileName)
	defer func() {
		_ = os.Remove(fileName)
	}()
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		assert.Failf(t, "file does not exist: %s", err.Error())
	}
}

func Test_LoadSchema(t *testing.T) {
	s, err := schema.LoadSchema()
	require.NoError(t, err)
	assert.Contains(t, s.ResourceSchemas, "azurerm_kubernetes_cluster")
}

func Test_Save_OverwriteExistingFile(t *testing.T) {
	// Create a temporary file
	tmpFile, err := os.CreateTemp("", "tmp.json")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Write some data to the file
	_, err = tmpFile.WriteString("existing data")
	require.NoError(t, err)

	// Extract AzureRM provider schema and save it to the file
	s, err := schema.LoadSchema()
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
