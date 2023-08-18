package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlVirtualMachineGroup = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sql_image_offer": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sql_image_sku": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
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
      },
      "wsfc_domain_profile": {
        "block": {
          "attributes": {
            "cluster_bootstrap_account_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cluster_operator_account_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cluster_subnet_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "fqdn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "organizational_unit_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sql_service_account_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_account_primary_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "storage_account_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermMssqlVirtualMachineGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlVirtualMachineGroup), &result)
	return &result
}
