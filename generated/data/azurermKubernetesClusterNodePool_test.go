package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/v4/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermKubernetesClusterNodePoolSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermKubernetesClusterNodePoolSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
