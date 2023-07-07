package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMediaContentKeyPolicy = `{
  "block": {
    "attributes": {
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
      }
    },
    "block_types": {
      "policy_option": {
        "block": {
          "attributes": {
            "clear_key_configuration_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "open_restriction_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "playready_response_custom_data": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "widevine_configuration_template": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "fairplay_configuration": {
              "block": {
                "attributes": {
                  "ask": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "pfx": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "pfx_password": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "rental_and_lease_key_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "rental_duration_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "offline_rental_configuration": {
                    "block": {
                      "attributes": {
                        "playback_duration_seconds": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "storage_duration_seconds": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
            "playready_configuration_license": {
              "block": {
                "attributes": {
                  "allow_test_devices": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "begin_date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_key_location_from_header_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "content_key_location_from_key_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "content_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "expiration_date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "grace_period": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "license_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "relative_begin_date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "relative_expiration_date": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "security_level": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "play_right": {
                    "block": {
                      "attributes": {
                        "agc_and_color_stripe_restriction": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "allow_passing_video_content_to_unknown_output": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "analog_video_opl": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "compressed_digital_audio_opl": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "compressed_digital_video_opl": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "digital_video_only_content_restriction": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "first_play_expiration": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "image_constraint_for_analog_component_video_restriction": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "image_constraint_for_analog_computer_monitor_restriction": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "scms_restriction": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "uncompressed_digital_audio_opl": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "uncompressed_digital_video_opl": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        }
                      },
                      "block_types": {
                        "explicit_analog_television_output_restriction": {
                          "block": {
                            "attributes": {
                              "best_effort_enforced": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "control_bits": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "number"
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
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "token_restriction": {
              "block": {
                "attributes": {
                  "audience": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "issuer": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "open_id_connect_discovery_document": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "primary_rsa_token_key_exponent": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "primary_rsa_token_key_modulus": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "primary_symmetric_token_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "primary_x509_token_key_raw": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "token_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "alternate_key": {
                    "block": {
                      "attributes": {
                        "rsa_token_key_exponent": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "rsa_token_key_modulus": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "symmetric_token_key": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        },
                        "x509_token_key_raw": {
                          "description_kind": "plain",
                          "optional": true,
                          "sensitive": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "required_claim": {
                    "block": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "set"
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

func AzurermMediaContentKeyPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMediaContentKeyPolicy), &result)
	return &result
}
