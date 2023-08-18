package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermApiManagement = `{
  "block": {
    "attributes": {
      "additional_location": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity": "number",
              "gateway_regional_url": "string",
              "location": "string",
              "private_ip_addresses": [
                "list",
                "string"
              ],
              "public_ip_address_id": "string",
              "public_ip_addresses": [
                "list",
                "string"
              ],
              "zones": [
                "list",
                "string"
              ]
            }
          ]
        ]
      },
      "developer_portal_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
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
      "hostname_configuration": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "developer_portal": [
                "list",
                [
                  "object",
                  {
                    "host_name": "string",
                    "key_vault_id": "string",
                    "negotiate_client_certificate": "bool"
                  }
                ]
              ],
              "management": [
                "list",
                [
                  "object",
                  {
                    "host_name": "string",
                    "key_vault_id": "string",
                    "negotiate_client_certificate": "bool"
                  }
                ]
              ],
              "portal": [
                "list",
                [
                  "object",
                  {
                    "host_name": "string",
                    "key_vault_id": "string",
                    "negotiate_client_certificate": "bool"
                  }
                ]
              ],
              "proxy": [
                "list",
                [
                  "object",
                  {
                    "default_ssl_binding": "bool",
                    "host_name": "string",
                    "key_vault_id": "string",
                    "negotiate_client_certificate": "bool"
                  }
                ]
              ],
              "scm": [
                "list",
                [
                  "object",
                  {
                    "host_name": "string",
                    "key_vault_id": "string",
                    "negotiate_client_certificate": "bool"
                  }
                ]
              ]
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
              "principal_id": "string",
              "tenant_id": "string",
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
      "management_api_url": {
        "computed": true,
        "description_kind": "plain",
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
        "type": "string"
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
        "computed": true,
        "description_kind": "plain",
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
      "publisher_email": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "publisher_name": {
        "computed": true,
        "description_kind": "plain",
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
      "tenant_access": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "enabled": "bool",
              "primary_key": "string",
              "secondary_key": "string",
              "tenant_id": "string"
            }
          ]
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

func AzurermApiManagementSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermApiManagement), &result)
	return &result
}
