package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermDataFactoryLinkedServiceAzureBlobStorage = `{
  "block": {
    "attributes": {
      "additional_properties": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "annotations": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "connection_string": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "connection_string_insecure": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "data_factory_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "description": {
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
      "integration_runtime_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "parameters": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "sas_uri": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "service_endpoint": {
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
        "type": "string"
      },
      "service_principal_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "service_principal_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_kind": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "tenant_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "use_managed_identity": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
      "key_vault_sas_token": {
        "block": {
          "attributes": {
            "linked_service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secret_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "deprecated": true,
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "sas_token_linked_key_vault_key": {
        "block": {
          "attributes": {
            "linked_service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secret_name": {
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
      "service_principal_linked_key_vault_key": {
        "block": {
          "attributes": {
            "linked_service_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secret_name": {
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

func AzurermDataFactoryLinkedServiceAzureBlobStorageSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermDataFactoryLinkedServiceAzureBlobStorage), &result)
	return &result
}
