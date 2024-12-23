package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermPostgresqlFlexibleServerVirtualEndpoint = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name of the Virtual Endpoint",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "replica_server_id": {
        "description": "The Resource ID of the *Replica* Postgres Flexible Server this should be associated with",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "source_server_id": {
        "description": "The Resource ID of the *Source* Postgres Flexible Server this should be associated with",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "type": {
        "description": "The type of Virtual Endpoint",
        "description_kind": "plain",
        "required": true,
        "type": "string"
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

func AzurermPostgresqlFlexibleServerVirtualEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermPostgresqlFlexibleServerVirtualEndpoint), &result)
	return &result
}
