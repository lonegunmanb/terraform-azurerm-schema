package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermRoleManagementPolicy = `{
  "block": {
    "attributes": {
      "activation_rules": {
        "computed": true,
        "description": "The activation rules of the policy",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "approval_stage": [
                "list",
                [
                  "object",
                  {
                    "primary_approver": [
                      "set",
                      [
                        "object",
                        {
                          "object_id": "string",
                          "type": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "maximum_duration": "string",
              "require_approval": "bool",
              "require_justification": "bool",
              "require_multifactor_authentication": "bool",
              "require_ticket_info": "bool",
              "required_conditional_access_authentication_context": "string"
            }
          ]
        ]
      },
      "active_assignment_rules": {
        "computed": true,
        "description": "The rules for active assignment of the policy",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "expiration_required": "bool",
              "expire_after": "string",
              "require_justification": "bool",
              "require_multifactor_authentication": "bool",
              "require_ticket_info": "bool"
            }
          ]
        ]
      },
      "description": {
        "computed": true,
        "description": "The Description of the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "eligible_assignment_rules": {
        "computed": true,
        "description": "The rules for eligible assignment of the policy",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "expiration_required": "bool",
              "expire_after": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the policy",
        "description_kind": "plain",
        "type": "string"
      },
      "notification_rules": {
        "computed": true,
        "description": "The notification rules of the policy",
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "active_assignments": [
                "list",
                [
                  "object",
                  {
                    "admin_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ],
                    "approver_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ],
                    "assignee_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "eligible_activations": [
                "list",
                [
                  "object",
                  {
                    "admin_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ],
                    "approver_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ],
                    "assignee_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "eligible_assignments": [
                "list",
                [
                  "object",
                  {
                    "admin_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ],
                    "approver_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ],
                    "assignee_notifications": [
                      "list",
                      [
                        "object",
                        {
                          "additional_recipients": [
                            "set",
                            "string"
                          ],
                          "default_recipients": "bool",
                          "notification_level": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "role_definition_id": {
        "description": "ID of the Azure Role to which this policy is assigned",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scope": {
        "description": "The scope of the role to which this policy will apply",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
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

func AzurermRoleManagementPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermRoleManagementPolicy), &result)
	return &result
}
