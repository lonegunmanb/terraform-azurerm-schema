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
      "agent_pool_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "auto_scaling_enabled": "bool",
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
              "node_public_ip_enabled": "bool",
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
                    "drain_timeout_in_minutes": "number",
                    "max_surge": "string",
                    "node_soak_duration_in_minutes": "number"
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
      "current_kubernetes_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_ca_trust_certificates_base64": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
      "key_management_service": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "key_vault_key_id": "string",
              "key_vault_network_access": "string"
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
        "sensitive": true,
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
        "sensitive": true,
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
      "microsoft_defender": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "log_analytics_workspace_id": "string"
            }
          ]
        ]
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
      "node_resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "oidc_issuer_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "oidc_issuer_url": {
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
              "msi_auth_for_monitoring_enabled": "bool",
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "role_based_access_control_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "service_mesh_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "certificate_authority": [
                "list",
                [
                  "object",
                  {
                    "cert_chain_object_name": "string",
                    "cert_object_name": "string",
                    "key_object_name": "string",
                    "key_vault_id": "string",
                    "root_cert_object_name": "string"
                  }
                ]
              ],
              "external_ingress_gateway_enabled": "bool",
              "internal_ingress_gateway_enabled": "bool",
              "mode": "string"
            }
          ]
        ]
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
      "storage_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "blob_driver_enabled": "bool",
              "disk_driver_enabled": "bool",
              "disk_driver_version": "string",
              "file_driver_enabled": "bool",
              "snapshot_controller_enabled": "bool"
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
  "version": 0
}`

func AzurermKubernetesClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKubernetesCluster), &result)
	return &result
}
