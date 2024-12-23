package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermEventgridDomain = `{
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
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_ids": [
                "list",
                "string"
              ],
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
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
      "input_mapping_default_values": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_version": "string",
              "event_type": "string",
              "subject": "string"
            }
          ]
        ]
      },
      "input_mapping_fields": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data_version": "string",
              "event_time": "string",
              "event_type": "string",
              "id": "string",
              "subject": "string",
              "topic": "string"
            }
          ]
        ]
      },
      "input_schema": {
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
      "primary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "public_network_access_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_access_key": {
        "computed": true,
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

func AzurermEventgridDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermEventgridDomain), &result)
	return &result
}
