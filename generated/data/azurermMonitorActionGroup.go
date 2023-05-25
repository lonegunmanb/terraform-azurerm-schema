package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMonitorActionGroup = `{
  "block": {
    "attributes": {
      "arm_role_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "role_id": "string",
              "use_common_alert_schema": "bool"
            }
          ]
        ]
      },
      "automation_runbook_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automation_account_id": "string",
              "is_global_runbook": "bool",
              "name": "string",
              "runbook_name": "string",
              "service_uri": "string",
              "use_common_alert_schema": "bool",
              "webhook_resource_id": "string"
            }
          ]
        ]
      },
      "azure_app_push_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "email_address": "string",
              "name": "string"
            }
          ]
        ]
      },
      "azure_function_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "function_app_resource_id": "string",
              "function_name": "string",
              "http_trigger_url": "string",
              "name": "string",
              "use_common_alert_schema": "bool"
            }
          ]
        ]
      },
      "email_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "email_address": "string",
              "name": "string",
              "use_common_alert_schema": "bool"
            }
          ]
        ]
      },
      "enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "itsm_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "connection_id": "string",
              "name": "string",
              "region": "string",
              "ticket_configuration": "string",
              "workspace_id": "string"
            }
          ]
        ]
      },
      "logic_app_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "callback_url": "string",
              "name": "string",
              "resource_id": "string",
              "use_common_alert_schema": "bool"
            }
          ]
        ]
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
      "short_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "sms_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "country_code": "string",
              "name": "string",
              "phone_number": "string"
            }
          ]
        ]
      },
      "voice_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "country_code": "string",
              "name": "string",
              "phone_number": "string"
            }
          ]
        ]
      },
      "webhook_receiver": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "aad_auth": [
                "list",
                [
                  "object",
                  {
                    "identifier_uri": "string",
                    "object_id": "string",
                    "tenant_id": "string"
                  }
                ]
              ],
              "name": "string",
              "service_uri": "string",
              "use_common_alert_schema": "bool"
            }
          ]
        ]
      }
    },
    "block_types": {
      "event_hub_receiver": {
        "block": {
          "attributes": {
            "event_hub_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "use_common_alert_schema": {
              "description_kind": "plain",
              "optional": true,
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
  "version": 1
}`

func AzurermMonitorActionGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMonitorActionGroup), &result)
	return &result
}
