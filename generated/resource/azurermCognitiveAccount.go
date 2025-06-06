package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCognitiveAccount = `{
  "block": {
    "attributes": {
      "custom_question_answering_search_service_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_question_answering_search_service_key": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "custom_subdomain_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dynamic_throttling_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fqdns": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kind": {
        "description_kind": "plain",
        "required": true,
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
      "metrics_advisor_aad_client_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metrics_advisor_aad_tenant_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metrics_advisor_super_user_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "metrics_advisor_website_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "outbound_network_access_restricted": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "qna_runtime_endpoint": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "sku_name": {
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
      "customer_managed_key": {
        "block": {
          "attributes": {
            "identity_client_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_key_id": {
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
      "network_acls": {
        "block": {
          "attributes": {
            "bypass": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ip_rules": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            }
          },
          "block_types": {
            "virtual_network_rules": {
              "block": {
                "attributes": {
                  "ignore_missing_vnet_service_endpoint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "subnet_id": {
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
        "nesting_mode": "list"
      },
      "storage": {
        "block": {
          "attributes": {
            "identity_client_id": {
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

func AzurermCognitiveAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCognitiveAccount), &result)
	return &result
}
