package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerApp = `{
  "block": {
    "attributes": {
      "container_app_environment_id": {
        "description": "The ID of the Container App Environment to host this Container App.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "custom_domain_verification_id": {
        "computed": true,
        "description": "The ID of the Custom Domain Verification for this Container App.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "latest_revision_fqdn": {
        "computed": true,
        "description": "The FQDN of the Latest Revision of the Container App.",
        "description_kind": "plain",
        "type": "string"
      },
      "latest_revision_name": {
        "computed": true,
        "description": "The name of the latest Container Revision.",
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name for this Container App.",
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
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision_mode": {
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
      "dapr": {
        "block": {
          "attributes": {
            "app_id": {
              "description": "The Dapr Application Identifier.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "app_port": {
              "description": "The port which the application is listening on. This is the same as the ` + "`" + `ingress` + "`" + ` port.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "app_protocol": {
              "description": "The protocol for the app. Possible values include ` + "`" + `http` + "`" + ` and ` + "`" + `grpc` + "`" + `. Defaults to ` + "`" + `http` + "`" + `.",
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
      "ingress": {
        "block": {
          "attributes": {
            "allow_insecure_connections": {
              "description": "Should this ingress allow insecure connections?",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "exposed_port": {
              "description": "The exposed port on the container for the Ingress traffic.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "external_enabled": {
              "description": "Is this an external Ingress.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "fqdn": {
              "computed": true,
              "description": "The FQDN of the ingress.",
              "description_kind": "plain",
              "type": "string"
            },
            "target_port": {
              "description": "The target port on the container for the Ingress traffic.",
              "description_kind": "plain",
              "required": true,
              "type": "number"
            },
            "transport": {
              "description": "The transport method for the Ingress. Possible values include ` + "`" + `auto` + "`" + `, ` + "`" + `http` + "`" + `, and ` + "`" + `http2` + "`" + `, ` + "`" + `tcp` + "`" + `. Defaults to ` + "`" + `auto` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "custom_domain": {
              "block": {
                "attributes": {
                  "certificate_binding_type": {
                    "description": "The Binding type. Possible values include ` + "`" + `Disabled` + "`" + ` and ` + "`" + `SniEnabled` + "`" + `. Defaults to ` + "`" + `Disabled` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "certificate_id": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "The hostname of the Certificate. Must be the CN or a named SAN in the certificate.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "deprecated": true,
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ip_security_restriction": {
              "block": {
                "attributes": {
                  "action": {
                    "description": "The action. Allow or Deny.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "description": {
                    "description": "Describe the IP restriction rule that is being sent to the container-app.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "ip_address_range": {
                    "description": "The incoming IP address or range of IP addresses (in CIDR notation).",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name": {
                    "description": "Name for the IP restriction rule.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "traffic_weight": {
              "block": {
                "attributes": {
                  "label": {
                    "description": "The label to apply to the revision as a name prefix for routing traffic.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "latest_revision": {
                    "description": "This traffic Weight relates to the latest stable Container Revision.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "percentage": {
                    "description": "The percentage of traffic to send to this revision.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "revision_suffix": {
                    "description": "The suffix string to append to the revision. This must be unique for the Container App's lifetime. A default hash created by the service will be used if this value is omitted.",
                    "description_kind": "plain",
                    "optional": true,
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
          "attributes": {
            "max_replicas": {
              "description": "The maximum number of replicas for this container.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "min_replicas": {
              "description": "The minimum number of replicas for this container.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "revision_suffix": {
              "computed": true,
              "description": "The suffix for the revision. This value must be unique for the lifetime of the Resource. If omitted the service will use a hash function to create one.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "azure_queue_scale_rule": {
              "block": {
                "attributes": {
                  "name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "queue_length": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "number"
                  },
                  "queue_name": {
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
                    "min_items": 1,
                    "nesting_mode": "list"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
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
                          "description": "The number of consecutive failures required to consider this probe as failed. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `10` + "`" + `. Defaults to ` + "`" + `3` + "`" + `.",
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
                          "description": "The time in seconds to wait after the container has started before the probe is started.",
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
                          "description": "The number of consecutive failures required to consider this probe as failed. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `10` + "`" + `. Defaults to ` + "`" + `3` + "`" + `.",
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
                          "description": "The number of consecutive failures required to consider this probe as failed. Possible values are between ` + "`" + `1` + "`" + ` and ` + "`" + `10` + "`" + `. Defaults to ` + "`" + `3` + "`" + `.",
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
            "custom_scale_rule": {
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
            },
            "http_scale_rule": {
              "block": {
                "attributes": {
                  "concurrent_requests": {
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
            "tcp_scale_rule": {
              "block": {
                "attributes": {
                  "concurrent_requests": {
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

func AzurermContainerAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerApp), &result)
	return &result
}
