package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppEnvironmentCertificate = `{
  "block": {
    "attributes": {
      "container_app_environment_id": {
        "description": "The Container App Managed Environment ID to configure this Certificate on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "expiration_date": {
        "computed": true,
        "description": "The expiration date for the Certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "issue_date": {
        "computed": true,
        "description": "The date of issue for the Certificate.",
        "description_kind": "plain",
        "type": "string"
      },
      "issuer": {
        "computed": true,
        "description": "The Certificate Issuer.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the Container Apps Certificate.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subject_name": {
        "computed": true,
        "description": "The Subject Name for the Certificate.",
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
      "thumbprint": {
        "computed": true,
        "description": "The Thumbprint of the Certificate.",
        "description_kind": "plain",
        "type": "string"
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

func AzurermContainerAppEnvironmentCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironmentCertificate), &result)
	return &result
}
