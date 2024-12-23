package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSignalrService = `{
  "block": {
    "attributes": {
      "aad_auth_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "connectivity_logs_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "http_request_logs_enabled": {
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
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "live_trace_enabled": {
        "deprecated": true,
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
      "messaging_logs_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "serverless_connection_timeout_in_seconds": {
        "description_kind": "plain",
        "optional": true,
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
      "cors": {
        "block": {
          "attributes": {
            "allowed_origins": {
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
      "live_trace": {
        "block": {
          "attributes": {
            "connectivity_logs_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "http_request_logs_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "messaging_logs_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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
      },
      "upstream_endpoint": {
        "block": {
          "attributes": {
            "category_pattern": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "event_pattern": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "hub_pattern": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "url_template": {
              "description_kind": "plain",
              "required": true,
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
        "nesting_mode": "set"
      }
    },
    "description_kind": "plain"
  },
  "version": 1
}`

func AzurermSignalrServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSignalrService), &result)
	return &result
}
