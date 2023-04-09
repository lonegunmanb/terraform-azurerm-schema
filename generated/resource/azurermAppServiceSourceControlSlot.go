package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAppServiceSourceControlSlot = `{
  "block": {
    "attributes": {
      "branch": {
        "computed": true,
        "description": "The URL for the repository",
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
      "repo_url": {
        "computed": true,
        "description": "The branch name to use for deployments.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "rollback_enabled": {
        "description": "Should the Deployment Rollback be enabled? Defaults to ` + "`" + `false` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "scm_type": {
        "computed": true,
        "description": "The SCM Type in use. This value is decoded by the service from the repository information supplied.",
        "description_kind": "plain",
        "type": "string"
      },
      "slot_id": {
        "description": "The ID of the Linux or Windows Web App Slot.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "use_local_git": {
        "description": "Should the Slot use local Git configuration.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "use_manual_integration": {
        "description": "Should code be deployed manually. Set to ` + "`" + `true` + "`" + ` to disable continuous integration, such as webhooks into online repos such as GitHub. Defaults to ` + "`" + `false` + "`" + `",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "use_mercurial": {
        "description": "The repository specified is Mercurial. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "uses_github_action": {
        "computed": true,
        "description": "Indicates if the Slot uses a GitHub action for deployment. This value is decoded by the service from the repository information supplied.",
        "description_kind": "plain",
        "type": "bool"
      }
    },
    "block_types": {
      "github_action_configuration": {
        "block": {
          "attributes": {
            "generate_workflow_file": {
              "description": "Should the service generate the GitHub Action Workflow file. Defaults to ` + "`" + `true` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "linux_action": {
              "computed": true,
              "description": "Denotes this action uses a Linux base image.",
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "block_types": {
            "code_configuration": {
              "block": {
                "attributes": {
                  "runtime_stack": {
                    "description": "The value to use for the Runtime Stack in the workflow file content for code base apps.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "runtime_version": {
                    "description": "The value to use for the Runtime Version in the workflow file content for code base apps.",
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
            "container_configuration": {
              "block": {
                "attributes": {
                  "image_name": {
                    "description": "The image name for the build.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "registry_password": {
                    "description": "The password used to upload the image to the container registry.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "registry_url": {
                    "description": "The server URL for the container registry where the build will be hosted.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "registry_username": {
                    "description": "The username used to upload the image to the container registry.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
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

func AzurermAppServiceSourceControlSlotSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAppServiceSourceControlSlot), &result)
	return &result
}
