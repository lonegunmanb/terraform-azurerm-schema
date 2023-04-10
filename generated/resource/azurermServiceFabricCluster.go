package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermServiceFabricCluster = `{
  "block": {
    "attributes": {
      "add_on_features": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "cluster_code_version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cluster_endpoint": {
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "management_endpoint": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "reliability_level": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_fabric_zonal_upgrade_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "upgrade_mode": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vm_image": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "vmss_zonal_upgrade_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "azure_active_directory": {
        "block": {
          "attributes": {
            "client_application_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "cluster_application_id": {
              "description_kind": "plain",
              "required": true,
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
      },
      "certificate": {
        "block": {
          "attributes": {
            "thumbprint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "thumbprint_secondary": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "x509_store_name": {
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
      "certificate_common_names": {
        "block": {
          "attributes": {
            "x509_store_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "common_names": {
              "block": {
                "attributes": {
                  "certificate_common_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "certificate_issuer_thumbprint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "client_certificate_common_name": {
        "block": {
          "attributes": {
            "common_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "is_admin": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "issuer_thumbprint": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "client_certificate_thumbprint": {
        "block": {
          "attributes": {
            "is_admin": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "thumbprint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "diagnostics_config": {
        "block": {
          "attributes": {
            "blob_endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protected_account_key_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "queue_endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_account_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "table_endpoint": {
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
      "fabric_settings": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parameters": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "node_type": {
        "block": {
          "attributes": {
            "capacities": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "client_endpoint_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "durability_level": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_endpoint_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "instance_count": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "is_primary": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "is_stateless": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "multiple_availability_zones": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "placement_properties": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "reverse_proxy_endpoint_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "application_ports": {
              "block": {
                "attributes": {
                  "end_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "start_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ephemeral_ports": {
              "block": {
                "attributes": {
                  "end_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "start_port": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "reverse_proxy_certificate": {
        "block": {
          "attributes": {
            "thumbprint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "thumbprint_secondary": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "x509_store_name": {
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
      "reverse_proxy_certificate_common_names": {
        "block": {
          "attributes": {
            "x509_store_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "common_names": {
              "block": {
                "attributes": {
                  "certificate_common_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "certificate_issuer_thumbprint": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "set"
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
      "upgrade_policy": {
        "block": {
          "attributes": {
            "force_restart_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "health_check_retry_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_stable_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_wait_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "upgrade_domain_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "upgrade_replica_set_check_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "upgrade_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "delta_health_policy": {
              "block": {
                "attributes": {
                  "max_delta_unhealthy_applications_percent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_delta_unhealthy_nodes_percent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_upgrade_domain_delta_unhealthy_nodes_percent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "health_policy": {
              "block": {
                "attributes": {
                  "max_unhealthy_applications_percent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "max_unhealthy_nodes_percent": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermServiceFabricClusterSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermServiceFabricCluster), &result)
	return &result
}
