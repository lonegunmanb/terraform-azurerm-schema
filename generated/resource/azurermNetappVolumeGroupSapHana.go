package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermNetappVolumeGroupSapHana = `{
  "block": {
    "attributes": {
      "account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "application_identifier": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "group_description": {
        "description_kind": "plain",
        "required": true,
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
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
      },
      "volume": {
        "block": {
          "attributes": {
            "capacity_pool_id": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "encryption_key_source": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "id": {
              "computed": true,
              "description_kind": "plain",
              "type": "string"
            },
            "key_vault_private_endpoint_id": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "mount_ip_addresses": {
              "computed": true,
              "description_kind": "plain",
              "type": [
                "list",
                "string"
              ]
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "network_features": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "protocols": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "proximity_placement_group_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "security_style": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "service_level": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "snapshot_directory_visible": {
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "storage_quota_in_gb": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
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
            },
            "throughput_in_mibps": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "volume_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "volume_spec_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "zone": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "data_protection_replication": {
              "block": {
                "attributes": {
                  "endpoint_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "remote_volume_location": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "remote_volume_resource_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "replication_frequency": {
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
            "data_protection_snapshot_policy": {
              "block": {
                "attributes": {
                  "snapshot_policy_id": {
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
            "export_policy_rule": {
              "block": {
                "attributes": {
                  "allowed_clients": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "nfsv3_enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "nfsv41_enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  },
                  "root_access_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "rule_index": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "unix_read_only": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "unix_read_write": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 5,
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 5,
        "min_items": 2,
        "nesting_mode": "list"
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermNetappVolumeGroupSapHanaSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermNetappVolumeGroupSapHana), &result)
	return &result
}
