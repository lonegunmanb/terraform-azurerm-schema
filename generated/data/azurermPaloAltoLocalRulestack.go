package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermPaloAltoLocalRulestack = `{
  "block": {
    "attributes": {
      "anti_spyware_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "anti_virus_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "description": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_subscription": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_blocking_profile": {
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
      "outbound_trust_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "outbound_untrust_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "url_filtering_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vulnerability_profile": {
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

func AzurermPaloAltoLocalRulestackSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermPaloAltoLocalRulestack), &result)
	return &result
}
