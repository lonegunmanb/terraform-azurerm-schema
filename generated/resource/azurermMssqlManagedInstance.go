package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlManagedInstance = `{
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
      "collation": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "database_format": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_zone": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_zone_partner_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hybrid_secondary_usage": {
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
      "license_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maintenance_configuration_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "minimum_tls_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "proxy_override": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_data_endpoint_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_principal_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sku_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_account_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_size_in_gb": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "subnet_id": {
        "description_kind": "plain",
        "required": true,
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
      "timezone_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vcores": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "zone_redundant_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "azure_active_directory_administrator": {
        "block": {
          "attributes": {
            "azuread_authentication_only_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "login_username": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "object_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "principal_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tenant_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AzurermMssqlManagedInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlManagedInstance), &result)
	return &result
}
