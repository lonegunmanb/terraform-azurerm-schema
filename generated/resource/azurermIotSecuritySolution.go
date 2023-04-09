package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermIotSecuritySolution = `{
  "block": {
    "attributes": {
      "disabled_data_sources": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "display_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "events_to_export": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "iothub_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "set",
          "string"
        ]
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_analytics_workspace_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "log_unmasked_ips_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "query_for_resources": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "query_subscription_ids": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "resource_group_name": {
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
      "additional_workspace": {
        "block": {
          "attributes": {
            "data_types": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "workspace_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "recommendations_enabled": {
        "block": {
          "attributes": {
            "acr_authentication": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "agent_send_unutilized_msg": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "baseline": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "edge_hub_mem_optimize": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "edge_logging_option": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "inconsistent_module_settings": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "install_agent": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ip_filter_deny_all": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ip_filter_permissive_rule": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "open_ports": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "permissive_firewall_policy": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "permissive_input_firewall_rules": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "permissive_output_firewall_rules": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "privileged_docker_options": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "shared_credentials": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "vulnerable_tls_cipher_suite": {
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

func AzurermIotSecuritySolutionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermIotSecuritySolution), &result)
	return &result
}
