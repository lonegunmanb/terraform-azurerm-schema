package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMobileNetworkPacketCoreControlPlane = `{
  "block": {
    "attributes": {
      "control_plane_access_ipv4_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "control_plane_access_ipv4_gateway": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "control_plane_access_ipv4_subnet": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "control_plane_access_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "core_network_technology": {
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
      "interoperability_settings_json": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "local_diagnostics_access": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "authentication_type": "string",
              "https_server_certificate_url": "string"
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "arc_kubernetes_cluster_id": "string",
              "custom_location_id": "string",
              "edge_device_id": "string",
              "stack_hci_cluster_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "sku": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "software_version": {
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
      },
      "user_equipment_mtu_in_bytes": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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

func AzurermMobileNetworkPacketCoreControlPlaneSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMobileNetworkPacketCoreControlPlane), &result)
	return &result
}
