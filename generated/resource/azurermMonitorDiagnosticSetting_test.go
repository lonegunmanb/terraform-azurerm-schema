package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azurerm-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzurermMonitorDiagnosticSettingSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzurermMonitorDiagnosticSettingSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
