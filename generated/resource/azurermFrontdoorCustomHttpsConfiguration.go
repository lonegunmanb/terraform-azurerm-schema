package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermFrontdoorCustomHttpsConfiguration = `{
  "block": {
    "attributes": {
      "custom_https_provisioning_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "frontend_endpoint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "custom_https_configuration": {
        "block": {
          "attributes": {
            "azure_key_vault_certificate_secret_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "azure_key_vault_certificate_secret_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "azure_key_vault_certificate_vault_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "certificate_source": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "minimum_tls_version": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "provisioning_state": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "provisioning_substate": {
              "computed": true,
              "description_kind": "plain",
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
  "version": 1
}`

func AzurermFrontdoorCustomHttpsConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermFrontdoorCustomHttpsConfiguration), &result)
	return &result
}
