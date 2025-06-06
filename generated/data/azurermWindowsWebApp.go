package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermWindowsWebApp = `{
  "block": {
    "attributes": {
      "app_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "auth_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active_directory": [
                "list",
                [
                  "object",
                  {
                    "allowed_audiences": [
                      "list",
                      "string"
                    ],
                    "client_id": "string",
                    "client_secret": "string",
                    "client_secret_setting_name": "string"
                  }
                ]
              ],
              "additional_login_parameters": [
                "map",
                "string"
              ],
              "allowed_external_redirect_urls": [
                "list",
                "string"
              ],
              "default_provider": "string",
              "enabled": "bool",
              "facebook": [
                "list",
                [
                  "object",
                  {
                    "app_id": "string",
                    "app_secret": "string",
                    "app_secret_setting_name": "string",
                    "oauth_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "github": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "client_secret": "string",
                    "client_secret_setting_name": "string",
                    "oauth_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "google": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "client_secret": "string",
                    "client_secret_setting_name": "string",
                    "oauth_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "issuer": "string",
              "microsoft": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "client_secret": "string",
                    "client_secret_setting_name": "string",
                    "oauth_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "runtime_version": "string",
              "token_refresh_extension_hours": "number",
              "token_store_enabled": "bool",
              "twitter": [
                "list",
                [
                  "object",
                  {
                    "consumer_key": "string",
                    "consumer_secret": "string",
                    "consumer_secret_setting_name": "string"
                  }
                ]
              ],
              "unauthenticated_client_action": "string"
            }
          ]
        ]
      },
      "auth_settings_v2": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active_directory_v2": [
                "list",
                [
                  "object",
                  {
                    "allowed_applications": [
                      "list",
                      "string"
                    ],
                    "allowed_audiences": [
                      "list",
                      "string"
                    ],
                    "allowed_groups": [
                      "list",
                      "string"
                    ],
                    "allowed_identities": [
                      "list",
                      "string"
                    ],
                    "client_id": "string",
                    "client_secret_certificate_thumbprint": "string",
                    "client_secret_setting_name": "string",
                    "jwt_allowed_client_applications": [
                      "list",
                      "string"
                    ],
                    "jwt_allowed_groups": [
                      "list",
                      "string"
                    ],
                    "login_parameters": [
                      "map",
                      "string"
                    ],
                    "tenant_auth_endpoint": "string",
                    "www_authentication_disabled": "bool"
                  }
                ]
              ],
              "apple_v2": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "client_secret_setting_name": "string",
                    "login_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "auth_enabled": "bool",
              "azure_static_web_app_v2": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string"
                  }
                ]
              ],
              "config_file_path": "string",
              "custom_oidc_v2": [
                "list",
                [
                  "object",
                  {
                    "authorisation_endpoint": "string",
                    "certification_uri": "string",
                    "client_credential_method": "string",
                    "client_id": "string",
                    "client_secret_setting_name": "string",
                    "issuer_endpoint": "string",
                    "name": "string",
                    "name_claim_type": "string",
                    "openid_configuration_endpoint": "string",
                    "scopes": [
                      "list",
                      "string"
                    ],
                    "token_endpoint": "string"
                  }
                ]
              ],
              "default_provider": "string",
              "excluded_paths": [
                "list",
                "string"
              ],
              "facebook_v2": [
                "list",
                [
                  "object",
                  {
                    "app_id": "string",
                    "app_secret_setting_name": "string",
                    "graph_api_version": "string",
                    "login_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "forward_proxy_convention": "string",
              "forward_proxy_custom_host_header_name": "string",
              "forward_proxy_custom_scheme_header_name": "string",
              "github_v2": [
                "list",
                [
                  "object",
                  {
                    "client_id": "string",
                    "client_secret_setting_name": "string",
                    "login_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "google_v2": [
                "list",
                [
                  "object",
                  {
                    "allowed_audiences": [
                      "list",
                      "string"
                    ],
                    "client_id": "string",
                    "client_secret_setting_name": "string",
                    "login_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "http_route_api_prefix": "string",
              "login": [
                "list",
                [
                  "object",
                  {
                    "allowed_external_redirect_urls": [
                      "list",
                      "string"
                    ],
                    "cookie_expiration_convention": "string",
                    "cookie_expiration_time": "string",
                    "logout_endpoint": "string",
                    "nonce_expiration_time": "string",
                    "preserve_url_fragments_for_logins": "bool",
                    "token_refresh_extension_time": "number",
                    "token_store_enabled": "bool",
                    "token_store_path": "string",
                    "token_store_sas_setting_name": "string",
                    "validate_nonce": "bool"
                  }
                ]
              ],
              "microsoft_v2": [
                "list",
                [
                  "object",
                  {
                    "allowed_audiences": [
                      "list",
                      "string"
                    ],
                    "client_id": "string",
                    "client_secret_setting_name": "string",
                    "login_scopes": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "require_authentication": "bool",
              "require_https": "bool",
              "runtime_version": "string",
              "twitter_v2": [
                "list",
                [
                  "object",
                  {
                    "consumer_key": "string",
                    "consumer_secret_setting_name": "string"
                  }
                ]
              ],
              "unauthenticated_action": "string"
            }
          ]
        ]
      },
      "backup": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "name": "string",
              "schedule": [
                "list",
                [
                  "object",
                  {
                    "frequency_interval": "number",
                    "frequency_unit": "string",
                    "keep_at_least_one_backup": "bool",
                    "last_execution_time": "string",
                    "retention_period_days": "number",
                    "start_time": "string"
                  }
                ]
              ],
              "storage_account_url": "string"
            }
          ]
        ]
      },
      "client_affinity_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "client_certificate_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "client_certificate_exclusion_paths": {
        "computed": true,
        "description": "Paths to exclude when using client certificates, separated by ;",
        "description_kind": "plain",
        "type": "string"
      },
      "client_certificate_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "type": "string",
              "value": "string"
            }
          ]
        ]
      },
      "custom_domain_verification_id": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "default_hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "ftp_publish_basic_authentication_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "hosting_environment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "https_only": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "kind": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "logs": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "application_logs": [
                "list",
                [
                  "object",
                  {
                    "azure_blob_storage": [
                      "list",
                      [
                        "object",
                        {
                          "level": "string",
                          "retention_in_days": "number",
                          "sas_url": "string"
                        }
                      ]
                    ],
                    "file_system_level": "string"
                  }
                ]
              ],
              "detailed_error_messages": "bool",
              "failed_request_tracing": "bool",
              "http_logs": [
                "list",
                [
                  "object",
                  {
                    "azure_blob_storage": [
                      "list",
                      [
                        "object",
                        {
                          "retention_in_days": "number",
                          "sas_url": "string"
                        }
                      ]
                    ],
                    "file_system": [
                      "list",
                      [
                        "object",
                        {
                          "retention_in_days": "number",
                          "retention_in_mb": "number"
                        }
                      ]
                    ]
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
      "public_network_access_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_plan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "site_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "always_on": "bool",
              "api_definition_url": "string",
              "api_management_api_id": "string",
              "app_command_line": "string",
              "application_stack": [
                "list",
                [
                  "object",
                  {
                    "current_stack": "string",
                    "docker_image_name": "string",
                    "docker_registry_password": "string",
                    "docker_registry_url": "string",
                    "docker_registry_username": "string",
                    "dotnet_core_version": "string",
                    "dotnet_version": "string",
                    "java_container": "string",
                    "java_container_version": "string",
                    "java_embedded_server_enabled": "bool",
                    "java_version": "string",
                    "node_version": "string",
                    "php_version": "string",
                    "python": "bool",
                    "python_version": "string",
                    "tomcat_version": "string"
                  }
                ]
              ],
              "auto_heal_setting": [
                "list",
                [
                  "object",
                  {
                    "action": [
                      "list",
                      [
                        "object",
                        {
                          "action_type": "string",
                          "custom_action": [
                            "list",
                            [
                              "object",
                              {
                                "executable": "string",
                                "parameters": "string"
                              }
                            ]
                          ],
                          "minimum_process_execution_time": "string"
                        }
                      ]
                    ],
                    "trigger": [
                      "list",
                      [
                        "object",
                        {
                          "private_memory_kb": "number",
                          "requests": [
                            "list",
                            [
                              "object",
                              {
                                "count": "number",
                                "interval": "string"
                              }
                            ]
                          ],
                          "slow_request": [
                            "list",
                            [
                              "object",
                              {
                                "count": "number",
                                "interval": "string",
                                "time_taken": "string"
                              }
                            ]
                          ],
                          "slow_request_with_path": [
                            "list",
                            [
                              "object",
                              {
                                "count": "number",
                                "interval": "string",
                                "path": "string",
                                "time_taken": "string"
                              }
                            ]
                          ],
                          "status_code": [
                            "set",
                            [
                              "object",
                              {
                                "count": "number",
                                "interval": "string",
                                "path": "string",
                                "status_code_range": "string",
                                "sub_status": "number",
                                "win32_status_code": "number"
                              }
                            ]
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "container_registry_managed_identity_client_id": "string",
              "container_registry_use_managed_identity": "bool",
              "cors": [
                "list",
                [
                  "object",
                  {
                    "allowed_origins": [
                      "list",
                      "string"
                    ],
                    "support_credentials": "bool"
                  }
                ]
              ],
              "default_documents": [
                "list",
                "string"
              ],
              "detailed_error_logging_enabled": "bool",
              "ftps_state": "string",
              "handler_mapping": [
                "set",
                [
                  "object",
                  {
                    "arguments": "string",
                    "extension": "string",
                    "script_processor_path": "string"
                  }
                ]
              ],
              "health_check_eviction_time_in_min": "number",
              "health_check_path": "string",
              "http2_enabled": "bool",
              "ip_restriction": [
                "list",
                [
                  "object",
                  {
                    "action": "string",
                    "description": "string",
                    "headers": [
                      "list",
                      [
                        "object",
                        {
                          "x_azure_fdid": [
                            "list",
                            "string"
                          ],
                          "x_fd_health_probe": [
                            "list",
                            "string"
                          ],
                          "x_forwarded_for": [
                            "list",
                            "string"
                          ],
                          "x_forwarded_host": [
                            "list",
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
              ],
              "ip_restriction_default_action": "string",
              "load_balancing_mode": "string",
              "local_mysql_enabled": "bool",
              "managed_pipeline_mode": "string",
              "minimum_tls_version": "string",
              "remote_debugging_enabled": "bool",
              "remote_debugging_version": "string",
              "scm_ip_restriction": [
                "list",
                [
                  "object",
                  {
                    "action": "string",
                    "description": "string",
                    "headers": [
                      "list",
                      [
                        "object",
                        {
                          "x_azure_fdid": [
                            "list",
                            "string"
                          ],
                          "x_fd_health_probe": [
                            "list",
                            "string"
                          ],
                          "x_forwarded_for": [
                            "list",
                            "string"
                          ],
                          "x_forwarded_host": [
                            "list",
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
              ],
              "scm_ip_restriction_default_action": "string",
              "scm_minimum_tls_version": "string",
              "scm_type": "string",
              "scm_use_main_ip_restriction": "bool",
              "use_32_bit_worker": "bool",
              "virtual_application": [
                "list",
                [
                  "object",
                  {
                    "physical_path": "string",
                    "preload": "bool",
                    "virtual_directory": [
                      "list",
                      [
                        "object",
                        {
                          "physical_path": "string",
                          "virtual_path": "string"
                        }
                      ]
                    ],
                    "virtual_path": "string"
                  }
                ]
              ],
              "vnet_route_all_enabled": "bool",
              "websockets_enabled": "bool",
              "windows_fx_version": "string",
              "worker_count": "number"
            }
          ]
        ]
      },
      "site_credential": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "password": "string"
            }
          ]
        ]
      },
      "sticky_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "app_setting_names": [
                "list",
                "string"
              ],
              "connection_string_names": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "storage_account": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "access_key": "string",
              "account_name": "string",
              "mount_path": "string",
              "name": "string",
              "share_name": "string",
              "type": "string"
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
      "virtual_network_backup_restore_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "virtual_network_subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "webdeploy_publish_basic_authentication_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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

func AzurermWindowsWebAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermWindowsWebApp), &result)
	return &result
}
