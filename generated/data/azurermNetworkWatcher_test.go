package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermNetworkWatcherSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermNetworkWatcherSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
