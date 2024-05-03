package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerGroup = `{
  "block": {
    "attributes": {
      "dns_name_label": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "dns_name_label_reuse_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "exposed_port": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          [
            "object",
            {
              "port": "number",
              "protocol": "string"
            }
          ]
        ]
      },
      "fqdn": {
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
      "ip_address": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "ip_address_type": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_vault_key_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "key_vault_user_assigned_identity_id": {
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
      "network_profile_id": {
        "computed": true,
        "deprecated": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "os_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "priority": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "restart_policy": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "sku": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "subnet_ids": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      },
      "tags": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "zones": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "set",
          "string"
        ]
      }
    },
    "block_types": {
      "container": {
        "block": {
          "attributes": {
            "commands": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "cpu": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "cpu_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "environment_variables": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "image": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "memory": {
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "memory_limit": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secure_environment_variables": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "gpu": {
              "block": {
                "attributes": {
                  "count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "sku": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "gpu_limit": {
              "block": {
                "attributes": {
                  "count": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "sku": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "liveness_probe": {
              "block": {
                "attributes": {
                  "exec": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "failure_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "initial_delay_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "period_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "success_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "http_get": {
                    "block": {
                      "attributes": {
                        "http_headers": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "path": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "scheme": {
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
            },
            "ports": {
              "block": {
                "attributes": {
                  "port": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "protocol": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "set"
            },
            "readiness_probe": {
              "block": {
                "attributes": {
                  "exec": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "failure_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "initial_delay_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "period_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "success_threshold": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "timeout_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "http_get": {
                    "block": {
                      "attributes": {
                        "http_headers": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "path": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "scheme": {
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
            },
            "security": {
              "block": {
                "attributes": {
                  "privilege_enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "volume": {
              "block": {
                "attributes": {
                  "empty_dir": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "mount_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "read_only": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "share_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_account_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "storage_account_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "git_repo": {
                    "block": {
                      "attributes": {
                        "directory": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "revision": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "url": {
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
            }
          },
          "description_kind": "plain"
        },
        "min_items": 1,
        "nesting_mode": "list"
      },
      "diagnostics": {
        "block": {
          "block_types": {
            "log_analytics": {
              "block": {
                "attributes": {
                  "log_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "metadata": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "workspace_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "workspace_key": {
                    "description_kind": "plain",
                    "required": true,
                    "sensitive": true,
                    "type": "string"
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
        "max_items": 1,
        "nesting_mode": "list"
      },
      "dns_config": {
        "block": {
          "attributes": {
            "nameservers": {
              "description_kind": "plain",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "options": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "set",
                "string"
              ]
            },
            "search_domains": {
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
      "image_registry_credential": {
        "block": {
          "attributes": {
            "password": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "server": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "user_assigned_identity_id": {
              "description": "The User Assigned Identity to use for Container Registry access.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "username": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "init_container": {
        "block": {
          "attributes": {
            "commands": {
              "computed": true,
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "environment_variables": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "image": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secure_environment_variables": {
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": [
                "map",
                "string"
              ]
            }
          },
          "block_types": {
            "security": {
              "block": {
                "attributes": {
                  "privilege_enabled": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "volume": {
              "block": {
                "attributes": {
                  "empty_dir": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "mount_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "read_only": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "secret": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "share_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_account_key": {
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "storage_account_name": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "git_repo": {
                    "block": {
                      "attributes": {
                        "directory": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "revision": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "url": {
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
            }
          },
          "description_kind": "plain"
        },
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

func AzurermContainerGroupSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerGroup), &result)
	return &result
}
