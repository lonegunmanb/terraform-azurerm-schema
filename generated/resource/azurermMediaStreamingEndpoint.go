package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMediaStreamingEndpoint = `{
  "block": {
    "attributes": {
      "auto_start_enabled": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cdn_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cdn_profile": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "cdn_provider": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_host_names": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "host_name": {
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
      "max_cache_age_seconds": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "media_services_account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "scale_units": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "sku": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "capacity": "number",
              "name": "string"
            }
          ]
        ]
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
      "access_control": {
        "block": {
          "block_types": {
            "akamai_signature_header_authentication_key": {
              "block": {
                "attributes": {
                  "base64_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expiration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "identifier": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "ip_allow": {
              "block": {
                "attributes": {
                  "address": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "subnet_prefix_length": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "cross_site_access_policy": {
        "block": {
          "attributes": {
            "client_access_policy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "cross_domain_policy": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
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
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 1
}`

func AzurermMediaStreamingEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMediaStreamingEndpoint), &result)
	return &result
}
