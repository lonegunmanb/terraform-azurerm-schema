package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermMediaTransform = `{
  "block": {
    "attributes": {
      "description": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "media_services_account_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "output": {
        "block": {
          "attributes": {
            "on_error_action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "relative_priority": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "audio_analyzer_preset": {
              "block": {
                "attributes": {
                  "audio_analysis_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "audio_language": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "experimental_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "builtin_preset": {
              "block": {
                "attributes": {
                  "preset_name": {
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  }
                },
                "block_types": {
                  "preset_configuration": {
                    "block": {
                      "attributes": {
                        "complexity": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "interleave_output": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        },
                        "key_frame_interval_in_seconds": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_bitrate_bps": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_height": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "max_layers": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min_bitrate_bps": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
                        },
                        "min_height": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "number"
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
              "max_items": 1,
              "nesting_mode": "list"
            },
            "custom_preset": {
              "block": {
                "attributes": {
                  "experimental_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  }
                },
                "block_types": {
                  "codec": {
                    "block": {
                      "block_types": {
                        "aac_audio": {
                          "block": {
                            "attributes": {
                              "bitrate": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "channels": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "profile": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sampling_rate": {
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
                        "copy_audio": {
                          "block": {
                            "attributes": {
                              "label": {
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
                        "copy_video": {
                          "block": {
                            "attributes": {
                              "label": {
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
                        "dd_audio": {
                          "block": {
                            "attributes": {
                              "bitrate": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "channels": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sampling_rate": {
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
                        "h264_video": {
                          "block": {
                            "attributes": {
                              "complexity": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_frame_interval": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "rate_control_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "scene_change_detection_enabled": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "stretch_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sync_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "layer": {
                                "block": {
                                  "attributes": {
                                    "adaptive_b_frame_enabled": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "b_frames": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "bitrate": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "buffer_window": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "crf": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "entropy_mode": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "frame_rate": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "height": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "label": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "level": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "max_bitrate": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "profile": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "reference_frames": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "slices": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "width": {
                                      "computed": true,
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
                          "nesting_mode": "list"
                        },
                        "h265_video": {
                          "block": {
                            "attributes": {
                              "complexity": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "key_frame_interval": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "scene_change_detection_enabled": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "bool"
                              },
                              "stretch_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sync_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "layer": {
                                "block": {
                                  "attributes": {
                                    "adaptive_b_frame_enabled": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "bool"
                                    },
                                    "b_frames": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "bitrate": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "number"
                                    },
                                    "buffer_window": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "crf": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "frame_rate": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "height": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "label": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "level": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "max_bitrate": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "profile": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "reference_frames": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "slices": {
                                      "computed": true,
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "width": {
                                      "computed": true,
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
                          "nesting_mode": "list"
                        },
                        "jpg_image": {
                          "block": {
                            "attributes": {
                              "key_frame_interval": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "range": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sprite_column": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "number"
                              },
                              "start": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "step": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "stretch_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sync_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "layer": {
                                "block": {
                                  "attributes": {
                                    "height": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "label": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "quality": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "width": {
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
                          "nesting_mode": "list"
                        },
                        "png_image": {
                          "block": {
                            "attributes": {
                              "key_frame_interval": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "label": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "range": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "start": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "step": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "stretch_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "sync_mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "layer": {
                                "block": {
                                  "attributes": {
                                    "height": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "label": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "width": {
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
                          "nesting_mode": "list"
                        }
                      },
                      "description_kind": "plain"
                    },
                    "min_items": 1,
                    "nesting_mode": "list"
                  },
                  "filter": {
                    "block": {
                      "attributes": {
                        "rotation": {
                          "description_kind": "plain",
                          "optional": true,
                          "type": "string"
                        }
                      },
                      "block_types": {
                        "crop_rectangle": {
                          "block": {
                            "attributes": {
                              "height": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "left": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "top": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "width": {
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
                        "deinterlace": {
                          "block": {
                            "attributes": {
                              "mode": {
                                "description_kind": "plain",
                                "optional": true,
                                "type": "string"
                              },
                              "parity": {
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
                        "fade_in": {
                          "block": {
                            "attributes": {
                              "duration": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "fade_color": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "start": {
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
                        "fade_out": {
                          "block": {
                            "attributes": {
                              "duration": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "fade_color": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              },
                              "start": {
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
                        "overlay": {
                          "block": {
                            "block_types": {
                              "audio": {
                                "block": {
                                  "attributes": {
                                    "audio_gain_level": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "end": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "fade_in_duration": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "fade_out_duration": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "input_label": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "start": {
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
                              "video": {
                                "block": {
                                  "attributes": {
                                    "audio_gain_level": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "end": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "fade_in_duration": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "fade_out_duration": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    },
                                    "input_label": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": "string"
                                    },
                                    "opacity": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "number"
                                    },
                                    "start": {
                                      "description_kind": "plain",
                                      "optional": true,
                                      "type": "string"
                                    }
                                  },
                                  "block_types": {
                                    "crop_rectangle": {
                                      "block": {
                                        "attributes": {
                                          "height": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "left": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "top": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "width": {
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
                                    "position": {
                                      "block": {
                                        "attributes": {
                                          "height": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "left": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "top": {
                                            "description_kind": "plain",
                                            "optional": true,
                                            "type": "string"
                                          },
                                          "width": {
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
                                "max_items": 1,
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
                  "format": {
                    "block": {
                      "block_types": {
                        "jpg": {
                          "block": {
                            "attributes": {
                              "filename_pattern": {
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
                        "mp4": {
                          "block": {
                            "attributes": {
                              "filename_pattern": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "output_file": {
                                "block": {
                                  "attributes": {
                                    "labels": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
                        "png": {
                          "block": {
                            "attributes": {
                              "filename_pattern": {
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
                        "transport_stream": {
                          "block": {
                            "attributes": {
                              "filename_pattern": {
                                "description_kind": "plain",
                                "required": true,
                                "type": "string"
                              }
                            },
                            "block_types": {
                              "output_file": {
                                "block": {
                                  "attributes": {
                                    "labels": {
                                      "description_kind": "plain",
                                      "required": true,
                                      "type": [
                                        "list",
                                        "string"
                                      ]
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
            "face_detector_preset": {
              "block": {
                "attributes": {
                  "analysis_resolution": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "blur_type": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "experimental_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "face_redactor_mode": {
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
            "video_analyzer_preset": {
              "block": {
                "attributes": {
                  "audio_analysis_mode": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "audio_language": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "experimental_options": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "insights_type": {
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
    "deprecated": true,
    "description_kind": "plain"
  },
  "version": 1
}`

func AzurermMediaTransformSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermMediaTransform), &result)
	return &result
}
