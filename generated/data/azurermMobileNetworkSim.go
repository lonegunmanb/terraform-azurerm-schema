package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkSim = `{
  "block": {
    "attributes": {
      "device_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "integrated_circuit_card_identifier": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "international_mobile_subscriber_identity": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mobile_network_sim_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sim_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sim_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "static_ip_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "attached_data_network_id": "string",
              "slice_id": "string",
              "static_ipv4_address": "string"
            }
          ]
        ]
      },
      "vendor_key_fingerprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vendor_name": {
        "computed": true,
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

func AzurermMobileNetworkSimSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMobileNetworkSim), &result)
	return &result
}
