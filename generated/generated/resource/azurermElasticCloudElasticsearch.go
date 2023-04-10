package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermElasticCloudElasticsearch = `{
  "block": {
    "attributes": {
      "elastic_cloud_deployment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "elastic_cloud_email_address": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "elastic_cloud_sso_default_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "elastic_cloud_user_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "elasticsearch_service_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "kibana_service_url": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "kibana_sso_uri": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "monitoring_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku_name": {
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
      "logs": {
        "block": {
          "attributes": {
            "send_activity_logs": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "send_azuread_logs": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "send_subscription_logs": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "filtering_tag": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
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

func AzurermElasticCloudElasticsearchSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermElasticCloudElasticsearch), &result)
	return &result
}
