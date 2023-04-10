package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/azurerm-provider-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermDiskAccessSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermDiskAccessSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}