package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermEmailCommunicationServiceDomain = `{
  "block": {
    "attributes": {
      "domain_management": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "email_service_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "from_sender_domain": {
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
      "mail_from_sender_domain": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
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
      },
      "user_engagement_tracking_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "verification_records": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dkim": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "ttl": "number",
                    "type": "string",
                    "value": "string"
                  }
                ]
              ],
              "dkim2": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "ttl": "number",
                    "type": "string",
                    "value": "string"
                  }
                ]
              ],
              "dmarc": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "ttl": "number",
                    "type": "string",
                    "value": "string"
                  }
                ]
              ],
              "domain": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "ttl": "number",
                    "type": "string",
                    "value": "string"
                  }
                ]
              ],
              "spf": [
                "list",
                [
                  "object",
                  {
                    "name": "string",
                    "ttl": "number",
                    "type": "string",
                    "value": "string"
                  }
                ]
              ]
            }
          ]
        ]
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

func AzurermEmailCommunicationServiceDomainSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermEmailCommunicationServiceDomain), &result)
	return &result
}
