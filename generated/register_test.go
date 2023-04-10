package generated_test

import (
	"testing"

	"github.com/lonegunmanb/azurerm-provider-schema/generated"
	"github.com/stretchr/testify/assert"
)

func TestResourceSchema(t *testing.T) {
	assert.NotEmpty(t, generated.Resources)
	assert.NotEmpty(t, generated.DataSources)
}