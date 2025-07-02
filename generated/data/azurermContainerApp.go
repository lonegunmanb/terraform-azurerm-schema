package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerApp = `{
  "block": {
    "attributes": {
      "container_app_environment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "custom_domain_verification_id": {
        "computed": true,
        "description": "The ID of the Custom Domain Verification for this Container App.",
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "dapr": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "app_id": "string",
              "app_port": "number",
              "app_protocol": "string"
            }
          ]
        ]
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity_ids": [
                "list",
                "string"
              ],
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "ingress": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "allow_insecure_connections": "bool",
              "client_certificate_mode": "string",
              "cors": [
                "list",
                [
                  "object",
                  {
                    "allow_credentials_enabled": "bool",
                    "allowed_headers": [
                      "list",
                      "string"
                    ],
                    "allowed_methods": [
                      "list",
                      "string"
                    ],
                    "allowed_origins": [
                      "list",
                      "string"
                    ],
                    "exposed_headers": [
                      "list",
                      "string"
                    ],
                    "max_age_in_seconds": "number"
                  }
                ]
              ],
              "custom_domain": [
                "list",
                [
                  "object",
                  {
                    "certificate_binding_type": "string",
                    "certificate_id": "string",
                    "name": "string"
                  }
                ]
              ],
              "exposed_port": "number",
              "external_enabled": "bool",
              "fqdn": "string",
              "ip_security_restriction": [
                "list",
                [
                  "object",
                  {
                    "action": "string",
                    "description": "string",
                    "ip_address_range": "string",
                    "name": "string"
                  }
                ]
              ],
              "target_port": "number",
              "traffic_weight": [
                "list",
                [
                  "object",
                  {
                    "label": "string",
                    "latest_revision": "bool",
                    "percentage": "number",
                    "revision_suffix": "string"
                  }
                ]
              ],
              "transport": "string"
            }
          ]
        ]
      },
      "latest_revision_fqdn": {
        "computed": true,
        "description": "The fully qualified domain name of the latest Container App.",
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
      "max_inactive_revisions": {
        "computed": true,
        "description_kind": "plain",
        "type": "number"
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
      "registry": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "identity": "string",
              "password_secret_name": "string",
              "server": "string",
              "username": "string"
            }
          ]
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "revision_mode": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "secret": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "set",
          [
            "object",
            {
              "identity": "string",
              "key_vault_secret_id": "string",
              "name": "string",
              "value": "string"
            }
          ]
        ]
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "template": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "azure_queue_scale_rule": [
                "list",
                [
                  "object",
                  {
                    "authentication": [
                      "list",
                      [
                        "object",
                        {
                          "secret_name": "string",
                          "trigger_parameter": "string"
                        }
                      ]
                    ],
                    "name": "string",
                    "queue_length": "number",
                    "queue_name": "string"
                  }
                ]
              ],
              "container": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "command": [
                      "list",
                      "string"
                    ],
                    "cpu": "number",
                    "env": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "secret_name": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "ephemeral_storage": "string",
                    "image": "string",
                    "liveness_probe": [
                      "list",
                      [
                        "object",
                        {
                          "failure_count_threshold": "number",
                          "header": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "host": "string",
                          "initial_delay": "number",
                          "interval_seconds": "number",
                          "path": "string",
                          "port": "number",
                          "termination_grace_period_seconds": "number",
                          "timeout": "number",
                          "transport": "string"
                        }
                      ]
                    ],
                    "memory": "string",
                    "name": "string",
                    "readiness_probe": [
                      "list",
                      [
                        "object",
                        {
                          "failure_count_threshold": "number",
                          "header": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "host": "string",
                          "initial_delay": "number",
                          "interval_seconds": "number",
                          "path": "string",
                          "port": "number",
                          "success_count_threshold": "number",
                          "timeout": "number",
                          "transport": "string"
                        }
                      ]
                    ],
                    "startup_probe": [
                      "list",
                      [
                        "object",
                        {
                          "failure_count_threshold": "number",
                          "header": [
                            "list",
                            [
                              "object",
                              {
                                "name": "string",
                                "value": "string"
                              }
                            ]
                          ],
                          "host": "string",
                          "initial_delay": "number",
                          "interval_seconds": "number",
                          "path": "string",
                          "port": "number",
                          "termination_grace_period_seconds": "number",
                          "timeout": "number",
                          "transport": "string"
                        }
                      ]
                    ],
                    "volume_mounts": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "path": "string",
                          "sub_path": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "custom_scale_rule": [
                "list",
                [
                  "object",
                  {
                    "authentication": [
                      "list",
                      [
                        "object",
                        {
                          "secret_name": "string",
                          "trigger_parameter": "string"
                        }
                      ]
                    ],
                    "custom_rule_type": "string",
                    "metadata": [
                      "map",
                      "string"
                    ],
                    "name": "string"
                  }
                ]
              ],
              "http_scale_rule": [
                "list",
                [
                  "object",
                  {
                    "authentication": [
                      "list",
                      [
                        "object",
                        {
                          "secret_name": "string",
                          "trigger_parameter": "string"
                        }
                      ]
                    ],
                    "concurrent_requests": "string",
                    "name": "string"
                  }
                ]
              ],
              "init_container": [
                "list",
                [
                  "object",
                  {
                    "args": [
                      "list",
                      "string"
                    ],
                    "command": [
                      "list",
                      "string"
                    ],
                    "cpu": "number",
                    "env": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "secret_name": "string",
                          "value": "string"
                        }
                      ]
                    ],
                    "ephemeral_storage": "string",
                    "image": "string",
                    "memory": "string",
                    "name": "string",
                    "volume_mounts": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "path": "string",
                          "sub_path": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "max_replicas": "number",
              "min_replicas": "number",
              "revision_suffix": "string",
              "tcp_scale_rule": [
                "list",
                [
                  "object",
                  {
                    "authentication": [
                      "list",
                      [
                        "object",
                        {
                          "secret_name": "string",
                          "trigger_parameter": "string"
                        }
                      ]
                    ],
                    "concurrent_requests": "string",
                    "name": "string"
                  }
                ]
              ],
              "termination_grace_period_seconds": "number",
              "volume": [
                "list",
                [
                  "object",
                  {
                    "mount_options": "string",
                    "name": "string",
                    "storage_name": "string",
                    "storage_type": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "workload_profile_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
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

func AzurermContainerAppSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerApp), &result)
	return &result
}
