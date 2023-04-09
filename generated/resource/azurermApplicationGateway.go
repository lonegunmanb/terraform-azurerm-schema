package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApplicationGateway = `{
  "block": {
    "attributes": {
      "enable_http2": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "fips_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "firewall_policy_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "force_firewall_policy_association": {
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
      "private_endpoint_connection": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "id": "string",
              "name": "string"
            }
          ]
        ]
      },
      "resource_group_name": {
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
      "zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "authentication_certificate": {
        "block": {
          "attributes": {
            "data": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
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
      "autoscale_configuration": {
        "block": {
          "attributes": {
            "max_capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_capacity": {
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
      "backend_address_pool": {
        "block": {
          "attributes": {
            "fqdns": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ip_addresses": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "backend_http_settings": {
        "block": {
          "attributes": {
            "affinity_cookie_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cookie_based_affinity": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "host_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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
            "pick_host_name_from_backend_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "probe_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "probe_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "request_timeout": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "trusted_root_certificate_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "block_types": {
            "authentication_certificate": {
              "block": {
                "attributes": {
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
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
            "connection_draining": {
              "block": {
                "attributes": {
                  "drain_timeout_sec": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
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
        "min_items": 1,
        "nesting_mode": "set"
      },
      "custom_error_configuration": {
        "block": {
          "attributes": {
            "custom_error_page_url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "status_code": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "frontend_ip_configuration": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "private_ip_address": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_ip_address_allocation": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "private_link_configuration_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "private_link_configuration_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_ip_address_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "subnet_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "frontend_port": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "gateway_ip_configuration": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 2,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "global": {
        "block": {
          "attributes": {
            "request_buffering_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "response_buffering_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "http_listener": {
        "block": {
          "attributes": {
            "firewall_policy_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "frontend_ip_configuration_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "frontend_ip_configuration_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "frontend_port_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "frontend_port_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "host_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "require_sni": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ssl_certificate_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_certificate_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ssl_profile_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "ssl_profile_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "custom_error_configuration": {
              "block": {
                "attributes": {
                  "custom_error_page_url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "status_code": {
                    "description_kind": "plain",
                    "required": true,
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
        "min_items": 1,
        "nesting_mode": "set"
      },
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
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
      "private_link_configuration": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "ip_configuration": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "primary": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "private_ip_address": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "private_ip_address_allocation": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "probe": {
        "block": {
          "attributes": {
            "host": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "interval": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "minimum_servers": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "pick_host_name_from_backend_http_settings": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "timeout": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "unhealthy_threshold": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "match": {
              "block": {
                "attributes": {
                  "body": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "status_code": {
                    "description_kind": "plain",
                    "required": true,
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "redirect_configuration": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "include_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "include_query_string": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "redirect_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "target_listener_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "target_listener_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "target_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "request_routing_rule": {
        "block": {
          "attributes": {
            "backend_address_pool_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "backend_address_pool_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "backend_http_settings_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "backend_http_settings_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_listener_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "http_listener_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "redirect_configuration_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "redirect_configuration_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rewrite_rule_set_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rewrite_rule_set_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "url_path_map_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "url_path_map_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
      },
      "rewrite_rule_set": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "rewrite_rule": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rule_sequence": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "condition": {
                    "block": {
                      "attributes": {
                        "ignore_case": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "negate": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "pattern": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "variable": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "request_header_configuration": {
                    "block": {
                      "attributes": {
                        "header_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "header_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "response_header_configuration": {
                    "block": {
                      "attributes": {
                        "header_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "header_value": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "url": {
                    "block": {
                      "attributes": {
                        "components": {
                          "computed": true,
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "path": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "query_string": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "reroute": {
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
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "sku": {
        "block": {
          "attributes": {
            "capacity": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tier": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "ssl_certificate": {
        "block": {
          "attributes": {
            "data": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_vault_secret_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "public_cert_data": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "ssl_policy": {
        "block": {
          "attributes": {
            "cipher_suites": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "disabled_protocols": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "min_protocol_version": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "policy_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "ssl_profile": {
        "block": {
          "attributes": {
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "trusted_client_certificate_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "verify_client_cert_issuer_dn": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "ssl_policy": {
              "block": {
                "attributes": {
                  "cipher_suites": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "disabled_protocols": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "min_protocol_version": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "policy_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "policy_type": {
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
      "trusted_client_certificate": {
        "block": {
          "attributes": {
            "data": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
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
      "trusted_root_certificate": {
        "block": {
          "attributes": {
            "data": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_vault_secret_id": {
              "description_kind": "plain",
              "optional": true,
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
      "url_path_map": {
        "block": {
          "attributes": {
            "default_backend_address_pool_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "default_backend_address_pool_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_backend_http_settings_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "default_backend_http_settings_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_redirect_configuration_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "default_redirect_configuration_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_rewrite_rule_set_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "default_rewrite_rule_set_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "path_rule": {
              "block": {
                "attributes": {
                  "backend_address_pool_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "backend_address_pool_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "backend_http_settings_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "backend_http_settings_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "firewall_policy_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "paths": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "redirect_configuration_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "redirect_configuration_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rewrite_rule_set_id": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "rewrite_rule_set_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "waf_configuration": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "file_upload_limit_mb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "firewall_mode": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "max_request_body_size_kb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "request_body_check": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "rule_set_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rule_set_version": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "disabled_rule_group": {
              "block": {
                "attributes": {
                  "rule_group_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "rules": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "exclusion": {
              "block": {
                "attributes": {
                  "match_variable": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "selector": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "selector_match_operator": {
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
