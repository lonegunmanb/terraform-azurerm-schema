package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermAttestationProvider = `{
  "block": {
    "attributes": {
      "attestation_uri": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "open_enclave_policy_base64": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "policy_signing_certificate_data": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sev_snp_policy_base64": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sgx_enclave_policy_base64": {
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
      "tpm_policy_base64": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "trust_model": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "policy": {
        "block": {
          "attributes": {
            "data": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "environment_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
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

func AzurermAttestationProviderSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermAttestationProvider), &result)
	return &result
}
