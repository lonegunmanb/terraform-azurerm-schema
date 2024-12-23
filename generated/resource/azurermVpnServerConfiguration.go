package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVpnServerConfiguration = `{
  "block": {
    "attributes": {
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
      },
      "vpn_authentication_types": {
        "description_kind": "plain",
        "required": true,
        "type": [
          "list",
          "string"
        ]
      },
      "vpn_protocols": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "azure_active_directory_authentication": {
        "block": {
          "attributes": {
            "audience": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "issuer": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "tenant": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "client_revoked_certificate": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "thumbprint": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "client_root_certificate": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "public_cert_data": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "ipsec_policy": {
        "block": {
          "attributes": {
            "dh_group": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ike_encryption": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ike_integrity": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ipsec_encryption": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "ipsec_integrity": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "pfs_group": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "sa_data_size_kilobytes": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "sa_lifetime_seconds": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "radius": {
        "block": {
          "block_types": {
            "client_root_certificate": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "thumbprint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "server": {
              "block": {
                "attributes": {
                  "address": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "score": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "secret": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "server_root_certificate": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "public_cert_data": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
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

func AzurermVpnServerConfigurationSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVpnServerConfiguration), &result)
	return &result
}
