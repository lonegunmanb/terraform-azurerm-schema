package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOrbitalContactProfile = `{
  "block": {
    "attributes": {
      "auto_tracking": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_hub_uri": {
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
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "minimum_elevation_degrees": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "minimum_variable_contact_duration": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_configuration_subnet_id": {
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
      }
    },
    "block_types": {
      "links": {
        "block": {
          "attributes": {
            "direction": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "polarization": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "channels": {
              "block": {
                "attributes": {
                  "bandwidth_mhz": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "center_frequency_mhz": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "demodulation_configuration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "modulation_configuration": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "end_point": {
                    "block": {
                      "attributes": {
                        "end_point_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "ip_address": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "protocol": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
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
  "version": 0
}`

func AzurermOrbitalContactProfileSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOrbitalContactProfile), &result)
	return &result
}
