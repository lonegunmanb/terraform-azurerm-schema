package resource

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
      "storage_account_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "rule": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "actions": {
              "block": {
                "block_types": {
                  "base_blob": {
                    "block": {
                      "attributes": {
                        "auto_tier_to_hot_from_cool_enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "delete_after_days_since_creation_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "delete_after_days_since_last_access_time_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "delete_after_days_since_modification_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_archive_after_days_since_creation_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_archive_after_days_since_last_access_time_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_archive_after_days_since_last_tier_change_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_archive_after_days_since_modification_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_cool_after_days_since_creation_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_cool_after_days_since_last_access_time_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_cool_after_days_since_modification_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "snapshot": {
                    "block": {
                      "attributes": {
                        "change_tier_to_archive_after_days_since_creation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "change_tier_to_cool_after_days_since_creation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "delete_after_days_since_creation_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_archive_after_days_since_last_tier_change_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "version": {
                    "block": {
                      "attributes": {
                        "change_tier_to_archive_after_days_since_creation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "change_tier_to_cool_after_days_since_creation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "delete_after_days_since_creation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "tier_to_archive_after_days_since_last_tier_change_greater_than": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "filters": {
              "block": {
                "attributes": {
                  "blob_types": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "prefix_match": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "match_blob_index_tag": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "operation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
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

func AzurermStorageManagementPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStorageManagementPolicy), &result)
	return &result
}
