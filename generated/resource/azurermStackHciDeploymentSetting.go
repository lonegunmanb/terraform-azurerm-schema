package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermStackHciDeploymentSetting = `{
  "block": {
    "attributes": {
      "arc_resource_ids": {
        "description_kind": "plain",
        "required": true,
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
      "stack_hci_cluster_id": {
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
    "block_types": {
      "scale_unit": {
        "block": {
          "attributes": {
            "active_directory_organizational_unit_path": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "bitlocker_boot_volume_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "bitlocker_data_volume_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "credential_guard_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "domain_fqdn": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "drift_control_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "drtm_protection_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "episodic_data_upload_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "eu_location_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "hvci_protection_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "name_prefix": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "secrets_location": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "side_channel_mitigation_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "smb_cluster_encryption_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "smb_signing_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "streaming_data_client_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "wdac_enabled": {
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            }
          },
          "block_types": {
            "cluster": {
              "block": {
                "attributes": {
                  "azure_service_endpoint": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "cloud_account_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "witness_path": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "witness_type": {
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
            "host_network": {
              "block": {
                "attributes": {
                  "storage_auto_ip_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "storage_connectivity_switchless_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "block_types": {
                  "intent": {
                    "block": {
                      "attributes": {
                        "adapter": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "adapter_property_override_enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "qos_policy_override_enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        },
                        "traffic_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "list",
                            "string"
                          ]
                        },
                        "virtual_switch_configuration_override_enabled": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "bool"
                        }
                      },
                      "block_types": {
                        "adapter_property_override": {
                          "block": {
                            "attributes": {
                              "jumbo_packet": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "network_direct": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "network_direct_technology": {
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
                        "qos_policy_override": {
                          "block": {
                            "attributes": {
                              "bandwidth_percentage_smb": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "priority_value8021_action_cluster": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "priority_value8021_action_smb": {
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
                        "virtual_switch_configuration_override": {
                          "block": {
                            "attributes": {
                              "enable_iov": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "load_balancing_algorithm": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "storage_network": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "network_adapter_name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "vlan_id": {
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
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "infrastructure_network": {
              "block": {
                "attributes": {
                  "dhcp_enabled": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "dns_server": {
                    "description_kind": "plain",
                    "required": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "gateway": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "subnet_mask": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "ip_pool": {
                    "block": {
                      "attributes": {
                        "ending_address": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "starting_address": {
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
            },
            "optional_service": {
              "block": {
                "attributes": {
                  "custom_location": {
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
            "physical_node": {
              "block": {
                "attributes": {
                  "ipv4_address": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
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
              "nesting_mode": "list"
            },
            "storage": {
              "block": {
                "attributes": {
                  "configuration_mode": {
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

func AzurermStackHciDeploymentSettingSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermStackHciDeploymentSetting), &result)
	return &result
}
