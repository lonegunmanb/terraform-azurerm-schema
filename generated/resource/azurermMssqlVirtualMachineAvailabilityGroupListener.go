package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMssqlVirtualMachineAvailabilityGroupListener = `{
  "block": {
    "attributes": {
      "availability_group_name": {
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
      "port": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "sql_virtual_machine_group_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "load_balancer_configuration": {
        "block": {
          "attributes": {
            "load_balancer_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "private_ip_address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "probe_port": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "sql_virtual_machine_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
            },
            "subnet_id": {
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
      "multi_subnet_ip_configuration": {
        "block": {
          "attributes": {
            "private_ip_address": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sql_virtual_machine_id": {
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
        "nesting_mode": "set"
      },
      "replica": {
        "block": {
          "attributes": {
            "commit": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "failover_mode": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "readable_secondary": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "role": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sql_virtual_machine_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
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

func AzurermMssqlVirtualMachineAvailabilityGroupListenerSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMssqlVirtualMachineAvailabilityGroupListener), &result)
	return &result
}
