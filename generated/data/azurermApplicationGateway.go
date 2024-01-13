package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApplicationGateway = `{
  "block": {
    "attributes": {
      "authentication_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string"
            }
          ]
        ]
      },
      "autoscale_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "max_capacity": "number",
              "min_capacity": "number"
            }
          ]
        ]
      },
      "backend_address_pool": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "fqdns": [
                "list",
                "string"
              ],
              "id": "string",
              "ip_addresses": [
                "list",
                "string"
              ],
              "name": "string"
            }
          ]
        ]
      },
      "backend_http_settings": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "affinity_cookie_name": "string",
              "authentication_certificate": [
                "list",
                [
                  "object",
                  {
                    "id": "string",
                    "name": "string"
                  }
                ]
              ],
              "connection_draining": [
                "list",
                [
                  "object",
                  {
                    "drain_timeout_sec": "number",
                    "enabled": "bool"
                  }
                ]
              ],
              "cookie_based_affinity": "string",
              "host_name": "string",
              "id": "string",
              "name": "string",
              "path": "string",
              "pick_host_name_from_backend_address": "bool",
              "port": "number",
              "probe_id": "string",
              "probe_name": "string",
              "protocol": "string",
              "request_timeout": "number",
              "trusted_root_certificate_names": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "custom_error_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_error_page_url": "string",
              "id": "string",
              "status_code": "string"
            }
          ]
        ]
      },
      "fips_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "firewall_policy_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "force_firewall_policy_association": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "frontend_ip_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "private_ip_address": "string",
              "private_ip_address_allocation": "string",
              "private_link_configuration_id": "string",
              "private_link_configuration_name": "string",
              "public_ip_address_id": "string",
              "subnet_id": "string"
            }
          ]
        ]
      },
      "frontend_port": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "port": "number"
            }
          ]
        ]
      },
      "gateway_ip_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "subnet_id": "string"
            }
          ]
        ]
      },
      "global": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "request_buffering_enabled": "bool",
              "response_buffering_enabled": "bool"
            }
          ]
        ]
      },
      "http2_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "http_listener": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "custom_error_configuration": [
                "list",
                [
                  "object",
                  {
                    "custom_error_page_url": "string",
                    "id": "string",
                    "status_code": "string"
                  }
                ]
              ],
              "firewall_policy_id": "string",
              "frontend_ip_configuration_id": "string",
              "frontend_ip_configuration_name": "string",
              "frontend_port_id": "string",
              "frontend_port_name": "string",
              "host_name": "string",
              "host_names": [
                "list",
                "string"
              ],
              "id": "string",
              "name": "string",
              "protocol": "string",
              "require_sni": "bool",
              "ssl_certificate_id": "string",
              "ssl_certificate_name": "string",
              "ssl_profile_id": "string",
              "ssl_profile_name": "string"
            }
          ]
        ]
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
              "type": "string"
            }
          ]
        ]
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
      "private_endpoint_connection": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string"
            }
          ]
        ]
      },
      "private_link_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "ip_configuration": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "primary": "bool",
                    "private_ip_address": "string",
                    "private_ip_address_allocation": "string",
                    "subnet_id": "string"
                  }
                ]
              ],
              "name": "string"
            }
          ]
        ]
      },
      "probe": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "host": "string",
              "id": "string",
              "interval": "number",
              "match": [
                "list",
                [
                  "object",
                  {
                    "body": "string",
                    "status_code": [
                      "list",
                      "string"
                    ]
                  }
                ]
              ],
              "minimum_servers": "number",
              "name": "string",
              "path": "string",
              "pick_host_name_from_backend_http_settings": "bool",
              "port": "number",
              "protocol": "string",
              "timeout": "number",
              "unhealthy_threshold": "number"
            }
          ]
        ]
      },
      "redirect_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "include_path": "bool",
              "include_query_string": "bool",
              "name": "string",
              "redirect_type": "string",
              "target_listener_id": "string",
              "target_listener_name": "string",
              "target_url": "string"
            }
          ]
        ]
      },
      "request_routing_rule": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "backend_address_pool_id": "string",
              "backend_address_pool_name": "string",
              "backend_http_settings_id": "string",
              "backend_http_settings_name": "string",
              "http_listener_id": "string",
              "http_listener_name": "string",
              "id": "string",
              "name": "string",
              "priority": "number",
              "redirect_configuration_id": "string",
              "redirect_configuration_name": "string",
              "rewrite_rule_set_id": "string",
              "rewrite_rule_set_name": "string",
              "rule_type": "string",
              "url_path_map_id": "string",
              "url_path_map_name": "string"
            }
          ]
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rewrite_rule_set": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "rewrite_rule": [
                "list",
                [
                  "object",
                  {
                    "condition": [
                      "list",
                      [
                        "object",
                        {
                          "ignore_case": "bool",
                          "negate": "bool",
                          "pattern": "string",
                          "variable": "string"
                        }
                      ]
                    ],
                    "name": "string",
                    "request_header_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "header_name": "string",
                          "header_value": "string"
                        }
                      ]
                    ],
                    "response_header_configuration": [
                      "list",
                      [
                        "object",
                        {
                          "header_name": "string",
                          "header_value": "string"
                        }
                      ]
                    ],
                    "rule_sequence": "number",
                    "url": [
                      "list",
                      [
                        "object",
                        {
                          "components": "string",
                          "path": "string",
                          "query_string": "string",
                          "reroute": "bool"
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
      "sku": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity": "number",
              "name": "string",
              "tier": "string"
            }
          ]
        ]
      },
      "ssl_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "key_vault_secret_id": "string",
              "name": "string",
              "public_cert_data": "string"
            }
          ]
        ]
      },
      "ssl_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "cipher_suites": [
                "list",
                "string"
              ],
              "disabled_protocols": [
                "list",
                "string"
              ],
              "min_protocol_version": "string",
              "policy_name": "string",
              "policy_type": "string"
            }
          ]
        ]
      },
      "ssl_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "name": "string",
              "ssl_policy": [
                "list",
                [
                  "object",
                  {
                    "cipher_suites": [
                      "list",
                      "string"
                    ],
                    "disabled_protocols": [
                      "list",
                      "string"
                    ],
                    "min_protocol_version": "string",
                    "policy_name": "string",
                    "policy_type": "string"
                  }
                ]
              ],
              "trusted_client_certificate_names": [
                "list",
                "string"
              ],
              "verify_client_certificate_issuer_dn": "bool",
              "verify_client_certificate_revocation": "string"
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
      "trusted_client_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "data": "string",
              "id": "string",
              "name": "string"
            }
          ]
        ]
      },
      "trusted_root_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "id": "string",
              "key_vault_secret_id": "string",
              "name": "string"
            }
          ]
        ]
      },
      "url_path_map": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "default_backend_address_pool_id": "string",
              "default_backend_address_pool_name": "string",
              "default_backend_http_settings_id": "string",
              "default_backend_http_settings_name": "string",
              "default_redirect_configuration_id": "string",
              "default_redirect_configuration_name": "string",
              "default_rewrite_rule_set_id": "string",
              "default_rewrite_rule_set_name": "string",
              "id": "string",
              "name": "string",
              "path_rule": [
                "list",
                [
                  "object",
                  {
                    "backend_address_pool_id": "string",
                    "backend_address_pool_name": "string",
                    "backend_http_settings_id": "string",
                    "backend_http_settings_name": "string",
                    "firewall_policy_id": "string",
                    "id": "string",
                    "name": "string",
                    "paths": [
                      "list",
                      "string"
                    ],
                    "redirect_configuration_id": "string",
                    "redirect_configuration_name": "string",
                    "rewrite_rule_set_id": "string",
                    "rewrite_rule_set_name": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "waf_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "disabled_rule_group": [
                "list",
                [
                  "object",
                  {
                    "rule_group_name": "string",
                    "rules": [
                      "list",
                      "number"
                    ]
                  }
                ]
              ],
              "enabled": "bool",
              "exclusion": [
                "list",
                [
                  "object",
                  {
                    "match_variable": "string",
                    "selector": "string",
                    "selector_match_operator": "string"
                  }
                ]
              ],
              "file_upload_limit_mb": "number",
              "firewall_mode": "string",
              "max_request_body_size_kb": "number",
              "request_body_check": "bool",
              "rule_set_type": "string",
              "rule_set_version": "string"
            }
          ]
        ]
      },
      "zones": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
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

func AzurermApplicationGatewaySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApplicationGateway), &result)
	return &result
}
