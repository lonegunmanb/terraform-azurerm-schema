package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermEventhubNamespaceAuthorizationRule = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "listen": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "manage": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "primary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_connection_string_alias": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "primary_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "secondary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_connection_string_alias": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "secondary_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "send": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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

func AzurermEventhubNamespaceAuthorizationRuleSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermEventhubNamespaceAuthorizationRule), &result)
	return &result
}
