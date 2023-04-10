package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/azurerm-provider-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermWebPubsubSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermWebPubsubSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
