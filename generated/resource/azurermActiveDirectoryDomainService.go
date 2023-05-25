package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermActiveDirectoryDomainService = `{
  "block": {
    "attributes": {
      "deployment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "filtered_sync_enabled": {
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sku": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sync_owner": {
        "computed": true,
        "description_kind": "plain",
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
      "tenant_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      }
    },
    "block_types": {
      "initial_replica_set": {
        "block": {
          "attributes": {
            "domain_controller_ip_addresses": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "external_access_ip_address": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "location": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "service_status": {
              "computed": true,
              "description_kind": "plain",
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
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "notifications": {
        "block": {
          "attributes": {
            "additional_recipients": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "notify_dc_admins": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "notify_global_admins": {
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
      "secure_ldap": {
        "block": {
          "attributes": {
            "certificate_expiry": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "certificate_thumbprint": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "external_access_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "pfx_certificate": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "pfx_certificate_password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "public_certificate": {
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
      "security": {
        "block": {
          "attributes": {
            "ntlm_v1_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sync_kerberos_passwords": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sync_ntlm_passwords": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "sync_on_prem_passwords": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "tls_v1_enabled": {
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

func AzurermActiveDirectoryDomainServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermActiveDirectoryDomainService), &result)
	return &result
}
