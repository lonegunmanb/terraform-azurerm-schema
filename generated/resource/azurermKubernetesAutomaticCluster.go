package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKubernetesAutomaticCluster = `{
  "block": {
    "attributes": {
      "current_kubernetes_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "fully_qualified_domain_name": {
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
      "node_resource_group_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "oidc_issuer_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "portal_fully_qualified_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_fully_qualified_domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "api_server_access": {
        "block": {
          "attributes": {
            "authorized_ip_ranges": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "hosted_system": {
        "block": {
          "attributes": {
            "node_subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "system_node_subnet_id": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "private_cluster": {
        "block": {
          "attributes": {
            "private_dns_zone_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_fully_qualified_domain_name_enabled": {
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
      "service_mesh": {
        "block": {
          "attributes": {
            "external_ingress_gateway_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "internal_ingress_gateway_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "proxy_redirect_mechanism": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "revisions": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "certificate_authority": {
              "block": {
                "attributes": {
                  "certificate_chain_object_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "certificate_object_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_object_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_vault_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "root_certificate_object_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
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
      },
      "web_app_routing_ingress": {
        "block": {
          "attributes": {
            "default_nginx_controller": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dns_zone_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "istio_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "web_app_routing_identity": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermKubernetesAutomaticClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKubernetesAutomaticCluster), &result)
	return &result
}
