package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v3/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermSecurityCenterSettingSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermSecurityCenterSettingSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
