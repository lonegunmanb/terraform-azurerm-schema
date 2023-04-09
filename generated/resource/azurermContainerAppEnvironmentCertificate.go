package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppEnvironmentCertificate = `{
  "block": {
    "attributes": {
      "certificate_blob_base64": {
        "description": "The Certificate Private Key as a base64 encoded PFX or PEM.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "certificate_password": {
        "description": "The password for the Certificate.",
        "description_kind": "plain",
        "required": true,
        "sensitive": true,
        "type": "string"
      },
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
        "description": "The name of the Container Apps Environment Certificate.",
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
        "description_kind": "plain",
        "optional": true,
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

func AzurermContainerAppEnvironmentCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironmentCertificate), &result)
	return &result
}
