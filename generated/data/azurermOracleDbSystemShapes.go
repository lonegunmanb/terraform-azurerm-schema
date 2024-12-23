package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermOracleDbSystemShapes = `{
  "block": {
    "attributes": {
      "db_system_shapes": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "available_core_count": "number",
              "available_core_count_per_node": "number",
              "available_data_storage_in_tbs": "number",
              "available_data_storage_per_server_in_tbs": "number",
              "available_db_node_per_node_in_gbs": "number",
              "available_db_node_storage_in_gbs": "number",
              "available_memory_in_gbs": "number",
              "available_memory_per_node_in_gbs": "number",
              "core_count_increment": "number",
              "maximum_node_count": "number",
              "maximum_storage_count": "number",
              "minimum_core_count": "number",
              "minimum_core_count_per_node": "number",
              "minimum_data_storage_in_tbs": "number",
              "minimum_db_node_storage_per_node_in_gbs": "number",
              "minimum_memory_per_node_in_gbs": "number",
              "minimum_node_count": "number",
              "minimum_storage_count": "number",
              "runtime_minimum_core_count": "number",
              "shape_family": "string"
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
      "location": {
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

func AzurermOracleDbSystemShapesSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermOracleDbSystemShapes), &result)
	return &result
}
