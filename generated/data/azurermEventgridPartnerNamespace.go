package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermEventgridPartnerNamespace = `{
  "block": {
    "attributes": {
      "endpoint": {
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
      "inbound_ip_rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "action": "string",
              "ip_mask": "string"
            }
          ]
        ]
      },
      "local_authentication_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "partner_registration_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "partner_topic_routing_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_network_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermEventgridPartnerNamespaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermEventgridPartnerNamespace), &result)
	return &result
}
