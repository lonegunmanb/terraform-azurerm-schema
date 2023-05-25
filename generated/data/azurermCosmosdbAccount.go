package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCosmosdbAccount = `{
  "block": {
    "attributes": {
      "capabilities": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string"
            }
          ]
        ]
      },
      "consistency_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "consistency_level": "string",
              "max_interval_in_seconds": "number",
              "max_staleness_prefix": "number"
            }
          ]
        ]
      },
      "enable_automatic_failover": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_free_tier": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "enable_multiple_write_locations": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "geo_location": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "failover_priority": "number",
              "id": "string",
              "location": "string"
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
      "ip_range_filter": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "is_virtual_network_filter_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "key_vault_key_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kind": {
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
      "offer_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "primary_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_master_key": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_readonly_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_readonly_master_key": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "read_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_master_key": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_readonly_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_readonly_master_key": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "virtual_network_rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string"
            }
          ]
        ]
      },
      "write_endpoints": {
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

func AzurermCosmosdbAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCosmosdbAccount), &result)
	return &result
}
