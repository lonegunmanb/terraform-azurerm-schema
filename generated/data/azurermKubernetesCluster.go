package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKubernetesCluster = `{
  "block": {
    "attributes": {
      "aci_connector_linux": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "subnet_name": "string"
            }
          ]
        ]
      },
      "addon_profile": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "azure_keyvault_secrets_provider": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool",
                    "secret_identity": [
                      "list",
                      [
                        "object",
                        {
                          "client_id": "string",
                          "object_id": "string",
                          "user_assigned_identity_id": "string"
                        }
                      ]
                    ],
                    "secret_rotation_enabled": "string",
                    "secret_rotation_interval": "string"
                  }
                ]
              ],
              "azure_policy": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "http_application_routing": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool",
                    "http_application_routing_zone_name": "string"
                  }
                ]
              ],
              "ingress_application_gateway": [
                "list",
                [
                  "object",
                  {
                    "effective_gateway_id": "string",
                    "enabled": "bool",
                    "gateway_id": "string",
                    "ingress_application_gateway_identity": [
                      "list",
                      [
                        "object",
                        {
                          "client_id": "string",
                          "object_id": "string",
                          "user_assigned_identity_id": "string"
                        }
                      ]
                    ],
                    "subnet_cidr": "string",
                    "subnet_id": "string"
                  }
                ]
              ],
              "kube_dashboard": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ],
              "oms_agent": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool",
                    "log_analytics_workspace_id": "string",
                    "oms_agent_identity": [
                      "list",
                      [
                        "object",
                        {
                          "client_id": "string",
                          "object_id": "string",
                          "user_assigned_identity_id": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "open_service_mesh": [
                "list",
                [
                  "object",
                  {
                    "enabled": "bool"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "agent_pool_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "availability_zones": [
                "list",
                "string"
              ],
              "count": "number",
              "enable_auto_scaling": "bool",
              "enable_node_public_ip": "bool",
              "max_count": "number",
              "max_pods": "number",
              "min_count": "number",
              "name": "string",
              "node_labels": [
                "map",
                "string"
              ],
              "node_public_ip_prefix_id": "string",
              "node_taints": [
                "list",
                "string"
              ],
              "orchestrator_version": "string",
              "os_disk_size_gb": "number",
              "os_type": "string",
              "tags": [
                "map",
                "string"
              ],
              "type": "string",
              "upgrade_settings": [
                "list",
                [
                  "object",
                  {
                    "max_surge": "string"
                  }
                ]
              ],
              "vm_size": "string",
              "vnet_subnet_id": "string",
              "zones": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "api_server_authorized_ip_ranges": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          "string"
        ]
      },
      "azure_active_directory_role_based_access_control": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "admin_group_object_ids": [
                "list",
                "string"
              ],
              "azure_rbac_enabled": "bool",
              "client_app_id": "string",
              "managed": "bool",
              "server_app_id": "string",
              "tenant_id": "string"
            }
          ]
        ]
      },
      "azure_policy_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "disk_encryption_set_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_prefix": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "http_application_routing_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "http_application_routing_zone_name": {
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
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string",
              "user_assigned_identity_id": "string"
            }
          ]
        ]
      },
      "ingress_application_gateway": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "effective_gateway_id": "string",
              "gateway_id": "string",
              "gateway_name": "string",
              "ingress_application_gateway_identity": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "object_id": "string",
                    "user_assigned_identity_id": "string"
                  }
                ]
              ],
              "subnet_cidr": "string",
              "subnet_id": "string"
            }
          ]
        ]
      },
      "key_vault_secrets_provider": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "secret_identity": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "object_id": "string",
                    "user_assigned_identity_id": "string"
                  }
                ]
              ],
              "secret_rotation_enabled": "bool",
              "secret_rotation_interval": "string"
            }
          ]
        ]
      },
      "kube_admin_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_certificate": "string",
              "client_key": "string",
              "cluster_ca_certificate": "string",
              "host": "string",
              "password": "string",
              "username": "string"
            }
          ]
        ]
      },
      "kube_admin_config_raw": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "kube_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_certificate": "string",
              "client_key": "string",
              "cluster_ca_certificate": "string",
              "host": "string",
              "password": "string",
              "username": "string"
            }
          ]
        ]
      },
      "kube_config_raw": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "kubelet_identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_id": "string",
              "object_id": "string",
              "user_assigned_identity_id": "string"
            }
          ]
        ]
      },
      "kubernetes_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "linux_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "admin_username": "string",
              "ssh_key": [
                "list",
                [
                  "object",
                  {
                    "key_data": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dns_service_ip": "string",
              "docker_bridge_cidr": "string",
              "load_balancer_sku": "string",
              "network_plugin": "string",
              "network_policy": "string",
              "pod_cidr": "string",
              "service_cidr": "string"
            }
          ]
        ]
      },
      "node_resource_group": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "oms_agent": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "log_analytics_workspace_id": "string",
              "oms_agent_identity": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "object_id": "string",
                    "user_assigned_identity_id": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "open_service_mesh_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "private_cluster_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "private_fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_link_enabled": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_based_access_control": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "azure_active_directory": [
                "list",
                [
                  "object",
                  {
                    "admin_group_object_ids": [
                      "list",
                      "string"
                    ],
                    "client_app_id": "string",
                    "managed": "bool",
                    "server_app_id": "string",
                    "tenant_id": "string"
                  }
                ]
              ],
              "enabled": "bool"
            }
          ]
        ]
      },
      "role_based_access_control_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "service_principal": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_id": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "windows_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "admin_username": "string"
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
  "version": 2
}`

func AzurermKubernetesClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKubernetesCluster), &result)
	return &result
}
