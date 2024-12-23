package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudConfigurationService = `{
  "block": {
    "attributes": {
      "generation": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "refresh_interval_in_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "spring_cloud_service_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "repository": {
        "block": {
          "attributes": {
            "ca_certificate_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_key": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "host_key_algorithm": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "label": {
              "description_kind": "plain",
              "required": true,
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
            "patterns": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "private_key": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "search_paths": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "strict_host_key_checking": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "optional": true,
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
  "version": 1
}`

func AzurermSpringCloudConfigurationServiceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudConfigurationService), &result)
	return &result
}
