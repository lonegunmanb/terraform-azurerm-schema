package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApiManagementCustomDomain = `{
  "block": {
    "attributes": {
      "api_management_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
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
            "key_vault_certificate_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_id": {
              "computed": true,
              "deprecated": true,
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
      "gateway": {
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
            "key_vault_certificate_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_id": {
              "computed": true,
              "deprecated": true,
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
            "key_vault_certificate_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_id": {
              "computed": true,
              "deprecated": true,
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
            "key_vault_certificate_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_id": {
              "computed": true,
              "deprecated": true,
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
            "key_vault_certificate_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_id": {
              "computed": true,
              "deprecated": true,
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

func AzurermApiManagementCustomDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApiManagementCustomDomain), &result)
	return &result
}
