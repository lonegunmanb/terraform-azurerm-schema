package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNginxDeployment = `{
  "block": {
    "attributes": {
      "auto_scale_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "max_capacity": "number",
              "min_capacity": "number",
              "name": "string"
            }
          ]
        ]
      },
      "automatic_upgrade_channel": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "capacity": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "diagnose_support_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "frontend_private": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allocation_method": "string",
              "ip_address": "string",
              "subnet_id": "string"
            }
          ]
        ]
      },
      "frontend_public": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ip_address": [
                "list",
                "string"
              ]
            }
          ]
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
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logging_storage_account": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "container_name": "string",
              "name": "string"
            }
          ]
        ]
      },
      "managed_resource_group": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "subnet_id": "string"
            }
          ]
        ]
      },
      "nginx_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermNginxDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNginxDeployment), &result)
	return &result
}
