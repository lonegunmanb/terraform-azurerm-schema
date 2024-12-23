package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCosmosdbSqlContainer = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "analytical_storage_ttl": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "database_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "default_ttl": {
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
      "partition_key_kind": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "partition_key_paths": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "partition_key_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "throughput": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "autoscale_settings": {
        "block": {
          "attributes": {
            "max_throughput": {
              "computed": true,
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
      "conflict_resolution_policy": {
        "block": {
          "attributes": {
            "conflict_resolution_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "conflict_resolution_procedure": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mode": {
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
      "indexing_policy": {
        "block": {
          "attributes": {
            "indexing_mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "composite_index": {
              "block": {
                "block_types": {
                  "index": {
                    "block": {
                      "attributes": {
                        "order": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "excluded_path": {
              "block": {
                "attributes": {
                  "path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "included_path": {
              "block": {
                "attributes": {
                  "path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "spatial_index": {
              "block": {
                "attributes": {
                  "path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "types": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "set",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
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
      },
      "unique_key": {
        "block": {
          "attributes": {
            "paths": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AzurermCosmosdbSqlContainerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCosmosdbSqlContainer), &result)
	return &result
}
