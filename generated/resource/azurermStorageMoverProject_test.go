package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v4/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermStorageMoverProjectSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermStorageMoverProjectSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
