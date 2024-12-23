package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermSpringCloudCustomizedAccelerator = `{
  "block": {
    "attributes": {
      "accelerator_tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "accelerator_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "display_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "icon_url": {
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
      "spring_cloud_accelerator_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "git_repository": {
        "block": {
          "attributes": {
            "branch": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "ca_certificate_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "commit": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "git_tag": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "interval_in_seconds": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "url": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "basic_auth": {
              "block": {
                "attributes": {
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "username": {
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
            "ssh_auth": {
              "block": {
                "attributes": {
                  "host_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "host_key_algorithm": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "private_key": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
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
        "min_items": 1,
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

func AzurermSpringCloudCustomizedAcceleratorSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermSpringCloudCustomizedAccelerator), &result)
	return &result
}
