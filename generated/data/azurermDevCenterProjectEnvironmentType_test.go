package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermDevCenterProjectEnvironmentTypeSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermDevCenterProjectEnvironmentTypeSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
