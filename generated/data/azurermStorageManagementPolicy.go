package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStorageManagementPolicy = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actions": [
                "list",
                [
                  "object",
                  {
                    "base_blob": [
                      "list",
                      [
                        "object",
                        {
                          "auto_tier_to_hot_from_cool_enabled": "bool",
                          "delete_after_days_since_creation_greater_than": "number",
                          "delete_after_days_since_last_access_time_greater_than": "number",
                          "delete_after_days_since_modification_greater_than": "number",
                          "tier_to_archive_after_days_since_creation_greater_than": "number",
                          "tier_to_archive_after_days_since_last_access_time_greater_than": "number",
                          "tier_to_archive_after_days_since_last_tier_change_greater_than": "number",
                          "tier_to_archive_after_days_since_modification_greater_than": "number",
                          "tier_to_cool_after_days_since_creation_greater_than": "number",
                          "tier_to_cool_after_days_since_last_access_time_greater_than": "number",
                          "tier_to_cool_after_days_since_modification_greater_than": "number"
                        }
                      ]
                    ],
                    "snapshot": [
                      "list",
                      [
                        "object",
                        {
                          "change_tier_to_archive_after_days_since_creation": "number",
                          "change_tier_to_cool_after_days_since_creation": "number",
                          "delete_after_days_since_creation_greater_than": "number",
                          "tier_to_archive_after_days_since_last_tier_change_greater_than": "number"
                        }
                      ]
                    ],
                    "version": [
                      "list",
                      [
                        "object",
                        {
                          "change_tier_to_archive_after_days_since_creation": "number",
                          "change_tier_to_cool_after_days_since_creation": "number",
                          "delete_after_days_since_creation": "number",
                          "tier_to_archive_after_days_since_last_tier_change_greater_than": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "enabled": "bool",
              "filters": [
                "list",
                [
                  "object",
                  {
                    "blob_types": [
                      "set",
                      "string"
                    ],
                    "match_blob_index_tag": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "operation": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "prefix_match": [
                      "set",
                      "string"
                    ]
                  }
                ]
              ],
              "name": "string"
            }
          ]
        ]
      },
      "storage_account_id": {
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

func AzurermStorageManagementPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageManagementPolicy), &result)
	return &result
}
