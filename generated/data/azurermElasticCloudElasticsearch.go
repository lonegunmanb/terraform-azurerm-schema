package data

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
        "computed": true,
        "description_kind": "plain",
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
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "monitoring_enabled": {
        "computed": true,
        "description_kind": "plain",
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
        "computed": true,
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
      }
    },
    "block_types": {
      "logs": {
        "block": {
          "attributes": {
            "filtering_tag": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                [
                  "object",
                  {
                    "action": "string",
                    "name": "string",
                    "value": "string"
                  }
                ]
              ]
            },
            "send_activity_logs": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "send_azuread_logs": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            },
            "send_subscription_logs": {
              "computed": true,
              "description_kind": "plain",
              "type": "bool"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
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

func AzurermElasticCloudElasticsearchSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermElasticCloudElasticsearch), &result)
	return &result
}
