package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppEnvironmentDaprComponent = `{
  "block": {
    "attributes": {
      "component_type": {
        "description": "The Dapr Component Type. For example ` + "`" + `state.azure.blobstorage` + "`" + `.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "container_app_environment_id": {
        "description": "The Container App Managed Environment ID to configure this Dapr component on.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "ignore_errors": {
        "description": "Should the Dapr sidecar to continue initialisation if the component fails to load. Defaults to ` + "`" + `false` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "init_timeout": {
        "description": "The component initialisation timeout in ISO8601 format. e.g. ` + "`" + `5s` + "`" + `, ` + "`" + `2h` + "`" + `, ` + "`" + `1m` + "`" + `. Defaults to ` + "`" + `5s` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description": "The name for this Dapr Component.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scopes": {
        "description": "A list of scopes to which this component applies. e.g. a Container App's ` + "`" + `dapr.app_id` + "`" + ` value.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "version": {
        "description": "The version of the component.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "metadata": {
        "block": {
          "attributes": {
            "name": {
              "description": "The name of the Metadata configuration item.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secret_name": {
              "description": "The name of a secret specified in the ` + "`" + `secrets` + "`" + ` block that contains the value for this metadata configuration item.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "value": {
              "description": "The value for this metadata configuration item.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "secret": {
        "block": {
          "attributes": {
            "identity": {
              "description": "The identity to use for accessing key vault reference.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_secret_id": {
              "description": "The Key Vault Secret ID. Could be either one of ` + "`" + `id` + "`" + ` or ` + "`" + `versionless_id` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "The secret name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for this secret.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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

func AzurermContainerAppEnvironmentDaprComponentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironmentDaprComponent), &result)
	return &result
}
