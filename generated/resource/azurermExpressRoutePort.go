package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermExpressRoutePort = `{
  "block": {
    "attributes": {
      "bandwidth_in_gbps": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "billing_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "encapsulation": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ethertype": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "guid": {
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mtu": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "peering_location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
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
      }
    },
    "block_types": {
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "principal_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description_kind": "plain",
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
      "link1": {
        "block": {
          "attributes": {
            "admin_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "connector_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "interface_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "macsec_cak_keyvault_secret_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "macsec_cipher": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "macsec_ckn_keyvault_secret_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "macsec_sci_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "patch_panel_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rack_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "router_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "link2": {
        "block": {
          "attributes": {
            "admin_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "connector_type": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "interface_name": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "macsec_cak_keyvault_secret_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "macsec_cipher": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "macsec_ckn_keyvault_secret_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "macsec_sci_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "patch_panel_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "rack_id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "router_name": {
              "computed": true,
              "description_kind": "plain",
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

func AzurermExpressRoutePortSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermExpressRoutePort), &result)
	return &result
}
