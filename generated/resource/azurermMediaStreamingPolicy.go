package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMediaStreamingPolicy = `{
  "block": {
    "attributes": {
      "default_content_key_policy_name": {
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
      "common_encryption_cbcs": {
        "block": {
          "block_types": {
            "clear_key_encryption": {
              "block": {
                "attributes": {
                  "custom_keys_acquisition_url_template": {
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
            "default_content_key": {
              "block": {
                "attributes": {
                  "label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "policy_name": {
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
            "drm_fairplay": {
              "block": {
                "attributes": {
                  "allow_persistent_license": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "custom_license_acquisition_url_template": {
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
            "enabled_protocols": {
              "block": {
                "attributes": {
                  "dash": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "download": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "hls": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "smooth_streaming": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "common_encryption_cenc": {
        "block": {
          "attributes": {
            "drm_widevine_custom_license_acquisition_url_template": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "clear_key_encryption": {
              "block": {
                "attributes": {
                  "custom_keys_acquisition_url_template": {
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
            "clear_track": {
              "block": {
                "block_types": {
                  "condition": {
                    "block": {
                      "attributes": {
                        "operation": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "property": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "value": {
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
              "nesting_mode": "set"
            },
            "content_key_to_track_mapping": {
              "block": {
                "attributes": {
                  "label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "policy_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "track": {
                    "block": {
                      "block_types": {
                        "condition": {
                          "block": {
                            "attributes": {
                              "operation": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "property": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
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
                    "nesting_mode": "set"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "default_content_key": {
              "block": {
                "attributes": {
                  "label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "policy_name": {
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
            "drm_playready": {
              "block": {
                "attributes": {
                  "custom_attributes": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "custom_license_acquisition_url_template": {
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
            "enabled_protocols": {
              "block": {
                "attributes": {
                  "dash": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "download": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "hls": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "smooth_streaming": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "envelope_encryption": {
        "block": {
          "attributes": {
            "custom_keys_acquisition_url_template": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "default_content_key": {
              "block": {
                "attributes": {
                  "label": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "policy_name": {
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
            "enabled_protocols": {
              "block": {
                "attributes": {
                  "dash": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "download": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "hls": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "smooth_streaming": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
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
      "no_encryption_enabled_protocols": {
        "block": {
          "attributes": {
            "dash": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "download": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hls": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "smooth_streaming": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
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

func AzurermMediaStreamingPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMediaStreamingPolicy), &result)
	return &result
}
