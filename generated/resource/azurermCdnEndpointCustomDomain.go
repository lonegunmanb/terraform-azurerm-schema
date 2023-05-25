package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCdnEndpointCustomDomain = `{
  "block": {
    "attributes": {
      "cdn_endpoint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "host_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "cdn_managed_https": {
        "block": {
          "attributes": {
            "certificate_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tls_version": {
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
      },
      "user_managed_https": {
        "block": {
          "attributes": {
            "key_vault_certificate_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tls_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermCdnEndpointCustomDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCdnEndpointCustomDomain), &result)
	return &result
}
