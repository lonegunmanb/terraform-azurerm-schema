package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppEnvironmentStorage = `{
  "block": {
    "attributes": {
      "access_key": {
        "description": "The Storage Account Access Key.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "access_mode": {
        "description": "The access mode to connect this storage to the Container App. Possible values include ` + "`" + `ReadOnly` + "`" + ` and ` + "`" + `ReadWrite` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "account_name": {
        "description": "The Azure Storage Account in which the Share to be used is located.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "container_app_environment_id": {
        "description": "The ID of the Container App Environment to which this storage belongs.",
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
        "description": "The name for this Storage.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "share_name": {
        "description": "The name of the Azure Storage Share to use.",
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

func AzurermContainerAppEnvironmentStorageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironmentStorage), &result)
	return &result
}
