package schema

import (
	"bytes"
	"fmt"
	"github.com/iancoleman/strcase"
	"go/format"
	"text/template"
)

const templateText = `package generated  
  
import (  
"encoding/json"  
  
tfjson "github.com/hashicorp/terraform-json"  
)  
  
const {{.ResourceName}} = ` + "`{{.ResourceSchema}}`" + `
  
func {{.ResourceNameCamel}}Schema() *tfjson.Schema {  
	var result tfjson.Schema  
	_ = json.Unmarshal([]byte({{.ResourceName}}), &result)  
	return &result  
}  
`

type TemplateData struct {
	ResourceName      string
	ResourceSchema    string
	ResourceNameCamel string
}

func GenerateGoFileContent(resourceName, resourceSchema string) (string, error) {
	data := TemplateData{
		ResourceName:      strcase.ToLowerCamel(resourceName),
		ResourceSchema:    resourceSchema,
		ResourceNameCamel: strcase.ToCamel(resourceName),
	}

	tmpl, err := template.New("goFileContent").Parse(templateText)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	formattedSource, err := format.Source(buf.Bytes())
	if err != nil {
		return "", fmt.Errorf("error formatting generated Go file content: %w", err)
	}

	return string(formattedSource), nil
}
