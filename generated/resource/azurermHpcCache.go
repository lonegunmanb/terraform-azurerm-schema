package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermHpcCache = `{
  "block": {
    "attributes": {
      "automatically_rotate_key_to_latest_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "cache_size_in_gb": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_vault_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "mount_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "mtu": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "ntp_server": {
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
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "subnet_id": {
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
      "default_access_policy": {
        "block": {
          "block_types": {
            "access_rule": {
              "block": {
                "attributes": {
                  "access": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "anonymous_gid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "anonymous_uid": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "filter": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "root_squash_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "scope": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "submount_access_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "suid_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 3,
              "min_items": 1,
              "nesting_mode": "set"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "directory_active_directory": {
        "block": {
          "attributes": {
            "cache_netbios_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "dns_primary_ip": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "dns_secondary_ip": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "domain_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "domain_netbios_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password": {
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            },
            "username": {
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
      "directory_flat_file": {
        "block": {
          "attributes": {
            "group_file_uri": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "password_file_uri": {
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
      "directory_ldap": {
        "block": {
          "attributes": {
            "base_dn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "certificate_validation_uri": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "download_certificate_automatically": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "encrypted": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "server": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "bind": {
              "block": {
                "attributes": {
                  "dn": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "password": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "dns": {
        "block": {
          "attributes": {
            "search_domain": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "servers": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
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

func AzurermHpcCacheSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermHpcCache), &result)
	return &result
}
