package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermElasticSanVolumeGroup = `{
  "block": {
    "attributes": {
      "elastic_san_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "encryption_type": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "protocol_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "encryption": {
        "block": {
          "attributes": {
            "current_versioned_key_expiration_timestamp": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "current_versioned_key_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_vault_key_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "last_key_rotation_timestamp": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "user_assigned_identity_id": {
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
      "network_rule": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
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

func AzurermElasticSanVolumeGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermElasticSanVolumeGroup), &result)
	return &result
}
