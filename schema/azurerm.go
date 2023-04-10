package schema

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/iancoleman/strcase"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
)

const registerTemplate = `package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/{{ .RepoOwner }}/{{ .GoModule }}/generated/data"
	"github.com/{{ .RepoOwner }}/{{ .GoModule }}/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)
	{{- range $name, $expression := .ResourceSchemas }}  
	resources["{{$name}}"] = {{$expression}}  
	{{- end }}  
	{{- range $name, $expression := .DataSourceSchemas }}  
	dataSources["{{$name}}"] = {{$expression}}  
	{{- end }}  
	Resources = resources
	DataSources = dataSources
}`

const registerTestTemplate = `package generated_test

import (
	"testing"

	"github.com/{{ .RepoOwner }}/{{ .GoModule }}/generated"
	"github.com/stretchr/testify/assert"
)

func TestResourceSchema(t *testing.T) {
	assert.NotEmpty(t, generated.Resources)
	assert.NotEmpty(t, generated.DataSources)
}`

type RegisterParameter struct {
	ResourceSchemas   map[string]string
	DataSourceSchemas map[string]string
	RepoOwner         string
	GoModule          string
}

func RefreshAzureRMSchema() error {
	s, err := ExtractAzureRMProviderSchema()

	if err != nil {
		return err
	}
	return SaveProviderSchema(s)
}

