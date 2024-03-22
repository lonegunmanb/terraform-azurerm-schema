package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMysqlServer = `{
  "block": {
    "attributes": {
      "administrator_login": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "auto_grow_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "backup_retention_days": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "geo_redundant_backup_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "infrastructure_encryption_enabled": {
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
      "restore_point_in_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sku_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ssl_enforcement_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ssl_minimal_tls_version_enforced": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_mb": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "threat_detection_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disabled_alerts": [
                "set",
                "string"
              ],
              "email_account_admins": "bool",
              "email_addresses": [
                "set",
                "string"
              ],
              "enabled": "bool",
              "retention_days": "number",
              "storage_account_access_key": "string",
              "storage_endpoint": "string"
            }
          ]
        ]
      },
      "version": {
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
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermMysqlServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMysqlServer), &result)
	return &result
}
