package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermContainerRegistryScopeMapSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermContainerRegistryScopeMapSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