func ExtractAzureRMProviderSchema() (*tfjson.ProviderSchema, error) {
	tmpFolder, err := files.CopyTerraformFolderToTemp("./", "azurermPrettier")
	if err != nil {
		return nil, fmt.Errorf("error creating temp TF code folder: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(tmpFolder)
	}()

	execPath, teardown, err := terraformExecPath()
	if err != nil {
		return nil, err
	}
	if teardown != nil {
		defer teardown()
	}
	workingDir := tmpFolder
	tf, err := tfexec.NewTerraform(workingDir, *execPath)
	if err != nil {
		return nil, fmt.Errorf("error running NewTerraform: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		return nil, fmt.Errorf("error running Init: %s", err)
	}
	schema, err := tf.ProvidersSchema(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error running providers: %s", err)
	}
	r := schema.Schemas["registry.terraform.io/hashicorp/azurerm"]

	return r, nil
}

func terraformExecPath() (*string, func(), error) {
	var teardown func()
	execPath, err := findTerraformExecPath()
	if err != nil {
		return nil, nil, fmt.Errorf("error finding Terraform: %s", err)
	}
	if execPath == nil {
		installer := &releases.LatestVersion{
			Product:                  product.Terraform,
			IncludePrereleases:       false,
			SkipChecksumVerification: false,
		}
		tp, err := installer.Install(context.Background())
		if err != nil {
			return nil, nil, fmt.Errorf("error installing Terraform: %s", err)
		}
		teardown = func() {
			_ = os.Remove(tp)
		}
		execPath = &tp
	}
	return execPath, teardown, nil
}

func findTerraformExecPath() (*string, error) {
	terraformExecName := "terraform"
	if os.Getenv("GOOS") == "windows" {
		terraformExecName = "terraform.exe"
	}

	pathEnv := os.Getenv("PATH")
	if pathEnv == "" {
		return nil, fmt.Errorf("PATH environment variable not set")
	}

	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	for _, p := range paths {
		execPath := filepath.Join(p, terraformExecName)
		if _, err := os.Stat(execPath); err == nil {
			return &execPath, nil
		}
	}

	return nil, nil
}

func SaveProviderSchema(s *tfjson.ProviderSchema) error {
	err := saveResourceSchemas(s)
	if err != nil {
		return fmt.Errorf("error saving resource schemas: %w", err)
	}
	err = saveDataSourceSchemas(s)
	if err != nil {
		return fmt.Errorf("error saving data source schemas: %w", err)
	}
	err = saveRegisterCode(s)
	if err != nil {
		return fmt.Errorf("error saving register code: %w", err)
	}
	return nil
}

func saveRegisterCode(s *tfjson.ProviderSchema) error {
	parameter := buildRegisterParameter(s)
	err := saveRegister(registerTemplate, parameter, "../generated/register.go")
	if err != nil {
		return fmt.Errorf("error saving register code: %w", err)
	}
	err = saveRegister(registerTestTemplate, parameter, "../generated/register_test.go")
	if err != nil {
		return fmt.Errorf("error saving register test code: %w", err)
	}

	return nil
}

func buildRegisterParameter(s *tfjson.ProviderSchema) RegisterParameter {
	parameter := RegisterParameter{
		ResourceSchemas:   make(map[string]string, 0),
		DataSourceSchemas: make(map[string]string, 0),
		RepoOwner:         "lonegunmanb",
		GoModule:          "azurerm-provider-schema",
	}
	linq.From(s.ResourceSchemas).OrderBy(byKey).ToMapBy(&parameter.ResourceSchemas, byKey, func(i interface{}) interface{} {
		pair := i.(linq.KeyValue)
		return fmt.Sprintf("resource.%sSchema()", strcase.ToCamel(pair.Key.(string)))
	})
	linq.From(s.DataSourceSchemas).OrderBy(byKey).ToMapBy(&parameter.DataSourceSchemas, byKey, func(i interface{}) interface{} {
		pair := i.(linq.KeyValue)
		return fmt.Sprintf("data.%sSchema()", strcase.ToCamel(pair.Key.(string)))
	})
	return parameter
}

func saveRegister(registerTemplate string, parameter RegisterParameter, destFilePath string) error {
	tmpl, err := template.New("register").Parse(registerTemplate)
	if err != nil {
		return err
	}
	buff := bytes.Buffer{}
	err = tmpl.Execute(&buff, parameter)
	if err != nil {
		return err
	}
	err = save(destFilePath, buff.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func saveDataSourceSchemas(s *tfjson.ProviderSchema) error {
	for dataName, schema := range s.DataSourceSchemas {
		err := SaveDataSourceSchema(dataName, schema)
		if err != nil {
			return fmt.Errorf("error saving data source schema: %s", err)
		}
	}
	return nil
}

func saveResourceSchemas(s *tfjson.ProviderSchema) error {
	for resourceName, schema := range s.ResourceSchemas {
		err := SaveResourceSchema(resourceName, schema)
		if err != nil {
			return fmt.Errorf("error saving resource schema: %s", err)
		}
	}
	return nil
}

func SaveDataSourceSchema(name string, s *tfjson.Schema) error {
	return saveSchema(name, s, PackageData)
}

func SaveResourceSchema(name string, s *tfjson.Schema) error {
	return saveSchema(name, s, PackageResource)
}

func saveSchema(name string, s *tfjson.Schema, pkg Package) error {
	jsonSchema, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("error marshalling schema: %s", err)
	}
	content, err := GenerateGoFileContent(name, string(jsonSchema), pkg)
	if err != nil {
		return fmt.Errorf("error generating go file content: %s", err)
	}
	fileName := strcase.ToLowerCamel(name)
	err = save(fmt.Sprintf("../generated/%s/%s.go", pkg, fileName), []byte(content))
	if err != nil {
		return fmt.Errorf("error saving file generated/%s/%s.go: %s", pkg, fileName, err)
	}
	content, err = GenerateGoTestFileContent(name, pkg)
	if err != nil {
		return fmt.Errorf("error generating go test file content: %w", err)
	}
	err = save(fmt.Sprintf("../generated/%s/%s_test.go", pkg, fileName), []byte(content))
	if err != nil {
		return fmt.Errorf("error saving file generated/%s/%s_test.go: %w", pkg, fileName, err)
	}
	return nil
}

func save(path string, bytes []byte) error {
	dir := filepath.Dir(path)
	// Create all folders in the path if they don't exist
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); err == nil {
		// File exists, truncate it
		err = os.Truncate(path, 0)
		if err != nil {
			return err
		}
	}

	// Write JSON bytes to file
	err = os.WriteFile(path, bytes, 0600)
	if err != nil {
		return err
	}

	return nil
}

func byKey(i interface{}) interface{} {
	pair := i.(linq.KeyValue)
	return pair.Key
}
