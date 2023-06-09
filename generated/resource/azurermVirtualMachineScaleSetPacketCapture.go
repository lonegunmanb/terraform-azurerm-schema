package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualMachineScaleSetPacketCapture = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "maximum_bytes_per_packet": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maximum_bytes_per_session": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "maximum_capture_duration_in_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_watcher_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "virtual_machine_scale_set_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "filter": {
        "block": {
          "attributes": {
            "local_ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "local_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocol": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "remote_ip_address": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "remote_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "machine_scope": {
        "block": {
          "attributes": {
            "exclude_instance_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "include_instance_ids": {
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
      "storage_location": {
        "block": {
          "attributes": {
            "file_path": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_account_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "storage_path": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
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

func AzurermVirtualMachineScaleSetPacketCaptureSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualMachineScaleSetPacketCapture), &result)
	return &result
}
