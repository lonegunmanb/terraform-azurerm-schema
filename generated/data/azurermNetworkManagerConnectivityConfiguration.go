package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetworkManagerConnectivityConfiguration = `{
  "block": {
    "attributes": {
      "applies_to_group": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "global_mesh_enabled": "bool",
              "group_connectivity": "string",
              "network_group_id": "string",
              "use_hub_gateway": "bool"
            }
          ]
        ]
      },
      "connectivity_topology": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "delete_existing_peering_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "global_mesh_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "hub": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resource_id": "string",
              "resource_type": "string"
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

func AzurermNetworkManagerConnectivityConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetworkManagerConnectivityConfiguration), &result)
	return &result
}
