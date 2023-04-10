package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermLogAnalyticsQueryPackSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermLogAnalyticsQueryPackSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
