package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppCustomDomain = `{
  "block": {
    "attributes": {
      "certificate_binding_type": {
        "description": "The Binding type. Possible values include ` + "`" + `Disabled` + "`" + ` and ` + "`" + `SniEnabled` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "container_app_environment_certificate_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "container_app_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The hostname of the Certificate. Must be the CN or a named SAN in the certificate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermContainerAppCustomDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppCustomDomain), &result)
	return &result
}
