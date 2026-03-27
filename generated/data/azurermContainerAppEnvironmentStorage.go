package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppEnvironmentStorage = `{
  "block": {
    "attributes": {
      "access_mode": {
        "computed": true,
        "description": "The access mode to connect this storage to the Container App.",
        "description_kind": "plain",
        "type": "string"
      },
      "account_name": {
        "computed": true,
        "description": "The Azure Storage Account in which the Share is located.",
        "description_kind": "plain",
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
        "description": "The name for this Container App Environment Storage.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nfs_server_url": {
        "computed": true,
        "description": "The NFS server URL for the Azure File Share.",
        "description_kind": "plain",
        "type": "string"
      },
      "share_name": {
        "computed": true,
        "description": "The name of the Azure Storage Share.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
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

func AzurermContainerAppEnvironmentStorageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironmentStorage), &result)
	return &result
}
