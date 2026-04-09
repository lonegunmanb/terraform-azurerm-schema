package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermManagedDevopsPool = `{
  "block": {
    "attributes": {
      "azure_devops_organization": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "organization": [
                "list",
                [
                  "object",
                  {
                    "parallelism": "number",
                    "projects": [
                      "list",
                      "string"
                    ],
                    "url": "string"
                  }
                ]
              ],
              "permission": [
                "list",
                [
                  "object",
                  {
                    "administrator_account": [
                      "list",
                      [
                        "object",
                        {
                          "groups": [
                            "list",
                            "string"
                          ],
                          "users": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "kind": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "dev_center_project_id": {
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
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_ids": [
                "list",
                "string"
              ],
              "type": "string"
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "maximum_concurrency": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "stateful_agent": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_resource_prediction": [
                "list",
                [
                  "object",
                  {
                    "prediction_preference": "string"
                  }
                ]
              ],
              "grace_period_time_span": "string",
              "manual_resource_prediction": [
                "list",
                [
                  "object",
                  {
                    "all_week_schedule": "number",
                    "friday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "monday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "saturday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "sunday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "thursday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "time_zone_name": "string",
                    "tuesday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "wednesday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "maximum_agent_lifetime": "string"
            }
          ]
        ]
      },
      "stateless_agent": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "automatic_resource_prediction": [
                "list",
                [
                  "object",
                  {
                    "prediction_preference": "string"
                  }
                ]
              ],
              "manual_resource_prediction": [
                "list",
                [
                  "object",
                  {
                    "all_week_schedule": "number",
                    "friday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "monday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "saturday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "sunday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "thursday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "time_zone_name": "string",
                    "tuesday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
                        }
                      ]
                    ],
                    "wednesday_schedule": [
                      "set",
                      [
                        "object",
                        {
                          "count": "number",
                          "time": "string"
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
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "virtual_machine_scale_set_fabric": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "image": [
                "list",
                [
                  "object",
                  {
                    "aliases": [
                      "list",
                      "string"
                    ],
                    "buffer": "string",
                    "id": "string",
                    "well_known_image_name": "string"
                  }
                ]
              ],
              "os_disk_storage_account_type": "string",
              "security": [
                "list",
                [
                  "object",
                  {
                    "interactive_logon_enabled": "bool",
                    "key_vault_management": [
                      "list",
                      [
                        "object",
                        {
                          "certificate_store_location": "string",
                          "certificate_store_name": "string",
                          "key_export_enabled": "bool",
                          "key_vault_certificate_ids": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "sku_name": "string",
              "storage": [
                "list",
                [
                  "object",
                  {
                    "caching": "string",
                    "disk_size_in_gb": "number",
                    "drive_letter": "string",
                    "storage_account_type": "string"
                  }
                ]
              ],
              "subnet_id": "string"
            }
          ]
        ]
      },
      "work_folder": {
        "computed": true,
        "description_kind": "plain",
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

func AzurermManagedDevopsPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermManagedDevopsPool), &result)
	return &result
}
