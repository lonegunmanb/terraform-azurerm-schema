package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAppServiceCertificateOrder = `{
  "block": {
    "attributes": {
      "app_service_certificate_not_renewable_reasons": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "auto_renew": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "certificates": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "certificate_name": "string",
              "key_vault_id": "string",
              "key_vault_secret_name": "string",
              "provisioning_state": "string"
            }
          ]
        ]
      },
      "csr": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "distinguished_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_verification_token": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_time": {
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
      "intermediate_thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "is_private_key_external": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "key_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "product_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "root_thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "signed_certificate_thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
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
      "validity_in_years": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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

func AzurermAppServiceCertificateOrderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAppServiceCertificateOrder), &result)
	return &result
}
