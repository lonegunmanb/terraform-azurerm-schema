package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermLogicAppStandard = `{
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
      "bundle_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_affinity_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_certificate_mode": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "ftp_publish_basic_authentication_enabled": {
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
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scm_publish_basic_authentication_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "storage_account_share_name": {
        "computed": true,
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
      "use_extension_bundle": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "virtual_network_subnet_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "vnet_content_share_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
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
              "computed": true,
              "description_kind": "plain",
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
            "public_network_access_enabled": {
              "computed": true,
              "deprecated": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "runtime_scale_monitoring_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scm_min_tls_version": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
                    "description": "Specifies a list of origins that should be allowed to make cross-origin calls.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "support_credentials": {
                    "description": "Are credentials allowed in CORS requests? Defaults to ` + "`" + `false` + "`" + `.",
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
            "ip_restriction": {
              "block": {
                "attributes": {
                  "action": {
                    "description": "The action to take. Possible values are ` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "description": "The description of the IP restriction rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "headers": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
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
                    ]
                  },
                  "ip_address": {
                    "description": "The CIDR notation of the IP or IP Range to match. For example: ` + "`" + `10.0.0.0/24` + "`" + ` or ` + "`" + `192.168.10.1/32` + "`" + ` or ` + "`" + `fe80::/64` + "`" + ` or ` + "`" + `13.107.6.152/31,13.107.128.0/22` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name which should be used for this ` + "`" + `ip_restriction` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description": "The priority value of this ` + "`" + `ip_restriction` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "service_tag": {
                    "description": "The Service Tag used for this IP Restriction.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "virtual_network_subnet_id": {
                    "description": "The Virtual Network Subnet ID used for this IP Restriction.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "scm_ip_restriction": {
              "block": {
                "attributes": {
                  "action": {
                    "description": "The action to take. Possible values are ` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "description": "The description of the IP restriction rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "headers": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
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
                    ]
                  },
                  "ip_address": {
                    "description": "The CIDR notation of the IP or IP Range to match. For example: ` + "`" + `10.0.0.0/24` + "`" + ` or ` + "`" + `192.168.10.1/32` + "`" + ` or ` + "`" + `fe80::/64` + "`" + ` or ` + "`" + `13.107.6.152/31,13.107.128.0/22` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name which should be used for this ` + "`" + `ip_restriction` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description": "The priority value of this ` + "`" + `ip_restriction` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "service_tag": {
                    "description": "The Service Tag used for this IP Restriction.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "virtual_network_subnet_id": {
                    "description": "The Virtual Network Subnet ID used for this IP Restriction.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
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

func AzurermLogicAppStandardSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermLogicAppStandard), &result)
	return &result
}
