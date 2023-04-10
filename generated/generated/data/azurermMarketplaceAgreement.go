package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMarketplaceAgreement = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_text_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "offer": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "plan": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "privacy_policy_link": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "publisher": {
        "description_kind": "plain",
        "required": true,
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermMarketplaceAgreementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMarketplaceAgreement), &result)
	return &result
}
