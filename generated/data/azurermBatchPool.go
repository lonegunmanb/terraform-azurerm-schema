package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermBatchPool = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "auto_scale": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "evaluation_interval": "string",
              "formula": "string"
            }
          ]
        ]
      },
      "certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "store_location": "string",
              "store_name": "string",
              "visibility": [
                "set",
                "string"
              ]
            }
          ]
        ]
      },
      "container_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "container_image_names": [
                "set",
                "string"
              ],
              "container_registries": [
                "list",
                [
                  "object",
                  {
                    "password": "string",
                    "registry_server": "string",
                    "user_name": "string"
                  }
                ]
              ],
              "type": "string"
            }
          ]
        ]
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fixed_scale": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "resize_timeout": "string",
              "target_dedicated_nodes": "number",
              "target_low_priority_nodes": "number"
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
      "max_tasks_per_node": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "endpoint_configuration": [
                "list",
                [
                  "object",
                  {
                    "backend_port": "number",
                    "frontend_port_range": "string",
                    "name": "string",
                    "network_security_group_rules": [
                      "list",
                      [
                        "object",
                        {
                          "access": "string",
                          "priority": "number",
                          "source_address_prefix": "string"
                        }
                      ]
                    ],
                    "protocol": "string"
                  }
                ]
              ],
              "subnet_id": "string"
            }
          ]
        ]
      },
      "node_agent_sku_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "start_task": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "command_line": "string",
              "common_environment_properties": [
                "map",
                "string"
              ],
              "environment": [
                "map",
                "string"
              ],
              "max_task_retry_count": "number",
              "resource_file": [
                "list",
                [
                  "object",
                  {
                    "auto_storage_container_name": "string",
                    "blob_prefix": "string",
                    "file_mode": "string",
                    "file_path": "string",
                    "http_url": "string",
                    "storage_container_url": "string"
                  }
                ]
              ],
              "task_retry_maximum": "number",
              "user_identity": [
                "list",
                [
                  "object",
                  {
                    "auto_user": [
                      "list",
                      [
                        "object",
                        {
                          "elevation_level": "string",
                          "scope": "string"
                        }
                      ]
                    ],
                    "user_name": "string"
                  }
                ]
              ],
              "wait_for_success": "bool"
            }
          ]
        ]
      },
      "storage_image_reference": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "offer": "string",
              "publisher": "string",
              "sku": "string",
              "version": "string"
            }
          ]
        ]
      },
      "vm_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AzurermBatchPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBatchPool), &result)
	return &result
}
