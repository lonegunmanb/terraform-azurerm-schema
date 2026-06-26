package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappVolumeBucketWithServer = `{
  "block": {
    "attributes": {
      "file_system_cifs_username": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "file_system_nfs_user": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "group_id": "number",
              "user_id": "number"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_vault": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "certificate_key_vault_uri": "string",
              "certificate_name": "string",
              "credentials_key_vault_uri": "string",
              "credentials_secret_name": "string"
            }
          ]
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "netapp_volume_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "path": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "permissions": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "server": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "certificate_pem": "string",
              "fqdn": "string",
              "on_certificate_conflict_action": "string"
            }
          ]
        ]
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

func AzurermNetappVolumeBucketWithServerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappVolumeBucketWithServer), &result)
	return &result
}
