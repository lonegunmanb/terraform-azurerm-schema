package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermVirtualMachineScaleSet = `{
  "block": {
    "attributes": {
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
      "instances": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "computer_name": "string",
              "instance_id": "string",
              "latest_model_applied": "bool",
              "name": "string",
              "power_state": "string",
              "private_ip_address": "string",
              "private_ip_addresses": [
                "list",
                "string"
              ],
              "public_ip_address": "string",
              "public_ip_addresses": [
                "list",
                "string"
              ],
              "virtual_machine_id": "string",
              "zone": "string"
            }
          ]
        ]
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "network_interface": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "dns_servers": [
                "list",
                "string"
              ],
              "enable_accelerated_networking": "bool",
              "enable_ip_forwarding": "bool",
              "ip_configuration": [
                "list",
                [
                  "object",
                  {
                    "application_gateway_backend_address_pool_ids": [
                      "list",
                      "string"
                    ],
                    "application_security_group_ids": [
                      "list",
                      "string"
                    ],
                    "load_balancer_backend_address_pool_ids": [
                      "list",
                      "string"
                    ],
                    "load_balancer_inbound_nat_rules_ids": [
                      "list",
                      "string"
                    ],
                    "name": "string",
                    "primary": "bool",
                    "public_ip_address": [
                      "list",
                      [
                        "object",
                        {
                          "domain_name_label": "string",
                          "idle_timeout_in_minutes": "number",
                          "ip_tag": [
                            "list",
                            [
                              "object",
                              {
                                "tag": "string",
                                "type": "string"
                              }
                            ]
                          ],
                          "name": "string",
                          "public_ip_prefix_id": "string",
                          "version": "string"
                        }
                      ]
                    ],
                    "subnet_id": "string",
                    "version": "string"
                  }
                ]
              ],
              "name": "string",
              "network_security_group_id": "string",
              "primary": "bool"
            }
          ]
        ]
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
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

func AzurermVirtualMachineScaleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermVirtualMachineScaleSet), &result)
	return &result
}
