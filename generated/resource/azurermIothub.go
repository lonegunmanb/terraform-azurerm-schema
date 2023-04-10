package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermIothub = `{
  "block": {
    "attributes": {
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "authentication_type": "string",
              "batch_frequency_in_seconds": "number",
              "connection_string": "string",
              "container_name": "string",
              "encoding": "string",
              "endpoint_uri": "string",
              "entity_path": "string",
              "file_name_format": "string",
              "identity_id": "string",
              "max_chunk_size_in_bytes": "number",
              "name": "string",
              "resource_group_name": "string",
              "type": "string"
            }
          ]
        ]
      },
      "enrichment": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "endpoint_names": [
                "list",
                "string"
              ],
              "key": "string",
              "value": "string"
            }
          ]
        ]
      },
      "event_hub_events_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_hub_events_namespace": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_hub_events_path": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_hub_operations_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_hub_operations_path": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "event_hub_partition_count": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "event_hub_retention_in_days": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "hostname": {
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "min_tls_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "route": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "condition": "string",
              "enabled": "bool",
              "endpoint_names": [
                "list",
                "string"
              ],
              "name": "string",
              "source": "string"
            }
          ]
        ]
      },
      "shared_access_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "key_name": "string",
              "permissions": "string",
              "primary_key": "string",
              "secondary_key": "string"
            }
          ]
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "cloud_to_device": {
        "block": {
          "attributes": {
            "default_ttl": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_delivery_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "feedback": {
              "block": {
                "attributes": {
                  "lock_duration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "max_delivery_count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "time_to_live": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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
      "fallback_route": {
        "block": {
          "attributes": {
            "condition": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "endpoint_names": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "source": {
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
      "file_upload": {
        "block": {
          "attributes": {
            "authentication_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "connection_string": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "container_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "default_ttl": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "identity_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "lock_duration": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "max_delivery_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "notifications": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sas_ttl": {
              "computed": true,
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
      "network_rule_set": {
        "block": {
          "attributes": {
            "apply_to_builtin_eventhub_endpoint": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "default_action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "ip_rule": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ip_mask": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "sku": {
        "block": {
          "attributes": {
            "capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
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
  "version": 1
}`

func AzurermIothubSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermIothub), &result)
	return &result
}
