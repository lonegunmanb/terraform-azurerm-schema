package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVmwarePrivateCloud = `{
  "block": {
    "attributes": {
      "circuit": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "express_route_id": "string",
              "express_route_private_peering_id": "string",
              "primary_subnet_cidr": "string",
              "secondary_subnet_cidr": "string"
            }
          ]
        ]
      },
      "hcx_cloud_manager_endpoint": {
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
      "internet_connection_enabled": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "management_cluster": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "hosts": [
                "list",
                "string"
              ],
              "id": "number",
              "size": "number"
            }
          ]
        ]
      },
      "management_subnet_cidr": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_subnet_cidr": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "nsxt_certificate_thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "nsxt_manager_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "provisioning_subnet_cidr": {
        "computed": true,
        "description_kind": "plain",
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
      },
      "vcenter_certificate_thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vcsa_endpoint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vmotion_subnet_cidr": {
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

func AzurermVmwarePrivateCloudSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVmwarePrivateCloud), &result)
	return &result
}
