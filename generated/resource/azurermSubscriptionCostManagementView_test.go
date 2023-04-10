package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermSubscriptionCostManagementViewSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermSubscriptionCostManagementViewSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
