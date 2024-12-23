package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetworkManagerConnectivityConfiguration = `{
  "block": {
    "attributes": {
      "connectivity_topology": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "delete_existing_peering_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "global_mesh_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_manager_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "applies_to_group": {
        "block": {
          "attributes": {
            "global_mesh_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "group_connectivity": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_group_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "use_hub_gateway": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "hub": {
        "block": {
          "attributes": {
            "resource_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "resource_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AzurermNetworkManagerConnectivityConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetworkManagerConnectivityConfiguration), &result)
	return &result
}
