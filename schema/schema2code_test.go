package schema_test

import (
	"github.com/lonegunmanb/azurermSchema/schema"
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateGoFileContent(t *testing.T) {
	resourceName := "azurerm_kubernetes_cluster"
	resourceSchema := `{"type":"object","properties":{"name":{"type":"string"},"age":{"type":"integer"}}}`

	goFileContent, err := schema.GenerateGoFileContent(resourceName, resourceSchema)
	assert.NoError(t, err, "GenerateGoFileContent should not return an error")

	formattedSource, err := format.Source([]byte(goFileContent))
	assert.NoError(t, err, "Generated Go file content should be formatted correctly")

	expectedOutput := `package generated  
  
import (  
	"encoding/json"  
	tfjson "github.com/hashicorp/terraform-json"  
)  
  
const azurermKubernetesCluster = ` + "`{\"type\":\"object\",\"properties\":{\"name\":{\"type\":\"string\"},\"age\":{\"type\":\"integer\"}}}`" + `  
  
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
