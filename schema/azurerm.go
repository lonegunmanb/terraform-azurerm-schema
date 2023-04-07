package schema

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gruntwork-io/terratest/modules/files"
	"github.com/hashicorp/hc-install/product"
	"github.com/hashicorp/hc-install/releases"
	"github.com/hashicorp/terraform-exec/tfexec"
	tfjson "github.com/hashicorp/terraform-json"
)

type ProviderSchema struct {
	*tfjson.ProviderSchema
	Version string `json:"version"`
}

func RefreshAzureRMSchema(path string) error {
	s, err := ExtractAzureRMProviderSchema()

	if err != nil {
		return err
	}
	return Save(s, path)
}

func ExtractAzureRMProviderSchema() (*ProviderSchema, error) {
	tmpFolder, err := files.CopyTerraformFolderToTemp("./", "azurermPrettier")
	if err != nil {
		return nil, fmt.Errorf("error creating temp TF code folder: %s", err)
	}
	defer func() {
		_ = os.RemoveAll(tmpFolder)
	}()
	installer := &releases.LatestVersion{
		Product:                  product.Terraform,
		IncludePrereleases:       false,
		SkipChecksumVerification: false,
	}
	execPath, err := installer.Install(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error installing Terraform: %s", err)
	}
	defer func() {
		_ = os.Remove(execPath)
	}()

	workingDir := tmpFolder
	tf, err := tfexec.NewTerraform(workingDir, execPath)
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
	_, providerVersions, err := tf.Version(context.Background(), false)
	if err != nil {
		return nil, fmt.Errorf("error running version: %s", err)
	}

	r := &ProviderSchema{
		ProviderSchema: schema.Schemas["registry.terraform.io/hashicorp/azurerm"],
	}
	v := providerVersions["registry.terraform.io/hashicorp/azurerm"]
	if v != nil {
		r.Version = v.String()
	}
	return r, nil
}

func Save(s *ProviderSchema, path string) error {
	//version := s.Version.String()
	// Convert ProviderSchema to JSON bytes
	bytes, err := json.Marshal(s)
	if err != nil {
		return err
	}

	err = save(path, bytes, false)
	return err
}

func save(path string, bytes []byte, skipIfExisted bool) error {
	if _, err := os.Stat(path); err == nil {
		if skipIfExisted {
			return nil
		}
		// File exists, truncate it
		err = os.Truncate(path, 0)
		if err != nil {
			return err
		}
	}

	// Write JSON bytes to file
	err := os.WriteFile(path, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Given such go function which takes `path` as file path that contains JSON bytes of `tfjson.ProviderSchema`, read the file content, unmarshal it to `tfjson.ProviderSchema` and return it.
func LoadSchema(path string) (*tfjson.ProviderSchema, error) {
	// Read file contents
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON data to ProviderSchema
	var schema tfjson.ProviderSchema
	err = json.Unmarshal(bytes, &schema)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}
