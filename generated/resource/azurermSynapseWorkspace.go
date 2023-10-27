package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSynapseWorkspace = `{
  "block": {
    "attributes": {
      "aad_admin": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "login": "string",
              "object_id": "string",
              "tenant_id": "string"
            }
          ]
        ]
      },
      "azuread_authentication_only": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "compute_subnet_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "connectivity_endpoints": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "data_exfiltration_protection_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "linking_allowed_for_aad_tenant_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_resource_group_name": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
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
      "public_network_access_enabled": {
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
      "sql_aad_admin": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          [
            "object",
            {
              "login": "string",
              "object_id": "string",
              "tenant_id": "string"
            }
          ]
        ]
      },
      "sql_administrator_login": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sql_administrator_login_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "sql_identity_control_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "storage_data_lake_gen2_filesystem_id": {
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
      "azure_devops_repo": {
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
            "last_commit_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "project_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "customer_managed_key": {
        "block": {
          "attributes": {
            "key_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_versionless_id": {
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
      "github_repo": {
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
              "optional": true,
              "type": "string"
            },
            "last_commit_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermSynapseWorkspaceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSynapseWorkspace), &result)
	return &result
}
