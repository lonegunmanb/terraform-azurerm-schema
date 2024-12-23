package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermEventgridDomain = `{
  "block": {
    "attributes": {
      "auto_create_topic_with_first_subscription": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "auto_delete_topic_with_last_subscription": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
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
      "inbound_ip_rule": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "action": "string",
              "ip_mask": "string"
            }
          ]
        ]
      },
      "input_schema": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "local_auth_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "primary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
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
      "secondary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
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
      "input_mapping_default_values": {
        "block": {
          "attributes": {
            "data_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "event_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subject": {
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
      "input_mapping_fields": {
        "block": {
          "attributes": {
            "data_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "event_time": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "event_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subject": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "topic": {
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

func AzurermEventgridDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermEventgridDomain), &result)
	return &result
}
