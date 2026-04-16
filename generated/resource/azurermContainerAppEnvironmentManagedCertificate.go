package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppEnvironmentManagedCertificate = `{
  "block": {
    "attributes": {
      "container_app_environment_id": {
        "description": "The Container App Managed Environment ID to configure this Managed Certificate on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "domain_control_validation": {
        "description": "The domain control validation type for the managed certificate. Possible values are ` + "`" + `CNAME` + "`" + ` and ` + "`" + `HTTP` + "`" + `. Defaults to ` + "`" + `HTTP` + "`" + `.",
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
        "description": "The name of the Container Apps Managed Certificate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subject_name": {
        "description": "The Subject Name of the Certificate. Must be a valid domain name.",
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
      "validation_token": {
        "computed": true,
        "description": "The validation token for the managed certificate.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func AzurermContainerAppEnvironmentManagedCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironmentManagedCertificate), &result)
	return &result
}
