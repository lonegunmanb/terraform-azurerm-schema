package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermWebApplicationFirewallPolicy = `{
  "block": {
    "attributes": {
      "http_listener_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
      "path_based_rule_ids": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
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
      "custom_rules": {
        "block": {
          "attributes": {
            "action": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "group_rate_limit_by": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "priority": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "rate_limit_duration": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "rate_limit_threshold": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "rule_type": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "block_types": {
            "match_conditions": {
              "block": {
                "attributes": {
                  "match_values": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "negation_condition": {
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
                      "set",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "match_variables": {
                    "block": {
                      "attributes": {
                        "selector": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "variable_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
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
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "managed_rules": {
        "block": {
          "block_types": {
            "exclusion": {
              "block": {
                "attributes": {
                  "match_variable": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "selector": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "selector_match_operator": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "excluded_rule_set": {
                    "block": {
                      "attributes": {
                        "type": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "version": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "rule_group": {
                          "block": {
                            "attributes": {
                              "excluded_rules": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "rule_group_name": {
                                "description_kind": "plain",
                                "required": true,
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
              "nesting_mode": "list"
            },
            "managed_rule_set": {
              "block": {
                "attributes": {
                  "type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "version": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "rule_group_override": {
                    "block": {
                      "attributes": {
                        "rule_group_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "rule": {
                          "block": {
                            "attributes": {
                              "action": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "enabled": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "id": {
                                "description_kind": "plain",
                                "required": true,
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
                    "nesting_mode": "list"
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
        "max_items": 1,
        "min_items": 1,
        "nesting_mode": "list"
      },
      "policy_settings": {
        "block": {
          "attributes": {
            "enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "file_upload_enforcement": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "file_upload_limit_in_mb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "js_challenge_cookie_expiration_in_minutes": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "max_request_body_size_in_kb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "mode": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "request_body_check": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "request_body_enforcement": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "request_body_inspect_limit_in_kb": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "log_scrubbing": {
              "block": {
                "attributes": {
                  "enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "rule": {
                    "block": {
                      "attributes": {
                        "enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "match_variable": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "selector": {
                          "description": "When matchVariable is a collection, operator used to specify which elements in the collection this rule applies to.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "selector_match_operator": {
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
  "version": 1
}`

func AzurermWebApplicationFirewallPolicySchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermWebApplicationFirewallPolicy), &result)
	return &result
}
