package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermManagedDevopsPool = `{
  "block": {
    "attributes": {
      "dev_center_project_id": {
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
      "maximum_concurrency": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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
      "work_folder": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "azure_devops_organization": {
        "block": {
          "block_types": {
            "organization": {
              "block": {
                "attributes": {
                  "parallelism": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "projects": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "url": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "permission": {
              "block": {
                "attributes": {
                  "kind": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "administrator_account": {
                    "block": {
                      "attributes": {
                        "groups": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "users": {
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
        "min_items": 1,
        "nesting_mode": "list"
      },
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "set",
                "string"
              ]
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
      "stateful_agent": {
        "block": {
          "attributes": {
            "grace_period_time_span": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "maximum_agent_lifetime": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "automatic_resource_prediction": {
              "block": {
                "attributes": {
                  "prediction_preference": {
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
            "manual_resource_prediction": {
              "block": {
                "attributes": {
                  "all_week_schedule": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "time_zone_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "friday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "monday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "saturday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "sunday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "thursday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "tuesday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "wednesday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
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
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "stateless_agent": {
        "block": {
          "block_types": {
            "automatic_resource_prediction": {
              "block": {
                "attributes": {
                  "prediction_preference": {
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
            "manual_resource_prediction": {
              "block": {
                "attributes": {
                  "all_week_schedule": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "time_zone_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "friday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "monday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "saturday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "sunday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "thursday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "tuesday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "wednesday_schedule": {
                    "block": {
                      "attributes": {
                        "count": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "time": {
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
      },
      "virtual_machine_scale_set_fabric": {
        "block": {
          "attributes": {
            "os_disk_storage_account_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "sku_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "subnet_id": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "image": {
              "block": {
                "attributes": {
                  "aliases": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "buffer": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "well_known_image_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "min_items": 1,
              "nesting_mode": "list"
            },
            "security": {
              "block": {
                "attributes": {
                  "interactive_logon_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "key_vault_management": {
                    "block": {
                      "attributes": {
                        "certificate_store_location": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "certificate_store_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_export_enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "key_vault_certificate_ids": {
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
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "storage": {
              "block": {
                "attributes": {
                  "caching": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "disk_size_in_gb": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "drive_letter": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_account_type": {
                    "description_kind": "plain",
                    "optional": true,
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
      }
    },
    "description_kind": "plain"
  },
  "version": 0
}`

func AzurermManagedDevopsPoolSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermManagedDevopsPool), &result)
	return &result
}
