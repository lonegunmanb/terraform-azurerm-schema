package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataFactory = `{
  "block": {
    "attributes": {
      "customer_managed_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "customer_managed_key_identity_id": {
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_virtual_network_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "public_network_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "purview_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
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
      }
    },
    "block_types": {
      "github_configuration": {
        "block": {
          "attributes": {
            "account_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "branch_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "git_url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "publishing_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "repository_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "root_folder": {
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
      "global_parameter": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "type": {
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
      },
      "vsts_configuration": {
        "block": {
          "attributes": {
            "account_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "branch_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "project_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "publishing_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "repository_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "root_folder": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tenant_id": {
              "description_kind": "plain",
              "required": true,
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
  "version": 2
}`

func AzurermDataFactorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataFactory), &result)
	return &result
}
