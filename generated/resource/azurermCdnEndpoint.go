package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCdnEndpoint = `{
  "block": {
    "attributes": {
      "content_types_to_compress": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
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
      "is_compression_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_http_allowed": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "is_https_allowed": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
      "optimization_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_host_header": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "origin_path": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "probe_path": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "profile_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "querystring_caching_behaviour": {
        "description_kind": "plain",
        "optional": true,
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
      "delivery_rule": {
        "block": {
          "attributes": {
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "order": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            }
          },
          "block_types": {
            "cache_expiration_action": {
              "block": {
                "attributes": {
                  "behavior": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "duration": {
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
            "cache_key_query_string_action": {
              "block": {
                "attributes": {
                  "behavior": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "parameters": {
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
            "cookies_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "selector": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "device_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
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
            "http_version_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "modify_request_header_action": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
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
            },
            "modify_response_header_action": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
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
            },
            "post_arg_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "selector": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "query_string_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "remote_address_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "request_body_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "request_header_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "selector": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "request_method_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
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
            "request_scheme_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
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
            "request_uri_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "url_file_extension_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "url_file_name_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "url_path_condition": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "negate_condition": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "transforms": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "url_redirect_action": {
              "block": {
                "attributes": {
                  "fragment": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "hostname": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query_string": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "redirect_type": {
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
            "url_rewrite_action": {
              "block": {
                "attributes": {
                  "destination": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "preserve_unmatched_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_pattern": {
                    "description_kind": "plain",
                    "required": true,
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
        "nesting_mode": "list"
      },
      "geo_filter": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "country_codes": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "relative_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "global_delivery_rule": {
        "block": {
          "block_types": {
            "cache_expiration_action": {
              "block": {
                "attributes": {
                  "behavior": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "duration": {
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
            "cache_key_query_string_action": {
              "block": {
                "attributes": {
                  "behavior": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "parameters": {
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
            "modify_request_header_action": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
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
            },
            "modify_response_header_action": {
              "block": {
                "attributes": {
                  "action": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
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
            },
            "url_redirect_action": {
              "block": {
                "attributes": {
                  "fragment": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "hostname": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "query_string": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "redirect_type": {
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
            "url_rewrite_action": {
              "block": {
                "attributes": {
                  "destination": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "preserve_unmatched_path": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "source_pattern": {
                    "description_kind": "plain",
                    "required": true,
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
      "origin": {
        "block": {
          "attributes": {
            "host_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "http_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "https_port": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
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
    "description_kind": "plain"
  },
  "version": 1
}`

func AzurermCdnEndpointSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCdnEndpoint), &result)
	return &result
}
