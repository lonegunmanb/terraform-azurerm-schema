package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azurermFunctionAppFlexConsumption = `{
  "block": {
    "attributes": {
      "app_settings": {
        "description": "A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.",
        "description_kind": "plain",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "client_certificate_enabled": {
        "description": "Should the function app use Client Certificates",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "client_certificate_exclusion_paths": {
        "description": "Paths to exclude when using client certificates, separated by ;",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "client_certificate_mode": {
        "description": "The mode of the Function App's client certificates requirement for incoming requests. Possible values are ` + "`" + `Required` + "`" + `, ` + "`" + `Optional` + "`" + `, and ` + "`" + `OptionalInteractiveUser` + "`" + ` ",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "custom_domain_verification_id": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": "string"
      },
      "default_hostname": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "enabled": {
        "description": "Is the Function App enabled.",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "hosting_environment_id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "https_only": {
        "description": "Can the Function App only be accessed via HTTPS?",
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "instance_memory_in_mb": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "kind": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "location": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "maximum_instance_count": {
        "description_kind": "plain",
        "optional": true,
        "type": "number"
      },
      "name": {
        "description": "Specifies the name of the Function App.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "outbound_ip_address_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "outbound_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "possible_outbound_ip_address_list": {
        "computed": true,
        "description_kind": "plain",
        "type": [
          "list",
          "string"
        ]
      },
      "possible_outbound_ip_addresses": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "public_network_access_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "resource_group_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_name": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "runtime_version": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "service_plan_id": {
        "description": "The ID of the App Service Plan within which to create this Function App",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "site_credential": {
        "computed": true,
        "description_kind": "plain",
        "sensitive": true,
        "type": [
          "list",
          [
            "object",
            {
              "name": "string",
              "password": "string"
            }
          ]
        ]
      },
      "storage_access_key": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "storage_authentication_type": {
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_container_endpoint": {
        "description": "The endpoint of the storage container where the function app's code is hosted.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_container_type": {
        "description": "The type of the storage container where the function app's code is hosted. Only ` + "`" + `blobContainer` + "`" + ` is supported currently.",
        "description_kind": "plain",
        "required": true,
        "type": "string"
      },
      "storage_user_assigned_identity_id": {
        "description_kind": "plain",
        "optional": true,
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
      "virtual_network_subnet_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "webdeploy_publish_basic_authentication_enabled": {
        "description_kind": "plain",
        "optional": true,
        "type": "bool"
      },
      "zip_deploy_file": {
        "computed": true,
        "description": "The local path and filename of the Zip packaged application to deploy to this Function App. **Note:** Using this value requires either ` + "`" + `WEBSITE_RUN_FROM_PACKAGE=1` + "`" + ` or ` + "`" + `SCM_DO_BUILD_DURING_DEPLOYMENT=true` + "`" + ` to be set on the App in ` + "`" + `app_settings` + "`" + `.",
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      }
    },
    "block_types": {
      "always_ready": {
        "block": {
          "attributes": {
            "instance_count": {
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "name": {
              "description_kind": "plain",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "auth_settings": {
        "block": {
          "attributes": {
            "additional_login_parameters": {
              "description": "Specifies a map of Login Parameters to send to the OpenID Connect authorization endpoint when a user logs in.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "map",
                "string"
              ]
            },
            "allowed_external_redirect_urls": {
              "computed": true,
              "description": "Specifies a list of External URLs that can be redirected to as part of logging in or logging out of the Windows Web App.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "default_provider": {
              "computed": true,
              "description": "The default authentication provider to use when multiple providers are configured. Possible values include: ` + "`" + `AzureActiveDirectory` + "`" + `, ` + "`" + `Facebook` + "`" + `, ` + "`" + `Google` + "`" + `, ` + "`" + `MicrosoftAccount` + "`" + `, ` + "`" + `Twitter` + "`" + `, ` + "`" + `Github` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "enabled": {
              "description": "Should the Authentication / Authorization feature be enabled?",
              "description_kind": "plain",
              "required": true,
              "type": "bool"
            },
            "issuer": {
              "description": "The OpenID Connect Issuer URI that represents the entity which issues access tokens.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "runtime_version": {
              "computed": true,
              "description": "The RuntimeVersion of the Authentication / Authorization feature in use.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "token_refresh_extension_hours": {
              "description": "The number of hours after session token expiration that a session token can be used to call the token refresh API. Defaults to ` + "`" + `72` + "`" + ` hours.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "token_store_enabled": {
              "description": "Should the Windows Web App durably store platform-specific security tokens that are obtained during login flows? Defaults to ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "unauthenticated_client_action": {
              "computed": true,
              "description": "The action to take when an unauthenticated client attempts to access the app. Possible values include: ` + "`" + `RedirectToLoginPage` + "`" + `, ` + "`" + `AllowAnonymous` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "active_directory": {
              "block": {
                "attributes": {
                  "allowed_audiences": {
                    "description": "Specifies a list of Allowed audience values to consider when validating JWTs issued by Azure Active Directory.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "client_id": {
                    "description": "The ID of the Client to use to authenticate with Azure Active Directory.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description": "The Client Secret for the Client ID. Cannot be used with ` + "`" + `client_secret_setting_name` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The App Setting name that contains the client secret of the Client. Cannot be used with ` + "`" + `client_secret` + "`" + `.",
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
            "facebook": {
              "block": {
                "attributes": {
                  "app_id": {
                    "description": "The App ID of the Facebook app used for login.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "app_secret": {
                    "description": "The App Secret of the Facebook app used for Facebook Login. Cannot be specified with ` + "`" + `app_secret_setting_name` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "app_secret_setting_name": {
                    "description": "The app setting name that contains the ` + "`" + `app_secret` + "`" + ` value used for Facebook Login. Cannot be specified with ` + "`" + `app_secret` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "description": "Specifies a list of OAuth 2.0 scopes to be requested as part of Facebook Login authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "github": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The ID of the GitHub app used for login.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description": "The Client Secret of the GitHub app used for GitHub Login. Cannot be specified with ` + "`" + `client_secret_setting_name` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The app setting name that contains the ` + "`" + `client_secret` + "`" + ` value used for GitHub Login. Cannot be specified with ` + "`" + `client_secret` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "description": "Specifies a list of OAuth 2.0 scopes that will be requested as part of GitHub Login authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "google": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The OpenID Connect Client ID for the Google web application.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description": "The client secret associated with the Google web application.  Cannot be specified with ` + "`" + `client_secret_setting_name` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The app setting name that contains the ` + "`" + `client_secret` + "`" + ` value used for Google Login. Cannot be specified with ` + "`" + `client_secret` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "description": "Specifies a list of OAuth 2.0 scopes that will be requested as part of Google Sign-In authentication. If not specified, \"openid\", \"profile\", and \"email\" are used as default scopes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "microsoft": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The OAuth 2.0 client ID that was created for the app used for authentication.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret": {
                    "description": "The OAuth 2.0 client secret that was created for the app used for authentication. Cannot be specified with ` + "`" + `client_secret_setting_name` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The app setting name containing the OAuth 2.0 client secret that was created for the app used for authentication. Cannot be specified with ` + "`" + `client_secret` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "oauth_scopes": {
                    "description": "The list of OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication. If not specified, ` + "`" + `wl.basic` + "`" + ` is used as the default scope.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "twitter": {
              "block": {
                "attributes": {
                  "consumer_key": {
                    "description": "The OAuth 1.0a consumer key of the Twitter application used for sign-in.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "consumer_secret": {
                    "description": "The OAuth 1.0a consumer secret of the Twitter application used for sign-in. Cannot be specified with ` + "`" + `consumer_secret_setting_name` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "sensitive": true,
                    "type": "string"
                  },
                  "consumer_secret_setting_name": {
                    "description": "The app setting name that contains the OAuth 1.0a consumer secret of the Twitter application used for sign-in. Cannot be specified with ` + "`" + `consumer_secret` + "`" + `.",
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
      },
      "auth_settings_v2": {
        "block": {
          "attributes": {
            "auth_enabled": {
              "description": "Should the AuthV2 Settings be enabled. Defaults to ` + "`" + `false` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "config_file_path": {
              "description": "The path to the App Auth settings. **Note:** Relative Paths are evaluated from the Site Root directory.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "default_provider": {
              "description": "The Default Authentication Provider to use when the ` + "`" + `unauthenticated_action` + "`" + ` is set to ` + "`" + `RedirectToLoginPage` + "`" + `. Possible values include: ` + "`" + `apple` + "`" + `, ` + "`" + `azureactivedirectory` + "`" + `, ` + "`" + `facebook` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `google` + "`" + `, ` + "`" + `twitter` + "`" + ` and the ` + "`" + `name` + "`" + ` of your ` + "`" + `custom_oidc_v2` + "`" + ` provider.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "excluded_paths": {
              "description": "The paths which should be excluded from the ` + "`" + `unauthenticated_action` + "`" + ` when it is set to ` + "`" + `RedirectToLoginPage` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "forward_proxy_convention": {
              "description": "The convention used to determine the url of the request made. Possible values include ` + "`" + `ForwardProxyConventionNoProxy` + "`" + `, ` + "`" + `ForwardProxyConventionStandard` + "`" + `, ` + "`" + `ForwardProxyConventionCustom` + "`" + `. Defaults to ` + "`" + `ForwardProxyConventionNoProxy` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "forward_proxy_custom_host_header_name": {
              "description": "The name of the header containing the host of the request.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "forward_proxy_custom_scheme_header_name": {
              "description": "The name of the header containing the scheme of the request.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http_route_api_prefix": {
              "description": "The prefix that should precede all the authentication and authorisation paths. Defaults to ` + "`" + `/.auth` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "require_authentication": {
              "description": "Should the authentication flow be used for all requests.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "require_https": {
              "description": "Should HTTPS be required on connections? Defaults to true.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "runtime_version": {
              "description": "The Runtime Version of the Authentication and Authorisation feature of this App. Defaults to ` + "`" + `~1` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "unauthenticated_action": {
              "description": "The action to take for requests made without authentication. Possible values include ` + "`" + `RedirectToLoginPage` + "`" + `, ` + "`" + `AllowAnonymous` + "`" + `, ` + "`" + `Return401` + "`" + `, and ` + "`" + `Return403` + "`" + `. Defaults to ` + "`" + `RedirectToLoginPage` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "block_types": {
            "active_directory_v2": {
              "block": {
                "attributes": {
                  "allowed_applications": {
                    "description": "The list of allowed Applications for the Default Authorisation Policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_audiences": {
                    "description": "Specifies a list of Allowed audience values to consider when validating JWTs issued by Azure Active Directory.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_groups": {
                    "description": "The list of allowed Group Names for the Default Authorisation Policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "allowed_identities": {
                    "description": "The list of allowed Identities for the Default Authorisation Policy.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "client_id": {
                    "description": "The ID of the Client to use to authenticate with Azure Active Directory.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret_certificate_thumbprint": {
                    "description": "The thumbprint of the certificate used for signing purposes.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The App Setting name that contains the client secret of the Client.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "jwt_allowed_client_applications": {
                    "description": "A list of Allowed Client Applications in the JWT Claim.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "jwt_allowed_groups": {
                    "description": "A list of Allowed Groups in the JWT Claim.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "login_parameters": {
                    "description": "A map of key-value pairs to send to the Authorisation Endpoint when a user logs in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "map",
                      "string"
                    ]
                  },
                  "tenant_auth_endpoint": {
                    "description": "The Azure Tenant Endpoint for the Authenticating Tenant. e.g. ` + "`" + `https://login.microsoftonline.com/v2.0/{tenant-guid}/` + "`" + `.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "www_authentication_disabled": {
                    "description": "Should the www-authenticate provider should be omitted from the request? Defaults to ` + "`" + `false` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "apple_v2": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The OpenID Connect Client ID for the Apple web application.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The app setting name that contains the ` + "`" + `client_secret` + "`" + ` value used for Apple Login.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "login_scopes": {
                    "computed": true,
                    "description_kind": "plain",
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "azure_static_web_app_v2": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The ID of the Client to use to authenticate with Azure Static Web App Authentication.",
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
            "custom_oidc_v2": {
              "block": {
                "attributes": {
                  "authorisation_endpoint": {
                    "computed": true,
                    "description": "The endpoint to make the Authorisation Request.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "certification_uri": {
                    "computed": true,
                    "description": "The endpoint that provides the keys necessary to validate the token.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_credential_method": {
                    "computed": true,
                    "description": "The Client Credential Method used. Currently the only supported value is ` + "`" + `ClientSecretPost` + "`" + `.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "client_id": {
                    "description": "The ID of the Client to use to authenticate with this Custom OIDC.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "computed": true,
                    "description": "The App Setting name that contains the secret for this Custom OIDC Client.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "issuer_endpoint": {
                    "computed": true,
                    "description": "The endpoint that issued the Token.",
                    "description_kind": "plain",
                    "type": "string"
                  },
                  "name": {
                    "description": "The name of the Custom OIDC Authentication Provider.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "name_claim_type": {
                    "description": "The name of the claim that contains the users name.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "openid_configuration_endpoint": {
                    "description": "The endpoint that contains all the configuration endpoints for this Custom OIDC provider.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "scopes": {
                    "description": "The list of the scopes that should be requested while authenticating.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "token_endpoint": {
                    "computed": true,
                    "description": "The endpoint used to request a Token.",
                    "description_kind": "plain",
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "facebook_v2": {
              "block": {
                "attributes": {
                  "app_id": {
                    "description": "The App ID of the Facebook app used for login.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "app_secret_setting_name": {
                    "description": "The app setting name that contains the ` + "`" + `app_secret` + "`" + ` value used for Facebook Login.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "graph_api_version": {
                    "computed": true,
                    "description": "The version of the Facebook API to be used while logging in.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "login_scopes": {
                    "description": "Specifies a list of scopes to be requested as part of Facebook Login authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "github_v2": {
              "block": {
                "attributes": {
                  "client_id": {
                    "description": "The ID of the GitHub app used for login.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The app setting name that contains the ` + "`" + `client_secret` + "`" + ` value used for GitHub Login.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "login_scopes": {
                    "description": "Specifies a list of OAuth 2.0 scopes that will be requested as part of GitHub Login authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "google_v2": {
              "block": {
                "attributes": {
                  "allowed_audiences": {
                    "description": "Specifies a list of Allowed Audiences that will be requested as part of Google Sign-In authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "client_id": {
                    "description": "The OpenID Connect Client ID for the Google web application.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The app setting name that contains the ` + "`" + `client_secret` + "`" + ` value used for Google Login.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "login_scopes": {
                    "description": "Specifies a list of Login scopes that will be requested as part of Google Sign-In authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "login": {
              "block": {
                "attributes": {
                  "allowed_external_redirect_urls": {
                    "description": "External URLs that can be redirected to as part of logging in or logging out of the app. This is an advanced setting typically only needed by Windows Store application backends. **Note:** URLs within the current domain are always implicitly allowed.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "cookie_expiration_convention": {
                    "description": "The method by which cookies expire. Possible values include: ` + "`" + `FixedTime` + "`" + `, and ` + "`" + `IdentityProviderDerived` + "`" + `. Defaults to ` + "`" + `FixedTime` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "cookie_expiration_time": {
                    "description": "The time after the request is made when the session cookie should expire. Defaults to ` + "`" + `08:00:00` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "logout_endpoint": {
                    "description": "The endpoint to which logout requests should be made.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "nonce_expiration_time": {
                    "description": "The time after the request is made when the nonce should expire. Defaults to ` + "`" + `00:05:00` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "preserve_url_fragments_for_logins": {
                    "description": "Should the fragments from the request be preserved after the login request is made. Defaults to ` + "`" + `false` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "token_refresh_extension_time": {
                    "description": "The number of hours after session token expiration that a session token can be used to call the token refresh API. Defaults to ` + "`" + `72` + "`" + ` hours.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "token_store_enabled": {
                    "description": "Should the Token Store configuration Enabled. Defaults to ` + "`" + `false` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  },
                  "token_store_path": {
                    "description": "The directory path in the App Filesystem in which the tokens will be stored.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "token_store_sas_setting_name": {
                    "description": "The name of the app setting which contains the SAS URL of the blob storage containing the tokens.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "validate_nonce": {
                    "description": "Should the nonce be validated while completing the login flow. Defaults to ` + "`" + `true` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "min_items": 1,
              "nesting_mode": "list"
            },
            "microsoft_v2": {
              "block": {
                "attributes": {
                  "allowed_audiences": {
                    "description": "Specifies a list of Allowed Audiences that will be requested as part of Microsoft Sign-In authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  },
                  "client_id": {
                    "description": "The OAuth 2.0 client ID that was created for the app used for authentication.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "client_secret_setting_name": {
                    "description": "The app setting name containing the OAuth 2.0 client secret that was created for the app used for authentication.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "login_scopes": {
                    "description": "The list of Login scopes that will be requested as part of Microsoft Account authentication.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      "string"
                    ]
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "twitter_v2": {
              "block": {
                "attributes": {
                  "consumer_key": {
                    "description": "The OAuth 1.0a consumer key of the Twitter application used for sign-in.",
                    "description_kind": "plain",
                    "required": true,
                    "type": "string"
                  },
                  "consumer_secret_setting_name": {
                    "description": "The app setting name that contains the OAuth 1.0a consumer secret of the Twitter application used for sign-in.",
                    "description_kind": "plain",
                    "required": true,
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
      },
      "connection_string": {
        "block": {
          "attributes": {
            "name": {
              "description": "The name which should be used for this Connection.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "type": {
              "description": "Type of database. Possible values include: ` + "`" + `MySQL` + "`" + `, ` + "`" + `SQLServer` + "`" + `, ` + "`" + `SQLAzure` + "`" + `, ` + "`" + `Custom` + "`" + `, ` + "`" + `NotificationHub` + "`" + `, ` + "`" + `ServiceBus` + "`" + `, ` + "`" + `EventHub` + "`" + `, ` + "`" + `APIHub` + "`" + `, ` + "`" + `DocDb` + "`" + `, ` + "`" + `RedisCache` + "`" + `, and ` + "`" + `PostgreSQL` + "`" + `.",
              "description_kind": "plain",
              "required": true,
              "type": "string"
            },
            "value": {
              "description": "The connection string value.",
              "description_kind": "plain",
              "required": true,
              "sensitive": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "set"
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
      "site_config": {
        "block": {
          "attributes": {
            "api_definition_url": {
              "description": "The URL of the API definition that describes this Linux Function App.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "api_management_api_id": {
              "description": "The ID of the API Management API for this Linux Function App.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "app_command_line": {
              "description": "The program and any arguments used to launch this app via the command line. (Example ` + "`" + `node myapp.js` + "`" + `).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "application_insights_connection_string": {
              "description": "The Connection String for linking the Linux Function App to Application Insights.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "application_insights_key": {
              "description": "The Instrumentation Key for connecting the Linux Function App to Application Insights.",
              "description_kind": "plain",
              "optional": true,
              "sensitive": true,
              "type": "string"
            },
            "container_registry_managed_identity_client_id": {
              "description": "The Client ID of the Managed Service Identity to use for connections to the Azure Container Registry.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "container_registry_use_managed_identity": {
              "description": "Should connections for Azure Container Registry use Managed Identity.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "default_documents": {
              "computed": true,
              "description": "Specifies a list of Default Documents for the Linux Web App.",
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "detailed_error_logging_enabled": {
              "computed": true,
              "description": "Is detailed error logging enabled",
              "description_kind": "plain",
              "type": "bool"
            },
            "elastic_instance_minimum": {
              "computed": true,
              "description": "The number of minimum instances for this Linux Function App. Only affects apps on Elastic Premium plans.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "health_check_eviction_time_in_min": {
              "description": "The amount of time in minutes that a node is unhealthy before being removed from the load balancer. Possible values are between ` + "`" + `2` + "`" + ` and ` + "`" + `10` + "`" + `. Only valid in conjunction with ` + "`" + `health_check_path` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            },
            "health_check_path": {
              "description": "The path to be checked for this function app health.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "http2_enabled": {
              "description": "Specifies if the http2 protocol should be enabled. Defaults to ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "ip_restriction_default_action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "load_balancing_mode": {
              "description": "The Site load balancing mode. Possible values include: ` + "`" + `WeightedRoundRobin` + "`" + `, ` + "`" + `LeastRequests` + "`" + `, ` + "`" + `LeastResponseTime` + "`" + `, ` + "`" + `WeightedTotalTraffic` + "`" + `, ` + "`" + `RequestHash` + "`" + `, ` + "`" + `PerSiteRoundRobin` + "`" + `. Defaults to ` + "`" + `LeastRequests` + "`" + ` if omitted.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "managed_pipeline_mode": {
              "description": "The Managed Pipeline mode. Possible values include: ` + "`" + `Integrated` + "`" + `, ` + "`" + `Classic` + "`" + `. Defaults to ` + "`" + `Integrated` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "minimum_tls_version": {
              "description": "The configures the minimum version of TLS required for SSL requests. Possible values include: ` + "`" + `1.0` + "`" + `, ` + "`" + `1.1` + "`" + `, and  ` + "`" + `1.2` + "`" + `. Defaults to ` + "`" + `1.2` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "remote_debugging_enabled": {
              "description": "Should Remote Debugging be enabled. Defaults to ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "remote_debugging_version": {
              "computed": true,
              "description": "The Remote Debugging Version. Possible values include ` + "`" + `VS2017` + "`" + `, ` + "`" + `VS2019` + "`" + `, and ` + "`" + `VS2022` + "`" + `` + "`" + `",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "runtime_scale_monitoring_enabled": {
              "description": "Should Functions Runtime Scale Monitoring be enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "scm_ip_restriction_default_action": {
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scm_minimum_tls_version": {
              "description": "Configures the minimum version of TLS required for SSL requests to the SCM site Possible values include: ` + "`" + `1.0` + "`" + `, ` + "`" + `1.1` + "`" + `, and  ` + "`" + `1.2` + "`" + `. Defaults to ` + "`" + `1.2` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "scm_type": {
              "computed": true,
              "description": "The SCM Type in use by the Linux Function App.",
              "description_kind": "plain",
              "type": "string"
            },
            "scm_use_main_ip_restriction": {
              "description": "Should the Linux Function App ` + "`" + `ip_restriction` + "`" + ` configuration be used for the SCM also.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "use_32_bit_worker": {
              "description": "Should the Linux Function App use a 32-bit worker.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "vnet_route_all_enabled": {
              "description": "Should the Linux Function App route all traffic through the virtual network.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "websockets_enabled": {
              "description": "Should Web Sockets be enabled. Defaults to ` + "`" + `false` + "`" + `.",
              "description_kind": "plain",
              "optional": true,
              "type": "bool"
            },
            "worker_count": {
              "computed": true,
              "description": "The number of Workers for this Linux Function App.",
              "description_kind": "plain",
              "optional": true,
              "type": "number"
            }
          },
          "block_types": {
            "app_service_logs": {
              "block": {
                "attributes": {
                  "disk_quota_mb": {
                    "description": "The amount of disk space to use for logs. Valid values are between ` + "`" + `25` + "`" + ` and ` + "`" + `100` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "retention_period_days": {
                    "description": "The retention period for logs in days. Valid values are between ` + "`" + `0` + "`" + ` and ` + "`" + `99999` + "`" + `. Defaults to ` + "`" + `0` + "`" + ` (never delete).",
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
            "cors": {
              "block": {
                "attributes": {
                  "allowed_origins": {
                    "description": "Specifies a list of origins that should be allowed to make cross-origin calls.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "set",
                      "string"
                    ]
                  },
                  "support_credentials": {
                    "description": "Are credentials allowed in CORS requests? Defaults to ` + "`" + `false` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "bool"
                  }
                },
                "description_kind": "plain"
              },
              "max_items": 1,
              "nesting_mode": "list"
            },
            "ip_restriction": {
              "block": {
                "attributes": {
                  "action": {
                    "description": "The action to take. Possible values are ` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "description": "The description of the IP restriction rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "headers": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "x_azure_fdid": [
                            "list",
                            "string"
                          ],
                          "x_fd_health_probe": [
                            "list",
                            "string"
                          ],
                          "x_forwarded_for": [
                            "list",
                            "string"
                          ],
                          "x_forwarded_host": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  },
                  "ip_address": {
                    "description": "The CIDR notation of the IP or IP Range to match. For example: ` + "`" + `10.0.0.0/24` + "`" + ` or ` + "`" + `192.168.10.1/32` + "`" + ` or ` + "`" + `fe80::/64` + "`" + ` or ` + "`" + `13.107.6.152/31,13.107.128.0/22` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name which should be used for this ` + "`" + `ip_restriction` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description": "The priority value of this ` + "`" + `ip_restriction` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "service_tag": {
                    "description": "The Service Tag used for this IP Restriction.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "virtual_network_subnet_id": {
                    "description": "The Virtual Network Subnet ID used for this IP Restriction.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  }
                },
                "description_kind": "plain"
              },
              "nesting_mode": "list"
            },
            "scm_ip_restriction": {
              "block": {
                "attributes": {
                  "action": {
                    "description": "The action to take. Possible values are ` + "`" + `Allow` + "`" + ` or ` + "`" + `Deny` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "description": {
                    "description": "The description of the IP restriction rule.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "headers": {
                    "description_kind": "plain",
                    "optional": true,
                    "type": [
                      "list",
                      [
                        "object",
                        {
                          "x_azure_fdid": [
                            "list",
                            "string"
                          ],
                          "x_fd_health_probe": [
                            "list",
                            "string"
                          ],
                          "x_forwarded_for": [
                            "list",
                            "string"
                          ],
                          "x_forwarded_host": [
                            "list",
                            "string"
                          ]
                        }
                      ]
                    ]
                  },
                  "ip_address": {
                    "description": "The CIDR notation of the IP or IP Range to match. For example: ` + "`" + `10.0.0.0/24` + "`" + ` or ` + "`" + `192.168.10.1/32` + "`" + ` or ` + "`" + `fe80::/64` + "`" + ` or ` + "`" + `13.107.6.152/31,13.107.128.0/22` + "`" + `",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "name": {
                    "computed": true,
                    "description": "The name which should be used for this ` + "`" + `ip_restriction` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "priority": {
                    "description": "The priority value of this ` + "`" + `ip_restriction` + "`" + `.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "number"
                  },
                  "service_tag": {
                    "description": "The Service Tag used for this IP Restriction.",
                    "description_kind": "plain",
                    "optional": true,
                    "type": "string"
                  },
                  "virtual_network_subnet_id": {
                    "description": "The Virtual Network Subnet ID used for this IP Restriction.",
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
      "sticky_settings": {
        "block": {
          "attributes": {
            "app_setting_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            },
            "connection_string_names": {
              "description_kind": "plain",
              "optional": true,
              "type": [
                "list",
                "string"
              ]
            }
          },
          "description_kind": "plain"
        },
        "max_items": 1,
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

func AzurermFunctionAppFlexConsumptionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azurermFunctionAppFlexConsumption), &result)
	return &result
}
