package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKubernetesFluxConfiguration = `{
  "block": {
    "attributes": {
      "cluster_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "continuous_reconciliation_enabled": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "namespace": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "blob_storage": {
        "block": {
          "attributes": {
            "account_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "container_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "local_auth_reference": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sas_token": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "sync_interval_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "managed_identity": {
              "block": {
                "attributes": {
                  "client_id": {
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
            "service_principal": {
              "block": {
                "attributes": {
                  "client_certificate_base64": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_certificate_password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_certificate_send_chain": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "client_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "tenant_id": {
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
      "bucket": {
        "block": {
          "attributes": {
            "access_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "bucket_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "local_auth_reference": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secret_key_base64": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "sync_interval_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "tls_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "url": {
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
      "git_repository": {
        "block": {
          "attributes": {
            "https_ca_cert_base64": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "https_key_base64": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "https_user": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_auth_reference": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "reference_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "reference_value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ssh_known_hosts_base64": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssh_private_key_base64": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "sync_interval_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "url": {
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
      "kustomizations": {
        "block": {
          "attributes": {
            "depends_on": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "garbage_collection_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "recreating_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "retry_interval_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "sync_interval_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "timeout_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
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

func AzurermKubernetesFluxConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKubernetesFluxConfiguration), &result)
	return &result
}
