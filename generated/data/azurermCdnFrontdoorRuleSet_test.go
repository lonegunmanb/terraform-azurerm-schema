package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v3/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermCdnFrontdoorRuleSetSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermCdnFrontdoorRuleSetSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
