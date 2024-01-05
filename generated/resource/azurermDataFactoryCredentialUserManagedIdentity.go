package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataFactoryCredentialUserManagedIdentity = `{
  "block": {
    "attributes": {
      "annotations": {
        "description": "(Optional) List of string annotations.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "data_factory_id": {
        "description": "The resource ID of the parent Data Factory",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description": "(Optional) Short text description",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity_id": {
        "description": "The resource ID of the User Assigned Managed Identity",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description": "The desired name of the credential resource",
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
            },
            "update": {
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

func AzurermDataFactoryCredentialUserManagedIdentitySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataFactoryCredentialUserManagedIdentity), &result)
	return &result
}
