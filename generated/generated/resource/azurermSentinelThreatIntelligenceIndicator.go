package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSentinelThreatIntelligenceIndicator = `{
  "block": {
    "attributes": {
      "confidence": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "created_by": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "created_on": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "defanged": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "extension": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "external_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "external_last_updated_time_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "guid": {
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
      "indicator_type": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "language": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "last_updated_time_utc": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "object_marking_refs": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "parsed_pattern": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "pattern_type_key": "string",
              "pattern_type_values": [
                "list",
                [
                  "object",
                  {
                    "value": "string",
                    "value_type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "pattern": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pattern_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "pattern_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "revoked": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "source": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "threat_types": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "validate_from_utc": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "validate_until_utc": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "external_reference": {
        "block": {
          "attributes": {
            "description": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "hashes": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "source_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "granular_marking": {
        "block": {
          "attributes": {
            "language": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "marking_ref": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "selectors": {
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
        "nesting_mode": "list"
      },
      "kill_chain_phase": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
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

func AzurermSentinelThreatIntelligenceIndicatorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSentinelThreatIntelligenceIndicator), &result)
	return &result
}
