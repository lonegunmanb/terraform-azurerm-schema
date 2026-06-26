package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappVolumeBucket = `{
  "block": {
    "attributes": {
      "file_system_cifs_username": {
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
      "path": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "permissions": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "server_certificate_common_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_certificate_expiry_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server_ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "volume_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "file_system_nfs_user": {
        "block": {
          "attributes": {
            "group_id": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "user_id": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "key_vault": {
        "block": {
          "attributes": {
            "certificate_key_vault_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "certificate_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "credentials_key_vault_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "credentials_secret_name": {
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

func AzurermNetappVolumeBucketSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappVolumeBucket), &result)
	return &result
}
