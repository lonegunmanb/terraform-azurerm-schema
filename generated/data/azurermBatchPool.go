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
                    "user_assigned_identity_id": "string",
                    "user_name": "string"
                  }
                ]
              ],
              "type": "string"
            }
          ]
        ]
      },
      "data_disks": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "caching": "string",
              "disk_size_gb": "number",
              "lun": "number",
              "storage_account_type": "string"
            }
          ]
        ]
      },
      "disk_encryption": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disk_encryption_target": "string"
            }
          ]
        ]
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "extensions": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_upgrade_minor_version": "bool",
              "name": "string",
              "protected_settings": "string",
              "provision_after_extensions": [
                "set",
                "string"
              ],
              "publisher": "string",
              "settings_json": "string",
              "type": "string",
              "type_handler_version": "string"
            }
          ]
        ]
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
      "inter_node_communication": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "license_type": {
        "computed": true,
        "description_kind": "plain",
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
      "mount": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "azure_blob_file_system": [
                "list",
                [
                  "object",
                  {
                    "account_key": "string",
                    "account_name": "string",
                    "blobfuse_options": "string",
                    "container_name": "string",
                    "identity_id": "string",
                    "relative_mount_path": "string",
                    "sas_key": "string"
                  }
                ]
              ],
              "azure_file_share": [
                "list",
                [
                  "object",
                  {
                    "account_key": "string",
                    "account_name": "string",
                    "azure_file_url": "string",
                    "mount_options": "string",
                    "relative_mount_path": "string"
                  }
                ]
              ],
              "cifs_mount": [
                "list",
                [
                  "object",
                  {
                    "mount_options": "string",
                    "password": "string",
                    "relative_mount_path": "string",
                    "source": "string",
                    "user_name": "string"
                  }
                ]
              ],
              "nfs_mount": [
                "list",
                [
                  "object",
                  {
                    "mount_options": "string",
                    "relative_mount_path": "string",
                    "source": "string"
                  }
                ]
              ]
            }
          ]
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
              "accelerated_networking_enabled": "bool",
              "dynamic_vnet_assignment_scope": "string",
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
                          "source_address_prefix": "string",
                          "source_port_ranges": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "protocol": "string"
                  }
                ]
              ],
              "public_address_provisioning_type": "string",
              "public_ips": [
                "set",
                "string"
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
      "node_placement": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "policy": "string"
            }
          ]
        ]
      },
      "os_disk_placement": {
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
              "container": [
                "list",
                [
                  "object",
                  {
                    "image_name": "string",
                    "registry": [
                      "list",
                      [
                        "object",
                        {
                          "password": "string",
                          "registry_server": "string",
                          "user_assigned_identity_id": "string",
                          "user_name": "string"
                        }
                      ]
                    ],
                    "run_options": "string",
                    "working_directory": "string"
                  }
                ]
              ],
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
                    "storage_container_url": "string",
                    "user_assigned_identity_id": "string"
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
      "task_scheduling_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "node_fill_type": "string"
            }
          ]
        ]
      },
      "user_accounts": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "elevation_level": "string",
              "linux_user_configuration": [
                "list",
                [
                  "object",
                  {
                    "gid": "number",
                    "ssh_private_key": "string",
                    "uid": "number"
                  }
                ]
              ],
              "name": "string",
              "password": "string",
              "windows_user_configuration": [
                "list",
                [
                  "object",
                  {
                    "login_mode": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "vm_size": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "windows": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enable_automatic_updates": "bool"
            }
          ]
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

func AzurermBatchPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermBatchPool), &result)
	return &result
}
