package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCdnFrontdoorCustomDomain = `{
  "block": {
    "attributes": {
      "cdn_frontdoor_profile_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_zone_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "expiration_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "host_name": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cdn_frontdoor_secret_id": "string",
              "certificate_type": "string",
              "minimum_tls_version": "string"
            }
          ]
        ]
      },
      "validation_token": {
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermCdnFrontdoorCustomDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCdnFrontdoorCustomDomain), &result)
	return &result
}
