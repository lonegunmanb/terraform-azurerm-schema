package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAadb2CDirectory = `{
  "block": {
    "attributes": {
      "billing_type": {
        "computed": true,
        "description": "The type of billing for the B2C tenant. Possible values include: ` + "`" + `MAU` + "`" + ` or ` + "`" + `Auths` + "`" + `.",
        "description_kind": "plain",
        "type": "string"
      },
      "country_code": {
        "computed": true,
        "description": "Country code of the B2C tenant. See https://aka.ms/B2CDataResidency for valid country codes.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_residency_location": {
        "description": "Location in which the B2C tenant is hosted and data resides. See https://aka.ms/B2CDataResidency for more information.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "display_name": {
        "computed": true,
        "description": "The initial display name of the B2C tenant.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "domain_name": {
        "description": "Domain name of the B2C tenant, including onmicrosoft.com suffix.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "effective_start_date": {
        "computed": true,
        "description": "The date from which the billing type took effect. May not be populated until after the first billing cycle.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sku_name": {
        "description": "Billing SKU for the B2C tenant. See https://aka.ms/b2cBilling for more information.",
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
      "tenant_id": {
        "computed": true,
        "description": "The Tenant ID for the B2C tenant.",
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
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

func AzurermAadb2CDirectorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAadb2CDirectory), &result)
	return &result
}
