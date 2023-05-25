package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSynapseSparkPool = `{
  "block": {
    "attributes": {
      "cache_size": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "compute_isolation_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "dynamic_executor_allocation_enabled": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "node_size": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_size_family": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "session_level_packages_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "spark_events_folder": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spark_log_folder": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "spark_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "synapse_workspace_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      }
    },
    "block_types": {
      "auto_pause": {
        "block": {
          "attributes": {
            "delay_in_minutes": {
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
      "auto_scale": {
        "block": {
          "attributes": {
            "max_node_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_node_count": {
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
      "library_requirement": {
        "block": {
          "attributes": {
            "content": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "filename": {
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
      "spark_config": {
        "block": {
          "attributes": {
            "content": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "filename": {
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

func AzurermSynapseSparkPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSynapseSparkPool), &result)
	return &result
}
