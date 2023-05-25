package data

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
      "data_residency_location": {
        "computed": true,
        "description": "Location in which the B2C tenant is hosted and data resides.",
        "description_kind": "plain",
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
        "computed": true,
        "description": "Billing SKU for the B2C tenant.",
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

func AzurermAadb2CDirectorySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAadb2CDirectory), &result)
	return &result
}
