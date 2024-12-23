package ephemeral

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKeyVaultCertificate = `{
  "block": {
    "attributes": {
      "certificate_count": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
      },
      "expiration_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "hex": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "key_vault_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "not_before_date": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "pem": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermKeyVaultCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKeyVaultCertificate), &result)
	return &result
}
