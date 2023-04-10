package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermConsumptionBudgetResourceGroupSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermConsumptionBudgetResourceGroupSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
