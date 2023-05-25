package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAppService = `{
  "block": {
    "attributes": {
      "app_service_plan_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "app_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "client_affinity_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "client_cert_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "connection_string": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
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
        "type": "string"
      },
      "default_site_hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "site_config": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "acr_use_managed_identity_credentials": "bool",
              "acr_user_managed_identity_client_id": "string",
              "always_on": "bool",
              "app_command_line": "string",
              "cors": [
                "list",
                [
                  "object",
                  {
                    "allowed_origins": [
                      "set",
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
              "dotnet_framework_version": "string",
              "ftps_state": "string",
              "health_check_path": "string",
              "http2_enabled": "bool",
              "ip_restriction": [
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
              ],
              "java_container": "string",
              "java_container_version": "string",
              "java_version": "string",
              "linux_fx_version": "string",
              "local_mysql_enabled": "bool",
              "managed_pipeline_mode": "string",
              "min_tls_version": "string",
              "number_of_workers": "number",
              "php_version": "string",
              "python_version": "string",
              "remote_debugging_enabled": "bool",
              "remote_debugging_version": "string",
              "scm_ip_restriction": [
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
              ],
              "scm_type": "string",
              "scm_use_main_ip_restriction": "bool",
              "use_32_bit_worker_process": "bool",
              "vnet_route_all_enabled": "bool",
              "websockets_enabled": "bool",
              "windows_fx_version": "string"
            }
          ]
        ]
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
      "source_control": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "branch": "string",
              "manual_integration": "bool",
              "repo_url": "string",
              "rollback_enabled": "bool",
              "use_mercurial": "bool"
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

func AzurermAppServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAppService), &result)
	return &result
}
