package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermPublicIp = `{
  "block": {
    "attributes": {
      "allocation_method": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ddos_protection_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ddos_protection_plan_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name_label": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edge_zone": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fqdn": {
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
      "idle_timeout_in_minutes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ip_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "ip_version": {
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
      "public_ip_prefix_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reverse_fqdn": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sku": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sku_tier": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermPublicIpSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermPublicIp), &result)
	return &result
}
