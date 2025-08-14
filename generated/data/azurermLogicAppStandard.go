package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLogicAppStandard = `{
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
      "bundle_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_affinity_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
      "name": {
        "description_kind": "plain",
        "required": true,
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
      "public_network_access": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scm_publish_basic_authentication_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
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
              "app_scale_limit": "number",
              "auto_swap_slot_name": "string",
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
              "dotnet_framework_version": "string",
              "elastic_instance_minimum": "number",
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
              "linux_fx_version": "string",
              "min_tls_version": "string",
              "pre_warmed_instance_count": "number",
              "public_network_access_enabled": "bool",
              "runtime_scale_monitoring_enabled": "bool",
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
              "scm_min_tls_version": "string",
              "scm_type": "string",
              "scm_use_main_ip_restriction": "bool",
              "use_32_bit_worker_process": "bool",
              "vnet_route_all_enabled": "bool",
              "websockets_enabled": "bool"
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
      "storage_account_access_key": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "storage_account_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "storage_account_share_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "use_extension_bundle": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "virtual_network_subnet_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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

func AzurermLogicAppStandardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLogicAppStandard), &result)
	return &result
}
