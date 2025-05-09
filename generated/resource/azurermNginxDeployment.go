package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNginxDeployment = `{
  "block": {
    "attributes": {
      "automatic_upgrade_channel": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "capacity": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "dataplane_api_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "diagnose_support_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "email": {
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
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "managed_resource_group": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "nginx_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku": {
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
      "auto_scale_profile": {
        "block": {
          "attributes": {
            "max_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "min_capacity": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "frontend_private": {
        "block": {
          "attributes": {
            "allocation_method": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ip_address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "frontend_public": {
        "block": {
          "attributes": {
            "ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
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
      "logging_storage_account": {
        "block": {
          "attributes": {
            "container_name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "network_interface": {
        "block": {
          "attributes": {
            "subnet_id": {
              "description_kind": "plain",
              "required": true,
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
      },
      "web_application_firewall": {
        "block": {
          "attributes": {
            "activation_state_enabled": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "status": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "attack_signatures_package": [
                      "list",
                      [
                        "object",
                        {
                          "revision_datetime": "string",
                          "version": "string"
                        }
                      ]
                    ],
                    "bot_signatures_package": [
                      "list",
                      [
                        "object",
                        {
                          "revision_datetime": "string",
                          "version": "string"
                        }
                      ]
                    ],
                    "component_versions": [
                      "list",
                      [
                        "object",
                        {
                          "waf_engine_version": "string",
                          "waf_nginx_version": "string"
                        }
                      ]
                    ],
                    "threat_campaigns_package": [
                      "list",
                      [
                        "object",
                        {
                          "revision_datetime": "string",
                          "version": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
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
  "version": 0
}`

func AzurermNginxDeploymentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNginxDeployment), &result)
	return &result
}
