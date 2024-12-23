package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataFactoryIntegrationRuntimeAzureSsis = `{
  "block": {
    "attributes": {
      "credential_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_factory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "edition": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "license_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "max_parallel_executions_per_node": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "node_size": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "number_of_nodes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "catalog_info": {
        "block": {
          "attributes": {
            "administrator_login": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "administrator_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "dual_standby_pair_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "elastic_pool_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pricing_tier": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server_endpoint": {
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
      "copy_compute_scale": {
        "block": {
          "attributes": {
            "data_integration_unit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "time_to_live": {
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
      "custom_setup_script": {
        "block": {
          "attributes": {
            "blob_container_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sas_token": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "express_custom_setup": {
        "block": {
          "attributes": {
            "environment": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "powershell_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "command_key": {
              "block": {
                "attributes": {
                  "password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "target_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "user_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "key_vault_password": {
                    "block": {
                      "attributes": {
                        "linked_service_name": {
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
                        },
                        "secret_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_version": {
                          "description_kind": "plain",
                          "optional": true,
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
              "nesting_mode": "list"
            },
            "component": {
              "block": {
                "attributes": {
                  "license": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "key_vault_license": {
                    "block": {
                      "attributes": {
                        "linked_service_name": {
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
                        },
                        "secret_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_version": {
                          "description_kind": "plain",
                          "optional": true,
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "express_vnet_integration": {
        "block": {
          "attributes": {
            "subnet_id": {
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
      "package_store": {
        "block": {
          "attributes": {
            "linked_service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "pipeline_external_compute_scale": {
        "block": {
          "attributes": {
            "number_of_external_nodes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "number_of_pipeline_nodes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "time_to_live": {
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
      "proxy": {
        "block": {
          "attributes": {
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "self_hosted_integration_runtime_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "staging_storage_linked_service_name": {
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
      "vnet_integration": {
        "block": {
          "attributes": {
            "public_ips": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "subnet_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "vnet_id": {
              "description_kind": "plain",
              "optional": true,
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
  "version": 0
}`

func AzurermDataFactoryIntegrationRuntimeAzureSsisSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataFactoryIntegrationRuntimeAzureSsis), &result)
	return &result
}
