package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMariadbServer = `{
  "block": {
    "attributes": {
      "administrator_login": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "administrator_login_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "auto_grow_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "backup_retention_days": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "create_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "creation_source_server_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "geo_redundant_backup_enabled": {
        "computed": true,
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
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restore_point_in_time": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sku_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ssl_enforcement_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "ssl_minimal_tls_version_enforced": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_mb": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "version": {
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
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermMariadbServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMariadbServer), &result)
	return &result
}
