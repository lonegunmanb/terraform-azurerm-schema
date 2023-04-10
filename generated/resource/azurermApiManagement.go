package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApiManagement = `{
  "block": {
    "attributes": {
      "client_certificate_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "developer_portal_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_disabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "gateway_regional_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "gateway_url": {
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
      "management_api_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "min_api_version": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "notification_sender_email": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "xml_content": "string",
              "xml_link": "string"
            }
          ]
        ]
      },
      "portal_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "public_ip_address_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "public_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "publisher_email": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "publisher_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scm_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sku_name": {
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
      "virtual_network_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
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
      "additional_location": {
        "block": {
          "attributes": {
            "capacity": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "gateway_disabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "gateway_regional_url": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "location": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "private_ip_addresses": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "public_ip_address_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "public_ip_addresses": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
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
            "virtual_network_configuration": {
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "certificate": {
        "block": {
          "attributes": {
            "certificate_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "encoded_certificate": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "expiry": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "store_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "subject": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "thumbprint": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 10,
        "nesting_mode": "list"
      },
      "delegation": {
        "block": {
          "attributes": {
            "subscriptions_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "user_registration_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "validation_key": {
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
      "hostname_configuration": {
        "block": {
          "block_types": {
            "developer_portal": {
              "block": {
                "attributes": {
                  "certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_source": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "certificate_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expiry": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_vault_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "negotiate_client_certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ssl_keyvault_identity_client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "thumbprint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "management": {
              "block": {
                "attributes": {
                  "certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_source": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "certificate_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expiry": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_vault_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "negotiate_client_certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ssl_keyvault_identity_client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "thumbprint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "portal": {
              "block": {
                "attributes": {
                  "certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_source": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "certificate_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expiry": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_vault_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "negotiate_client_certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ssl_keyvault_identity_client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "thumbprint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "proxy": {
              "block": {
                "attributes": {
                  "certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_source": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "certificate_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "default_ssl_binding": {
                    "computed": true,
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "expiry": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_vault_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "negotiate_client_certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ssl_keyvault_identity_client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "thumbprint": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "scm": {
              "block": {
                "attributes": {
                  "certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "certificate_source": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "certificate_status": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "expiry": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "host_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "key_vault_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "negotiate_client_certificate": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "ssl_keyvault_identity_client_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subject": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "thumbprint": {
                    "computed": true,
                    "description_kind": "plain",
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
      "protocols": {
        "block": {
          "attributes": {
            "enable_http2": {
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
      "security": {
        "block": {
          "attributes": {
            "enable_backend_ssl30": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_backend_tls10": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_backend_tls11": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_frontend_ssl30": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_frontend_tls10": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "enable_frontend_tls11": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_ecdhe_ecdsa_with_aes128_cbc_sha_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_ecdhe_ecdsa_with_aes256_cbc_sha_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_ecdhe_rsa_with_aes128_cbc_sha_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_ecdhe_rsa_with_aes256_cbc_sha_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_rsa_with_aes128_cbc_sha256_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_rsa_with_aes128_cbc_sha_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_rsa_with_aes128_gcm_sha256_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_rsa_with_aes256_cbc_sha256_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_rsa_with_aes256_cbc_sha_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_rsa_with_aes256_gcm_sha384_ciphers_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "triple_des_ciphers_enabled": {
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
      "sign_in": {
        "block": {
          "attributes": {
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
      },
      "sign_up": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            }
          },
          "block_types": {
            "terms_of_service": {
              "block": {
                "attributes": {
                  "consent_required": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "text": {
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
      "tenant_access": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "primary_key": {
              "computed": true,
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            },
            "secondary_key": {
              "computed": true,
              "description_kind": "plain",
              "sensitive": true,
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
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
      "virtual_network_configuration": {
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermApiManagementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApiManagement), &result)
	return &result
}
