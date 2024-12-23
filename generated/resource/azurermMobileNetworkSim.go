package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkSim = `{
  "block": {
    "attributes": {
      "authentication_key": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "device_type": {
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
      "integrated_circuit_card_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "international_mobile_subscriber_identity": {
        "description_kind": "plain",
        "required": true,
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
      "operator_key_code": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "sim_policy_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sim_state": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "static_ip_configuration": {
        "block": {
          "attributes": {
            "attached_data_network_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "slice_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "static_ipv4_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AzurermMobileNetworkSimSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMobileNetworkSim), &result)
	return &result
}
