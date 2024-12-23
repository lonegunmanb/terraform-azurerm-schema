package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkPacketCoreControlPlane = `{
  "block": {
    "attributes": {
      "control_plane_access_ipv4_address": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "control_plane_access_ipv4_gateway": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "control_plane_access_ipv4_subnet": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "control_plane_access_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "core_network_technology": {
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
      "interoperability_settings_json": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
      "site_ids": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "sku": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "software_version": {
        "description_kind": "plain",
        "optional": true,
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
      "user_equipment_mtu_in_bytes": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      }
    },
    "block_types": {
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
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
      "local_diagnostics_access": {
        "block": {
          "attributes": {
            "authentication_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "https_server_certificate_url": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "platform": {
        "block": {
          "attributes": {
            "arc_kubernetes_cluster_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "custom_location_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "edge_device_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "stack_hci_cluster_id": {
              "description_kind": "plain",
              "optional": true,
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

func AzurermMobileNetworkPacketCoreControlPlaneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMobileNetworkPacketCoreControlPlane), &result)
	return &result
}
