package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCdnFrontdoorBatchRuleSet = `{
  "block": {
    "attributes": {
      "cdn_frontdoor_profile_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "rule": {
        "block": {
          "attributes": {
            "behaviour_on_match": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
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
            "actions": {
              "block": {
                "block_types": {
                  "modify_request_header": {
                    "block": {
                      "attributes": {
                        "header_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "header_value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
                  "modify_response_header": {
                    "block": {
                      "attributes": {
                        "header_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "header_value": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
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
                  "route_configuration_override": {
                    "block": {
                      "block_types": {
                        "caching": {
                          "block": {
                            "attributes": {
                              "behaviour": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "compression_enabled": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "duration": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "query_string_behaviour": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "query_string_parameters": {
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
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "origin_group": {
                          "block": {
                            "attributes": {
                              "cdn_frontdoor_origin_group_id": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "forwarding_protocol": {
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
                  "url_redirect": {
                    "block": {
                      "attributes": {
                        "destination_fragment": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "destination_host_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "destination_path": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "query_string": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "redirect_protocol": {
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
                  "url_rewrite": {
                    "block": {
                      "attributes": {
                        "destination_path": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "preserve_unmatched_path_enabled": {
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
              "min_items": 1,
              "nesting_mode": "list"
            },
            "conditions": {
              "block": {
                "block_types": {
                  "client_port": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                  "device_type": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                    "nesting_mode": "list"
                  },
                  "host_name": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transforms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "http_version": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "post_argument": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "query_string": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transforms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "remote_address": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                    "nesting_mode": "list"
                  },
                  "request_body": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transforms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "request_cookies": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "request_file_extension": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transforms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "request_filename": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transforms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "request_header": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "request_method": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "request_path": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transforms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "request_scheme": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                    "nesting_mode": "list"
                  },
                  "request_url": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "transforms": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        },
                        "values": {
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
                  "server_port": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "set",
                            "string"
                          ]
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "socket_address": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
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
                    "nesting_mode": "list"
                  },
                  "ssl_protocol": {
                    "block": {
                      "attributes": {
                        "operator": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "values": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "set",
                            "string"
                          ]
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
        "max_items": 100,
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermCdnFrontdoorBatchRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCdnFrontdoorBatchRuleSet), &result)
	return &result
}
