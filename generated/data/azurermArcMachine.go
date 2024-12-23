package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermArcMachine = `{
  "block": {
    "attributes": {
      "active_directory_fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "agent": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "extensions_allow_list": [
                "list",
                [
                  "object",
                  {
                    "publisher": "string",
                    "type": "string"
                  }
                ]
              ],
              "extensions_block_list": [
                "list",
                [
                  "object",
                  {
                    "publisher": "string",
                    "type": "string"
                  }
                ]
              ],
              "extensions_enabled": "bool",
              "guest_configuration_enabled": "bool",
              "incoming_connections_ports": [
                "list",
                "string"
              ],
              "proxy_bypass": [
                "list",
                "string"
              ],
              "proxy_url": "string"
            }
          ]
        ]
      },
      "agent_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "client_public_key": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "cloud_metadata": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "provider": "string"
            }
          ]
        ]
      },
      "detected_properties": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "display_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "dns_fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "domain_name": {
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
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "principal_id": "string",
              "tenant_id": "string",
              "type": "string"
            }
          ]
        ]
      },
      "last_status_change_time": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location_data": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "city": "string",
              "country_or_region": "string",
              "district": "string",
              "name": "string"
            }
          ]
        ]
      },
      "machine_fqdn": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "mssql_discovered": {
        "computed": true,
        "description_kind": "plain",
        "type": "bool"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "os_name": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "os_profile": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "computer_name": "string",
              "linux": [
                "list",
                [
                  "object",
                  {
                    "patch": [
                      "list",
                      [
                        "object",
                        {
                          "assessment_mode": "string",
                          "patch_mode": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "windows": [
                "list",
                [
                  "object",
                  {
                    "patch": [
                      "list",
                      [
                        "object",
                        {
                          "assessment_mode": "string",
                          "patch_mode": "string"
                        }
                      ]
                    ]
                  }
                ]
              ]
            }
          ]
        ]
      },
      "os_sku": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "os_type": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "os_version": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "parent_cluster_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "private_link_scope_resource_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_status": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "extension_service": [
                "list",
                [
                  "object",
                  {
                    "startup_type": "string",
                    "status": "string"
                  }
                ]
              ],
              "guest_configuration_service": [
                "list",
                [
                  "object",
                  {
                    "startup_type": "string",
                    "status": "string"
                  }
                ]
              ]
            }
          ]
        ]
      },
      "status": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "tags": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "map",
          "string"
        ]
      },
      "vm_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "vm_uuid": {
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

func AzurermArcMachineSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermArcMachine), &result)
	return &result
}
