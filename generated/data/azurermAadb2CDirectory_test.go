package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermAadb2CDirectorySchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermAadb2CDirectorySchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
