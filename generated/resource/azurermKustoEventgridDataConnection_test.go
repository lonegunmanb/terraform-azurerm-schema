package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v4/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermKustoEventgridDataConnectionSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermKustoEventgridDataConnectionSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
