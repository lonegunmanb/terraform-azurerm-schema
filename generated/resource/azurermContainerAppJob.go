package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppJob = `{
  "block": {
    "attributes": {
      "container_app_environment_id": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "event_stream_endpoint": {
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
      "outbound_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "replica_retry_limit": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "replica_timeout_in_seconds": {
        "description_kind": "plain",
        "required": true,
        "type": "number"
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
      "workload_profile_name": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "event_trigger_config": {
        "block": {
          "attributes": {
            "parallelism": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replica_completion_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "scale": {
              "block": {
                "attributes": {
                  "max_executions": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "min_executions": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "polling_interval_in_seconds": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  }
                },
                "block_types": {
                  "rules": {
                    "block": {
                      "attributes": {
                        "custom_rule_type": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "metadata": {
                          "description_kind": "plain",
                          "required": true,
                          "type": [
                            "map",
                            "string"
                          ]
                        },
                        "name": {
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "authentication": {
                          "block": {
                            "attributes": {
                              "secret_name": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "trigger_parameter": {
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
              "nesting_mode": "list"
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
      "manual_trigger_config": {
        "block": {
          "attributes": {
            "parallelism": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replica_completion_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "registry": {
        "block": {
          "attributes": {
            "identity": {
              "description": "ID of the System or User Managed Identity used to pull images from the Container Registry",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "password_secret_name": {
              "description": "The name of the Secret Reference containing the password value for this user on the Container Registry.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "server": {
              "description": "The hostname for the Container Registry.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "username": {
              "description": "The username to use for this Container Registry.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "schedule_trigger_config": {
        "block": {
          "attributes": {
            "cron_expression": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "parallelism": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "replica_completion_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
        "nesting_mode": "list"
      },
      "secret": {
        "block": {
          "attributes": {
            "identity": {
              "description": "The identity to use for accessing key vault reference.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "key_vault_secret_id": {
              "description": "The Key Vault Secret ID. Could be either one of ` + "`" + `id` + "`" + ` or ` + "`" + `versionless_id` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "name": {
              "description": "The secret name.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The value for this secret.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
      },
      "template": {
        "block": {
          "block_types": {
            "container": {
              "block": {
                "attributes": {
                  "args": {
                    "description": "A list of args to pass to the container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "command": {
                    "description": "A command to pass to the container to override the default. This is provided as a list of command line elements without spaces.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cpu": {
                    "description": "The amount of vCPU to allocate to the container. Possible values include ` + "`" + `0.25` + "`" + `, ` + "`" + `0.5` + "`" + `, ` + "`" + `0.75` + "`" + `, ` + "`" + `1.0` + "`" + `, ` + "`" + `1.25` + "`" + `, ` + "`" + `1.5` + "`" + `, ` + "`" + `1.75` + "`" + `, and ` + "`" + `2.0` + "`" + `. **NOTE:** ` + "`" + `cpu` + "`" + ` and ` + "`" + `memory` + "`" + ` must be specified in ` + "`" + `0.25'/'0.5Gi` + "`" + ` combination increments. e.g. ` + "`" + `1.0` + "`" + ` / ` + "`" + `2.0` + "`" + ` or ` + "`" + `0.5` + "`" + ` / ` + "`" + `1.0` + "`" + `. When there's a workload profile specified, there's no such constraint.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "ephemeral_storage": {
                    "computed": true,
                    "description": "The amount of ephemeral storage available to the Container App.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "image": {
                    "description": "The image to use to create the container.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "memory": {
                    "description": "The amount of memory to allocate to the container. Possible values include ` + "`" + `0.5Gi` + "`" + `, ` + "`" + `1.0Gi` + "`" + `, ` + "`" + `1.5Gi` + "`" + `, ` + "`" + `2.0Gi` + "`" + `, ` + "`" + `2.5Gi` + "`" + `, ` + "`" + `3.0Gi` + "`" + `, ` + "`" + `3.5Gi` + "`" + `, and ` + "`" + `4.0Gi` + "`" + `. **NOTE:** ` + "`" + `cpu` + "`" + ` and ` + "`" + `memory` + "`" + ` must be specified in ` + "`" + `0.25'/'0.5Gi` + "`" + ` combination increments. e.g. ` + "`" + `1.25` + "`" + ` / ` + "`" + `2.5Gi` + "`" + ` or ` + "`" + `0.75` + "`" + ` / ` + "`" + `1.5Gi` + "`" + `. When there's a workload profile specified, there's no such constraint.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The name of the container.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "env": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "The name of the environment variable for the container.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_name": {
                          "description": "The name of the secret that contains the value for this environment variable.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "The value for this environment variable. **NOTE:** This value is ignored if ` + "`" + `secret_name` + "`" + ` is used",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "liveness_probe": {
                    "block": {
                      "attributes": {
                        "failure_count_threshold": {
                          "description": "The number of consecutive failures required to consider this probe as failed. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `30` + "`" + `. Defaults to ` + "`" + `3` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "host": {
                          "description": "The probe hostname. Defaults to the pod IP address. Setting a value for ` + "`" + `Host` + "`" + ` in ` + "`" + `headers` + "`" + ` can be used to override this for ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ` type probes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "initial_delay": {
                          "description": "The number of seconds elapsed after the container has started before the probe is initiated. Possible values are between ` + "`" + `0` + "`" + ` and ` + "`" + `60` + "`" + `. Defaults to ` + "`" + `1` + "`" + ` seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_seconds": {
                          "description": "How often, in seconds, the probe should run. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `240` + "`" + `. Defaults to ` + "`" + `10` + "`" + `",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "path": {
                          "computed": true,
                          "description": "The URI to use with the ` + "`" + `host` + "`" + ` for http type probes. Not valid for ` + "`" + `TCP` + "`" + ` type probes. Defaults to ` + "`" + `/` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "description": "The port number on which to connect. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "termination_grace_period_seconds": {
                          "computed": true,
                          "description": "The time in seconds after the container is sent the termination signal before the process if forcibly killed.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "timeout": {
                          "description": "Time in seconds after which the probe times out. Possible values are between ` + "`" + `1` + "`" + ` an ` + "`" + `240` + "`" + `. Defaults to ` + "`" + `1` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "transport": {
                          "description": "Type of probe. Possible values are ` + "`" + `TCP` + "`" + `, ` + "`" + `HTTP` + "`" + `, and ` + "`" + `HTTPS` + "`" + `.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "header": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "The HTTP Header Name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "The HTTP Header value.",
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
                  },
                  "readiness_probe": {
                    "block": {
                      "attributes": {
                        "failure_count_threshold": {
                          "description": "The number of consecutive failures required to consider this probe as failed. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `30` + "`" + `. Defaults to ` + "`" + `3` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "host": {
                          "description": "The probe hostname. Defaults to the pod IP address. Setting a value for ` + "`" + `Host` + "`" + ` in ` + "`" + `headers` + "`" + ` can be used to override this for ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ` type probes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "initial_delay": {
                          "description": "The number of seconds elapsed after the container has started before the probe is initiated. Possible values are between ` + "`" + `0` + "`" + ` and ` + "`" + `60` + "`" + `. Defaults to ` + "`" + `0` + "`" + ` seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_seconds": {
                          "description": "How often, in seconds, the probe should run. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `240` + "`" + `. Defaults to ` + "`" + `10` + "`" + `",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "path": {
                          "computed": true,
                          "description": "The URI to use for http type probes. Not valid for ` + "`" + `TCP` + "`" + ` type probes. Defaults to ` + "`" + `/` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "description": "The port number on which to connect. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "success_count_threshold": {
                          "description": "The number of consecutive successful responses required to consider this probe as successful. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `10` + "`" + `. Defaults to ` + "`" + `3` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "timeout": {
                          "description": "Time in seconds after which the probe times out. Possible values are between ` + "`" + `1` + "`" + ` an ` + "`" + `240` + "`" + `. Defaults to ` + "`" + `1` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "transport": {
                          "description": "Type of probe. Possible values are ` + "`" + `TCP` + "`" + `, ` + "`" + `HTTP` + "`" + `, and ` + "`" + `HTTPS` + "`" + `.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "header": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "The HTTP Header Name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "The HTTP Header value.",
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
                  },
                  "startup_probe": {
                    "block": {
                      "attributes": {
                        "failure_count_threshold": {
                          "description": "The number of consecutive failures required to consider this probe as failed. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `30` + "`" + `. Defaults to ` + "`" + `3` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "host": {
                          "description": "The probe hostname. Defaults to the pod IP address. Setting a value for ` + "`" + `Host` + "`" + ` in ` + "`" + `headers` + "`" + ` can be used to override this for ` + "`" + `http` + "`" + ` and ` + "`" + `https` + "`" + ` type probes.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "initial_delay": {
                          "description": "The number of seconds elapsed after the container has started before the probe is initiated. Possible values are between ` + "`" + `0` + "`" + ` and ` + "`" + `60` + "`" + `. Defaults to ` + "`" + `0` + "`" + ` seconds.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "interval_seconds": {
                          "description": "How often, in seconds, the probe should run. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `240` + "`" + `. Defaults to ` + "`" + `10` + "`" + `",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "path": {
                          "computed": true,
                          "description": "The URI to use with the ` + "`" + `host` + "`" + ` for http type probes. Not valid for ` + "`" + `TCP` + "`" + ` type probes. Defaults to ` + "`" + `/` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "port": {
                          "description": "The port number on which to connect. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `65535` + "`" + `.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "number"
                        },
                        "termination_grace_period_seconds": {
                          "computed": true,
                          "description": "The time in seconds after the container is sent the termination signal before the process if forcibly killed.",
                          "description_kind": "plain",
                          "type": "number"
                        },
                        "timeout": {
                          "description": "Time in seconds after which the probe times out. Possible values are between ` + "`" + `1` + "`" + ` an ` + "`" + `240` + "`" + `. Defaults to ` + "`" + `1` + "`" + `.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "transport": {
                          "description": "Type of probe. Possible values are ` + "`" + `TCP` + "`" + `, ` + "`" + `HTTP` + "`" + `, and ` + "`" + `HTTPS` + "`" + `.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "header": {
                          "block": {
                            "attributes": {
                              "name": {
                                "description": "The HTTP Header Name.",
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "value": {
                                "description": "The HTTP Header value.",
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
                  },
                  "volume_mounts": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "The name of the Volume to be mounted in the container.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description": "The path in the container at which to mount this volume.",
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
              "min_items": 1,
              "nesting_mode": "list"
            },
            "init_container": {
              "block": {
                "attributes": {
                  "args": {
                    "description": "A list of args to pass to the container.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "command": {
                    "description": "A command to pass to the container to override the default. This is provided as a list of command line elements without spaces.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cpu": {
                    "description": "The amount of vCPU to allocate to the container. Possible values include ` + "`" + `0.25` + "`" + `, ` + "`" + `0.5` + "`" + `, ` + "`" + `0.75` + "`" + `, ` + "`" + `1.0` + "`" + `, ` + "`" + `1.25` + "`" + `, ` + "`" + `1.5` + "`" + `, ` + "`" + `1.75` + "`" + `, and ` + "`" + `2.0` + "`" + `. **NOTE:** ` + "`" + `cpu` + "`" + ` and ` + "`" + `memory` + "`" + ` must be specified in ` + "`" + `0.25'/'0.5Gi` + "`" + ` combination increments. e.g. ` + "`" + `1.0` + "`" + ` / ` + "`" + `2.0` + "`" + ` or ` + "`" + `0.5` + "`" + ` / ` + "`" + `1.0` + "`" + `. When there's a workload profile specified, there's no such constraint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "ephemeral_storage": {
                    "computed": true,
                    "description": "The amount of ephemeral storage available to the Container App.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "image": {
                    "description": "The image to use to create the container.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "memory": {
                    "description": "The amount of memory to allocate to the container. Possible values include ` + "`" + `0.5Gi` + "`" + `, ` + "`" + `1.0Gi` + "`" + `, ` + "`" + `1.5Gi` + "`" + `, ` + "`" + `2.0Gi` + "`" + `, ` + "`" + `2.5Gi` + "`" + `, ` + "`" + `3.0Gi` + "`" + `, ` + "`" + `3.5Gi` + "`" + `, and ` + "`" + `4.0Gi` + "`" + `. **NOTE:** ` + "`" + `cpu` + "`" + ` and ` + "`" + `memory` + "`" + ` must be specified in ` + "`" + `0.25'/'0.5Gi` + "`" + ` combination increments. e.g. ` + "`" + `1.25` + "`" + ` / ` + "`" + `2.5Gi` + "`" + ` or ` + "`" + `0.75` + "`" + ` / ` + "`" + `1.5Gi` + "`" + `. When there's a workload profile specified, there's no such constraint.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The name of the container.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "env": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "The name of the environment variable for the container.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "secret_name": {
                          "description": "The name of the secret that contains the value for this environment variable.",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "value": {
                          "description": "The value for this environment variable. **NOTE:** This value is ignored if ` + "`" + `secret_name` + "`" + ` is used",
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "nesting_mode": "list"
                  },
                  "volume_mounts": {
                    "block": {
                      "attributes": {
                        "name": {
                          "description": "The name of the Volume to be mounted in the container.",
                          "description_kind": "plain",
                          "required": true,
                          "type": "string"
                        },
                        "path": {
                          "description": "The path in the container at which to mount this volume.",
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
            },
            "volume": {
              "block": {
                "attributes": {
                  "name": {
                    "description": "The name of the volume.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "storage_name": {
                    "description": "The name of the ` + "`" + `AzureFile` + "`" + ` storage. Required when ` + "`" + `storage_type` + "`" + ` is ` + "`" + `AzureFile` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "storage_type": {
                    "description": "The type of storage volume. Possible values include ` + "`" + `AzureFile` + "`" + ` and ` + "`" + `EmptyDir` + "`" + `. Defaults to ` + "`" + `EmptyDir` + "`" + `.",
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

func AzurermContainerAppJobSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppJob), &result)
	return &result
}
