package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermKeyVaultCertificate = `{
  "block": {
    "attributes": {
      "certificate_data": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_data_base64": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "certificate_policy": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "issuer_parameters": [
                "list",
                [
                  "object",
                  {
                    "name": "string"
                  }
                ]
              ],
              "key_properties": [
                "list",
                [
                  "object",
                  {
                    "curve": "string",
                    "exportable": "bool",
                    "key_size": "number",
                    "key_type": "string",
                    "reuse_key": "bool"
                  }
                ]
              ],
              "lifetime_action": [
                "list",
                [
                  "object",
                  {
                    "action": [
                      "list",
                      [
                        "object",
                        {
                          "action_type": "string"
                        }
                      ]
                    ],
                    "trigger": [
                      "list",
                      [
                        "object",
                        {
                          "days_before_expiry": "number",
                          "lifetime_percentage": "number"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "secret_properties": [
                "list",
                [
                  "object",
                  {
                    "content_type": "string"
                  }
                ]
              ],
              "x509_certificate_properties": [
                "list",
                [
                  "object",
                  {
                    "extended_key_usage": [
                      "list",
                      "string"
                    ],
                    "key_usage": [
                      "list",
                      "string"
                    ],
                    "subject": "string",
                    "subject_alternative_names": [
                      "list",
                      [
                        "object",
                        {
                          "dns_names": [
                            "list",
                            "string"
                          ],
                          "emails": [
                            "list",
                            "string"
                          ],
                          "upns": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "validity_in_months": "number"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "expires": {
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
      "not_before": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secret_id": {
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
      "thumbprint": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "version": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "versionless_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "versionless_secret_id": {
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

func AzurermKeyVaultCertificateSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermKeyVaultCertificate), &result)
	return &result
}
