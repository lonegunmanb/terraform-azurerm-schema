package resource

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
      "dapr_application_insights_connection_string": {
        "description": "Application Insights connection string used by Dapr to export Service to Service communication telemetry.",
        "description_kind": "plain",
        "optional": true,
        "sensitive": true,
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
      "infrastructure_resource_group_name": {
        "description": "Name of the platform-managed resource group created for the Managed Environment to host infrastructure resources. **Note:** Only valid if a ` + "`" + `workload_profile` + "`" + ` is specified. If ` + "`" + `infrastructure_subnet_id` + "`" + ` is specified, this resource group will be created in the same subscription as ` + "`" + `infrastructure_subnet_id` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "infrastructure_subnet_id": {
        "description": "The existing Subnet to use for the Container Apps Control Plane. **NOTE:** The Subnet must have a ` + "`" + `/21` + "`" + ` or larger address space.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "internal_load_balancer_enabled": {
        "description": "Should the Container Environment operate in Internal Load Balancing Mode? Defaults to ` + "`" + `false` + "`" + `. **Note:** can only be set to ` + "`" + `true` + "`" + ` if ` + "`" + `infrastructure_subnet_id` + "`" + ` is specified.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "log_analytics_workspace_id": {
        "description": "The ID for the Log Analytics Workspace to link this Container Apps Managed Environment to.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "logs_destination": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "mutual_tls_enabled": {
        "description": "Should mutual transport layer security (mTLS) be enabled? Defaults to ` + "`" + `false` + "`" + `. **Note:** This feature is in public preview. Enabling mTLS for your applications may increase response latency and reduce maximum throughput in high-load scenarios.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
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
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "zone_redundancy_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      }
    },
    "block_types": {
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
      "workload_profile": {
        "block": {
          "attributes": {
            "maximum_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "minimum_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "workload_profile_type": {
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
  "version": 0
}`

func AzurermContainerAppEnvironmentSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermContainerAppEnvironment), &result)
	return &result
}
