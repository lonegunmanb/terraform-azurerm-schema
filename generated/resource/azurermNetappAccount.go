package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappAccount = `{
  "block": {
    "attributes": {
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
      "active_directory": {
        "block": {
          "attributes": {
            "aes_encryption_enabled": {
              "description": "If enabled, AES encryption will be enabled for SMB communication.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "dns_servers": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "domain": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "kerberos_ad_name": {
              "description": "Name of the active directory machine. This optional parameter is used only while creating kerberos volume.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "kerberos_kdc_ip": {
              "description": "IP address of the KDC server (usually same the DC). This optional parameter is used only while creating kerberos volume.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ldap_over_tls_enabled": {
              "description": "Specifies whether or not the LDAP traffic needs to be secured via TLS.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ldap_signing_enabled": {
              "description": "Specifies whether or not the LDAP traffic needs to be signed.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "local_nfs_users_with_ldap_allowed": {
              "description": "If enabled, NFS client local users can also (in addition to LDAP users) access the NFS volumes.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "organizational_unit": {
              "description": "The Organizational Unit (OU) within the Windows Active Directory where machines will be created. If blank, defaults to 'CN=Computers'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "server_root_ca_certificate": {
              "description": "When LDAP over SSL/TLS is enabled, the LDAP client is required to have base64 encoded Active Directory Certificate Service's self-signed root CA certificate, this optional parameter is used only for dual protocol with LDAP user-mapping volumes.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "site_name": {
              "description": "The Active Directory site the service will limit Domain Controller discovery to. If blank, defaults to 'Default-First-Site-Name'",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "smb_server_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "username": {
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

func AzurermNetappAccountSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappAccount), &result)
	return &result
}
