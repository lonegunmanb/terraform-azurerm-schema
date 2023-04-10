package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/azurerm-provider-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzurermMonitorDiagnosticCategoriesSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzurermMonitorDiagnosticCategoriesSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
