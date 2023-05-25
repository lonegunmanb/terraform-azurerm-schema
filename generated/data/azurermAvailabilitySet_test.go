package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermAvailabilitySetSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermAvailabilitySetSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
