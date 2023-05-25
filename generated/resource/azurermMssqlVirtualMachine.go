package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlVirtualMachine = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "r_services_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "sql_connectivity_port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sql_connectivity_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sql_connectivity_update_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "sql_connectivity_update_username": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "sql_license_type": {
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
      "virtual_machine_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "auto_backup": {
        "block": {
          "attributes": {
            "encryption_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "encryption_password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "retention_period_in_days": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "storage_account_access_key": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_blob_endpoint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "system_databases_backup_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "manual_schedule": {
              "block": {
                "attributes": {
                  "full_backup_frequency": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "full_backup_start_hour": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "full_backup_window_in_hours": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "log_backup_frequency_in_minutes": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
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
        "nesting_mode": "list"
      },
      "auto_patching": {
        "block": {
          "attributes": {
            "day_of_week": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "maintenance_window_duration_in_minutes": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "maintenance_window_starting_hour": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "key_vault_credential": {
        "block": {
          "attributes": {
            "key_vault_url": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_principal_name": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "service_principal_secret": {
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
      },
      "storage_configuration": {
        "block": {
          "attributes": {
            "disk_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "storage_workload_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_settings": {
              "block": {
                "attributes": {
                  "default_file_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "luns": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "log_settings": {
              "block": {
                "attributes": {
                  "default_file_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "luns": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "number"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "temp_db_settings": {
              "block": {
                "attributes": {
                  "default_file_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "luns": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "number"
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

func AzurermMssqlVirtualMachineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlVirtualMachine), &result)
	return &result
}
