package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermWebPubsubSocketio = `{
  "block": {
    "attributes": {
      "aad_auth_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "external_ip": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "live_trace_connectivity_logs_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "live_trace_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "live_trace_http_request_logs_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "live_trace_messaging_logs_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "primary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "public_network_access": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "secondary_connection_string": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "server_port": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "service_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "tls_client_cert_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "sku": {
        "block": {
          "attributes": {
            "capacity": {
              "description_kind": "plain",
              "optional": true,
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
  "version": 0
}`

func AzurermWebPubsubSocketioSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermWebPubsubSocketio), &result)
	return &result
}
