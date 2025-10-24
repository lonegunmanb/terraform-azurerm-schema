package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermContainerAppEnvironment = `{
  "block": {
    "attributes": {
      "custom_domain_verification_id": {
        "computed": true,
        "description": "The ID of the Custom Domain Verification for this Container App Environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "default_domain": {
        "computed": true,
        "description": "The default publicly resolvable name of this Container App Environment",
        "description_kind": "plain",
        "type": "string"
      },
      "docker_bridge_cidr": {
        "computed": true,
        "description": "The network addressing in which the Container Apps in this Container App Environment will reside in CIDR notation.",
        "description_kind": "plain",
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "infrastructure_subnet_id": {
        "computed": true,
        "description": "The existing Subnet in use by the Container Apps Control Plane.",
        "description_kind": "plain",
        "type": "string"
      },
      "internal_load_balancer_enabled": {
        "computed": true,
        "description": "Does the Container Environment operate in Internal Load Balancing Mode?",
        "description_kind": "plain",
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "log_analytics_workspace_name": {
        "computed": true,
        "description": "The name of the Log Analytics Workspace this Container Apps Managed Environment is linked to.",
        "description_kind": "plain",
        "type": "string"
      },
      "name": {
        "description": "The name of the Container Apps Managed Environment.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "platform_reserved_cidr": {
        "computed": true,
        "description": "The IP range, in CIDR notation, that is reserved for environment infrastructure IP addresses.",
        "description_kind": "plain",
        "type": "string"
      },
      "platform_reserved_dns_ip_address": {
        "computed": true,
        "description": "The IP address from the IP range defined by ` + "`" + `platform_reserved_cidr` + "`" + ` that is reserved for the internal DNS server.",
        "description_kind": "plain",
        "type": "string"
      },
      "public_network_access": {
        "computed": true,
        "description": "The public network access setting for this Container App Environment.",
        "description_kind": "plain",
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "static_ip_address": {
        "computed": true,
        "description": "The Static IP Address of the Environment.",
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

func AzurermContainerAppEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironment), &result)
	return &result
}
