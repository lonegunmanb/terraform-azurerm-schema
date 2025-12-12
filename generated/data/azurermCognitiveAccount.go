package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCognitiveAccount = `{
  "block": {
    "attributes": {
      "custom_question_answering_search_service_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_subdomain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "customer_managed_key": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_client_id": "string",
              "key_vault_key_id": "string"
            }
          ]
        ]
      },
      "dynamic_throttling_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fqdns": {
        "computed": true,
        "description_kind": "plain",
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
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_ids": [
                "list",
                "string"
              ],
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "kind": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "local_auth_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metrics_advisor_aad_client_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metrics_advisor_aad_tenant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metrics_advisor_super_user_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "metrics_advisor_website_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_acls": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "bypass": "string",
              "default_action": "string",
              "ip_rules": [
                "list",
                "string"
              ],
              "virtual_network_rules": [
                "list",
                [
                  "object",
                  {
                    "ignore_missing_vnet_service_endpoint": "bool",
                    "subnet_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "network_injection": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "scenario": "string",
              "subnet_id": "string"
            }
          ]
        ]
      },
      "outbound_network_access_restricted": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "primary_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "project_management_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "public_network_access_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "qna_runtime_endpoint": {
        "computed": true,
        "description_kind": "plain",
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_client_id": "string",
              "storage_account_id": "string"
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

func AzurermCognitiveAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCognitiveAccount), &result)
	return &result
}
