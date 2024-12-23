package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermFunctionAppSlot = `{
  "block": {
    "attributes": {
      "app_service_plan_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "app_settings": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "daily_memory_time_quota": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "default_hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enable_builtin_logging": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "function_app_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "https_only": {
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
      "kind": {
        "computed": true,
        "description_kind": "plain",
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
      "os_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "outbound_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "possible_outbound_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_credential": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "password": "string",
              "username": "string"
            }
          ]
        ]
      },
      "storage_account_access_key": {
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
      "storage_account_name": {
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
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "auth_settings": {
        "block": {
          "attributes": {
            "additional_login_params": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "allowed_external_redirect_urls": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "default_provider": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "issuer": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "runtime_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "token_refresh_extension_hours": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "token_store_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "unauthenticated_client_action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "active_directory": {
              "block": {
                "attributes": {
                  "allowed_audiences": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "facebook": {
              "block": {
                "attributes": {
                  "app_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "app_secret": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "google": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "microsoft": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "twitter": {
              "block": {
                "attributes": {
                  "consumer_key": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "consumer_secret": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "connection_string": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
        "nesting_mode": "list"
      },
      "site_config": {
        "block": {
          "attributes": {
            "always_on": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "app_scale_limit": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "auto_swap_slot_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "dotnet_framework_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "elastic_instance_minimum": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "ftps_state": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "health_check_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http2_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ip_restriction": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "action": "string",
                    "headers": [
                      "list",
                      [
                        "object",
                        {
                          "x_azure_fdid": [
                            "set",
                            "string"
                          ],
                          "x_fd_health_probe": [
                            "set",
                            "string"
                          ],
                          "x_forwarded_for": [
                            "set",
                            "string"
                          ],
                          "x_forwarded_host": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "ip_address": "string",
                    "name": "string",
                    "priority": "number",
                    "service_tag": "string",
                    "virtual_network_subnet_id": "string"
                  }
                ]
              ]
            },
            "java_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "linux_fx_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "min_tls_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "pre_warmed_instance_count": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "runtime_scale_monitoring_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scm_ip_restriction": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                [
                  "object",
                  {
                    "action": "string",
                    "headers": [
                      "list",
                      [
                        "object",
                        {
                          "x_azure_fdid": [
                            "set",
                            "string"
                          ],
                          "x_fd_health_probe": [
                            "set",
                            "string"
                          ],
                          "x_forwarded_for": [
                            "set",
                            "string"
                          ],
                          "x_forwarded_host": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "ip_address": "string",
                    "name": "string",
                    "priority": "number",
                    "service_tag": "string",
                    "virtual_network_subnet_id": "string"
                  }
                ]
              ]
            },
            "scm_type": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scm_use_main_ip_restriction": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_32_bit_worker_process": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "vnet_route_all_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "websockets_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "cors": {
              "block": {
                "attributes": {
                  "allowed_origins": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "support_credentials": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      }
    },
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermFunctionAppSlotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermFunctionAppSlot), &result)
	return &result
}
