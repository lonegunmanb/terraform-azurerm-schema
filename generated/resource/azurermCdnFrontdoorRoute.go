package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCdnFrontdoorRoute = `{
  "block": {
    "attributes": {
      "cdn_frontdoor_custom_domain_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cdn_frontdoor_endpoint_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cdn_frontdoor_origin_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "cdn_frontdoor_origin_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "cdn_frontdoor_origin_path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cdn_frontdoor_rule_set_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "forwarding_protocol": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "https_redirect_enabled": {
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
      "link_to_default_domain": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "patterns_to_match": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "supported_protocols": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "cache": {
        "block": {
          "attributes": {
            "compression_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "content_types_to_compress": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "query_string_caching_behavior": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "query_strings": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
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

func AzurermCdnFrontdoorRouteSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCdnFrontdoorRoute), &result)
	return &result
}
