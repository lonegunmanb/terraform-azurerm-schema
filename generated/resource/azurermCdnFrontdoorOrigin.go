package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCdnFrontdoorOrigin = `{
  "block": {
    "attributes": {
      "cdn_frontdoor_origin_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_name_check_enabled": {
        "description_kind": "plain",
        "required": true,
        "type": "bool"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "health_probes_enabled": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "host_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "http_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "https_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
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
      "origin_host_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "weight": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "private_link": {
        "block": {
          "attributes": {
            "location": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "private_link_target_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "request_message": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_type": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermCdnFrontdoorOriginSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCdnFrontdoorOrigin), &result)
	return &result
}
