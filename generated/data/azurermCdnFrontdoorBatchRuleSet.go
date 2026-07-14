package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermCdnFrontdoorBatchRuleSet = `{
  "block": {
    "attributes": {
      "cdn_frontdoor_profile_id": {
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
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "profile_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "rules": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          [
            "object",
            {
              "actions": [
                "list",
                [
                  "object",
                  {
                    "modify_request_header": [
                      "list",
                      [
                        "object",
                        {
                          "header_name": "string",
                          "header_value": "string",
                          "operator": "string"
                        }
                      ]
                    ],
                    "modify_response_header": [
                      "list",
                      [
                        "object",
                        {
                          "header_name": "string",
                          "header_value": "string",
                          "operator": "string"
                        }
                      ]
                    ],
                    "route_configuration_override": [
                      "list",
                      [
                        "object",
                        {
                          "caching": [
                            "list",
                            [
                              "object",
                              {
                                "behaviour": "string",
                                "compression_enabled": "bool",
                                "duration": "string",
                                "query_string_behaviour": "string",
                                "query_string_parameters": [
                                  "list",
                                  "string"
                                ]
                              }
                            ]
                          ],
                          "origin_group": [
                            "list",
                            [
                              "object",
                              {
                                "cdn_frontdoor_origin_group_id": "string",
                                "forwarding_protocol": "string"
                              }
                            ]
                          ]
                        }
                      ]
                    ],
                    "url_redirect": [
                      "list",
                      [
                        "object",
                        {
                          "destination_fragment": "string",
                          "destination_host_name": "string",
                          "destination_path": "string",
                          "query_string": "string",
                          "redirect_protocol": "string",
                          "redirect_type": "string"
                        }
                      ]
                    ],
                    "url_rewrite": [
                      "list",
                      [
                        "object",
                        {
                          "destination_path": "string",
                          "preserve_unmatched_path_enabled": "bool",
                          "source_pattern": "string"
                        }
                      ]
                    ]
                  }
                ]
              ],
              "behaviour_on_match": "string",
              "conditions": [
                "list",
                [
                  "object",
                  {
                    "client_port": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "device_type": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "host_name": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "http_version": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "post_argument": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "query_string": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "remote_address": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_body": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_cookies": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_file_extension": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_filename": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_header": [
                      "list",
                      [
                        "object",
                        {
                          "name": "string",
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_method": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_path": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_scheme": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "request_url": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "transforms": [
                            "set",
                            "string"
                          ],
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "server_port": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "socket_address": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ],
                    "ssl_protocol": [
                      "list",
                      [
                        "object",
                        {
                          "operator": "string",
                          "values": [
                            "set",
                            "string"
                          ]
                        }
                      ]
                    ]
                  }
                ]
              ],
              "name": "string",
              "order": "number"
            }
          ]
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

func AzurermCdnFrontdoorBatchRuleSetSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermCdnFrontdoorBatchRuleSet), &result)
	return &result
}
