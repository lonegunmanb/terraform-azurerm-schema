package resource

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
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
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
        "description_kind": "plain",
        "required": true,
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
      "nsxt_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "vcenter_certificate_thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vcenter_password": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "management_cluster": {
        "block": {
          "attributes": {
            "hosts": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "number"
            },
            "size": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
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

func AzurermVmwarePrivateCloudSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVmwarePrivateCloud), &result)
	return &result
}
