package data

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
      "domain_configuration_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "filtered_sync_enabled": {
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
      "notifications": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "additional_recipients": [
                "list",
                "string"
              ],
              "notify_dc_admins": "bool",
              "notify_global_admins": "bool"
            }
          ]
        ]
      },
      "replica_sets": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "domain_controller_ip_addresses": [
                "list",
                "string"
              ],
              "external_access_ip_address": "string",
              "id": "string",
              "location": "string",
              "service_status": "string",
              "subnet_id": "string"
            }
          ]
        ]
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
      "secure_ldap": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "certificate_expiry": "string",
              "certificate_thumbprint": "string",
              "enabled": "bool",
              "external_access_enabled": "bool",
              "public_certificate": "string"
            }
          ]
        ]
      },
      "security": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "ntlm_v1_enabled": "bool",
              "sync_kerberos_passwords": "bool",
              "sync_ntlm_passwords": "bool",
              "sync_on_prem_passwords": "bool",
              "tls_v1_enabled": "bool"
            }
          ]
        ]
      },
      "sku": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sync_owner": {
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

func AzurermActiveDirectoryDomainServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermActiveDirectoryDomainService), &result)
	return &result
}
