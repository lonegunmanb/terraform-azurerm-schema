package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVpnServerConfiguration = `{
  "block": {
    "attributes": {
      "azure_active_directory_authentication": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "audience": "string",
              "issuer": "string",
              "tenant": "string"
            }
          ]
        ]
      },
      "client_revoked_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "thumbprint": "string"
            }
          ]
        ]
      },
      "client_root_certificate": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
          [
            "object",
            {
              "name": "string",
              "public_cert_data": "string"
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
      "ipsec_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dh_group": "string",
              "ike_encryption": "string",
              "ike_integrity": "string",
              "ipsec_encryption": "string",
              "ipsec_integrity": "string",
              "pfs_group": "string",
              "sa_data_size_kilobytes": "number",
              "sa_lifetime_seconds": "number"
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
      "radius": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "client_root_certificate": [
                "set",
                [
                  "object",
                  {
                    "name": "string",
                    "thumbprint": "string"
                  }
                ]
              ],
              "server": [
                "list",
                [
                  "object",
                  {
                    "address": "string",
                    "score": "number",
                    "secret": "string"
                  }
                ]
              ],
              "server_root_certificate": [
                "set",
                [
                  "object",
                  {
                    "name": "string",
                    "public_cert_data": "string"
                  }
                ]
              ]
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
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "vpn_authentication_types": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "vpn_protocols": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "set",
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

func AzurermVpnServerConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVpnServerConfiguration), &result)
	return &result
}
