package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudService = `{
  "block": {
    "attributes": {
      "config_server_git_setting": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "http_basic_auth": [
                "list",
                [
                  "object",
                  {
                    "password": "string",
                    "username": "string"
                  }
                ]
              ],
              "label": "string",
              "repository": [
                "list",
                [
                  "object",
                  {
                    "http_basic_auth": [
                      "list",
                      [
                        "object",
                        {
                          "password": "string",
                          "username": "string"
                        }
                      ]
                    ],
                    "label": "string",
                    "name": "string",
                    "pattern": [
                      "list",
                      "string"
                    ],
                    "search_paths": [
                      "list",
                      "string"
                    ],
                    "ssh_auth": [
                      "list",
                      [
                        "object",
                        {
                          "host_key": "string",
                          "host_key_algorithm": "string",
                          "private_key": "string",
                          "strict_host_key_checking_enabled": "bool"
                        }
                      ]
                    ],
                    "uri": "string"
                  }
                ]
              ],
              "search_paths": [
                "list",
                "string"
              ],
              "ssh_auth": [
                "list",
                [
                  "object",
                  {
                    "host_key": "string",
                    "host_key_algorithm": "string",
                    "private_key": "string",
                    "strict_host_key_checking_enabled": "bool"
                  }
                ]
              ],
              "uri": "string"
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
      "outbound_public_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "required_network_traffic_rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "direction": "string",
              "fqdns": [
                "list",
                "string"
              ],
              "ip_addresses": [
                "list",
                "string"
              ],
              "port": "number",
              "protocol": "string"
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

func AzurermSpringCloudServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudService), &result)
	return &result
}
