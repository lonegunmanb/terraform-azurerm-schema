package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAppService = `{
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
      "client_affinity_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_cert_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_cert_mode": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_domain_verification_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "default_site_hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "key_vault_reference_identity_id": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "outbound_ip_address_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "outbound_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "possible_outbound_ip_address_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
      "backup": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_account_url": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "block_types": {
            "schedule": {
              "block": {
                "attributes": {
                  "frequency_interval": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "frequency_unit": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "keep_at_least_one_backup": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "retention_period_in_days": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "start_time": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
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
                "list",
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
      "logs": {
        "block": {
          "attributes": {
            "detailed_error_messages_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "failed_request_tracing_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "application_logs": {
              "block": {
                "attributes": {
                  "file_system_level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "azure_blob_storage": {
                    "block": {
                      "attributes": {
                        "level": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "retention_in_days": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "sas_url": {
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
            "http_logs": {
              "block": {
                "block_types": {
                  "azure_blob_storage": {
                    "block": {
                      "attributes": {
                        "retention_in_days": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "sas_url": {
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
                  "file_system": {
                    "block": {
                      "attributes": {
                        "retention_in_days": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "retention_in_mb": {
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
              "max_items": 1,
              "nesting_mode": "list"
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
            "acr_use_managed_identity_credentials": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "acr_user_managed_identity_client_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "always_on": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "app_command_line": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "auto_swap_slot_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_documents": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "dotnet_framework_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
            "java_container": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "java_container_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
            "local_mysql_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "managed_pipeline_mode": {
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
            "number_of_workers": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "php_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "python_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "remote_debugging_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "remote_debugging_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "windows_fx_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
      "source_control": {
        "block": {
          "attributes": {
            "branch": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "manual_integration": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "repo_url": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rollback_enabled": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_mercurial": {
              "computed": true,
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
      "storage_account": {
        "block": {
          "attributes": {
            "access_key": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "account_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "mount_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "share_name": {
              "description_kind": "plain",
              "required": true,
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

func AzurermAppServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAppService), &result)
	return &result
}
