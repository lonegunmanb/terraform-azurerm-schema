package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleExascaleDatabaseStorageVault = `{
  "block": {
    "attributes": {
      "additional_flash_cache_percentage": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "high_capacity_database_storage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "available_size_in_gb": "number",
              "total_size_in_gb": "number"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "lifecycle_details": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "lifecycle_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "oci_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ocid": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "time_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "virtual_machine_cluster_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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

func AzurermOracleExascaleDatabaseStorageVaultSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleExascaleDatabaseStorageVault), &result)
	return &result
}
