package schema

import (
	"encoding/json"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateGoFileContent(t *testing.T) {
	resourceName := "azurerm_kubernetes_cluster"
	resourceSchema := `{"type":"object","properties":{"name":{"type":"string"},"age":{"type":"integer"}}}`

	goFileContent, err := GenerateGoFileContent(resourceName, resourceSchema, PackageResource)
	assert.NoError(t, err, "GenerateGoFileContent should not return an error")

	formattedSource, err := format.Source([]byte(goFileContent))
	assert.NoError(t, err, "Generated Go file content should be formatted correctly")

	expectedSchema := `{"type":"object","properties":{"name":{"type":"string"},"age":{"type":"integer"}}}`
	var schemaMap map[string]interface{}
	if err := json.Unmarshal([]byte(expectedSchema), &schemaMap); err != nil {
		t.Fatal(err.Error())
	}
	// Marshal the schema map back to a well-formatted JSON string
	formattedSchemaBytes, err := json.MarshalIndent(schemaMap, "", "  ")
	require.NoError(t, err)
	formattedSchema := string(formattedSchemaBytes)
	expectedOutput := `package resource  
  
import (  
	"encoding/json"  
	tfjson "github.com/hashicorp/terraform-json"  
)  
  
const azurermKubernetesCluster = ` + fmt.Sprintf("`%s`", formattedSchema) + `  
  
func AzurermKubernetesClusterSchema() *tfjson.Schema {  
	var result tfjson.Schema  
	_ = json.Unmarshal([]byte(azurermKubernetesCluster), &result)  
	return &result  
}  
`

	assert.Equal(t, removeNewLinesAndSpaces(string(formattedSource)), removeNewLinesAndSpaces(expectedOutput), "Generated Go file content should match the expected output")
	// Parse the generated Go file content to check for syntax errors
	fileSet := token.NewFileSet()
	_, err = parser.ParseFile(fileSet, "", goFileContent, parser.AllErrors)
	assert.NoError(t, err, "Generated Go file content should have no syntax errors")
}

func removeNewLinesAndSpaces(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, "\n", ""), " ", "")
}
