package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermWorkloadsSapThreeTierVirtualInstance = `{
  "block": {
    "attributes": {
      "app_location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "environment": {
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
      "managed_resource_group_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "managed_resources_network_access_type": {
        "description_kind": "plain",
        "optional": true,
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
      },
      "sap_fqdn": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "sap_product": {
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
      "three_tier_configuration": {
        "block": {
          "attributes": {
            "app_resource_group_name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "high_availability_type": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "secondary_ip_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "application_server_configuration": {
              "block": {
                "attributes": {
                  "instance_count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "virtual_machine_configuration": {
                    "block": {
                      "attributes": {
                        "virtual_machine_size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "image": {
                          "block": {
                            "attributes": {
                              "offer": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "publisher": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "sku": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "version": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "os_profile": {
                          "block": {
                            "attributes": {
                              "admin_username": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "ssh_private_key": {
                                "description_kind": "plain",
                                "required": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "ssh_public_key": {
                                "description_kind": "plain",
                                "required": true,
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
            "central_server_configuration": {
              "block": {
                "attributes": {
                  "instance_count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "virtual_machine_configuration": {
                    "block": {
                      "attributes": {
                        "virtual_machine_size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "image": {
                          "block": {
                            "attributes": {
                              "offer": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "publisher": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "sku": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "version": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "os_profile": {
                          "block": {
                            "attributes": {
                              "admin_username": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "ssh_private_key": {
                                "description_kind": "plain",
                                "required": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "ssh_public_key": {
                                "description_kind": "plain",
                                "required": true,
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
            "database_server_configuration": {
              "block": {
                "attributes": {
                  "database_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "instance_count": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "subnet_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "disk_volume_configuration": {
                    "block": {
                      "attributes": {
                        "number_of_disks": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "size_in_gb": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "sku_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "volume_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "set"
                  },
                  "virtual_machine_configuration": {
                    "block": {
                      "attributes": {
                        "virtual_machine_size": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "image": {
                          "block": {
                            "attributes": {
                              "offer": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "publisher": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "sku": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "version": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "description_kind": "plain"
                          },
                          "max_items": 1,
                          "min_items": 1,
                          "nesting_mode": "list"
                        },
                        "os_profile": {
                          "block": {
                            "attributes": {
                              "admin_username": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "ssh_private_key": {
                                "description_kind": "plain",
                                "required": true,
                                "sensitive": true,
                                "type": "string"
                              },
                              "ssh_public_key": {
                                "description_kind": "plain",
                                "required": true,
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
            "resource_names": {
              "block": {
                "block_types": {
                  "application_server": {
                    "block": {
                      "attributes": {
                        "availability_set_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "virtual_machine": {
                          "block": {
                            "attributes": {
                              "host_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "network_interface_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "os_disk_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "virtual_machine_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "data_disk": {
                                "block": {
                                  "attributes": {
                                    "names": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "volume_name": {
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
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "central_server": {
                    "block": {
                      "attributes": {
                        "availability_set_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "load_balancer": {
                          "block": {
                            "attributes": {
                              "backend_pool_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "frontend_ip_configuration_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "health_probe_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "name": {
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
                        "virtual_machine": {
                          "block": {
                            "attributes": {
                              "host_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "network_interface_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "os_disk_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "virtual_machine_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "data_disk": {
                                "block": {
                                  "attributes": {
                                    "names": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "volume_name": {
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
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "database_server": {
                    "block": {
                      "attributes": {
                        "availability_set_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "load_balancer": {
                          "block": {
                            "attributes": {
                              "backend_pool_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "frontend_ip_configuration_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "health_probe_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "name": {
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
                        "virtual_machine": {
                          "block": {
                            "attributes": {
                              "host_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "network_interface_names": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": [
                                  "list",
                                  "string"
                                ]
                              },
                              "os_disk_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "virtual_machine_name": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "data_disk": {
                                "block": {
                                  "attributes": {
                                    "names": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
                                    },
                                    "volume_name": {
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
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "max_items": 1,
                    "nesting_mode": "list"
                  },
                  "shared_storage": {
                    "block": {
                      "attributes": {
                        "account_name": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "private_endpoint_name": {
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
              "nesting_mode": "list"
            },
            "transport_create_and_mount": {
              "block": {
                "attributes": {
                  "resource_group_id": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_account_name": {
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

func AzurermWorkloadsSapThreeTierVirtualInstanceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermWorkloadsSapThreeTierVirtualInstance), &result)
	return &result
}
