package api

func init() {
	Swagger.Add("compliance_profiles_profiles", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/compliance/profiles/profiles.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/compliance/market/read/{name}/version/{version}": {
      "get": {
        "operationId": "ReadFromMarket",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Profile"
            }
          }
        },
        "parameters": [
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "owner",
            "in": "query",
            "required": false,
            "type": "string"
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    },
    "/compliance/profiles/read/{owner}/{name}/version/{version}": {
      "get": {
        "operationId": "Read",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Profile"
            }
          }
        },
        "parameters": [
          {
            "name": "owner",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    },
    "/compliance/profiles/search": {
      "post": {
        "operationId": "List",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Profiles"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Query"
            }
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    },
    "/compliance/profiles/{owner}/{name}/version/{version}": {
      "delete": {
        "operationId": "Delete",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "properties": {}
            }
          }
        },
        "parameters": [
          {
            "name": "owner",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "version",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "ProfilesService"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.compliance.profiles.v1.Attribute": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "options": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Option"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.CheckMessage": {
      "type": "object",
      "properties": {
        "file": {
          "type": "string"
        },
        "line": {
          "type": "integer",
          "format": "int32"
        },
        "column": {
          "type": "integer",
          "format": "int32"
        },
        "control_id": {
          "type": "string"
        },
        "msg": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.CheckResult": {
      "type": "object",
      "properties": {
        "summary": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.ResultSummary"
        },
        "errors": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.CheckMessage"
          }
        },
        "warnings": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.CheckMessage"
          }
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Chunk": {
      "type": "object",
      "properties": {
        "data": {
          "type": "string",
          "format": "byte"
        },
        "position": {
          "type": "string",
          "format": "int64"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Control": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "code": {
          "type": "string"
        },
        "desc": {
          "type": "string"
        },
        "impact": {
          "type": "number",
          "format": "float"
        },
        "title": {
          "type": "string"
        },
        "source_location": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.SourceLocation"
        },
        "results": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Result"
          }
        },
        "refs": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Ref"
          }
        },
        "tags": {
          "type": "object",
          "additionalProperties": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Dependency": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "url": {
          "type": "string"
        },
        "path": {
          "type": "string"
        },
        "git": {
          "type": "string"
        },
        "branch": {
          "type": "string"
        },
        "tag": {
          "type": "string"
        },
        "commit": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "supermarket": {
          "type": "string"
        },
        "github": {
          "type": "string"
        },
        "compliance": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Group": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "controls": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.ListFilter": {
      "type": "object",
      "properties": {
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "type": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Metadata": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "content_type": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Option": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string"
        },
        "default": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Profile": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "title": {
          "type": "string"
        },
        "maintainer": {
          "type": "string"
        },
        "copyright": {
          "type": "string"
        },
        "copyright_email": {
          "type": "string"
        },
        "license": {
          "type": "string"
        },
        "summary": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "owner": {
          "type": "string"
        },
        "supports": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Support"
          }
        },
        "depends": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Dependency"
          }
        },
        "sha256": {
          "type": "string"
        },
        "groups": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Group"
          }
        },
        "controls": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Control"
          }
        },
        "attributes": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Attribute"
          }
        },
        "latest_version": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.ProfileData": {
      "type": "object",
      "properties": {
        "owner": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "data": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Profiles": {
      "type": "object",
      "properties": {
        "profiles": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Profile"
          }
        },
        "total": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Query": {
      "type": "object",
      "properties": {
        "filters": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.ListFilter"
          }
        },
        "order": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.Query.OrderType"
        },
        "sort": {
          "type": "string"
        },
        "page": {
          "type": "integer",
          "format": "int32"
        },
        "per_page": {
          "type": "integer",
          "format": "int32"
        },
        "owner": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "version": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Query.OrderType": {
      "type": "string",
      "enum": [
        "ASC",
        "DESC"
      ],
      "default": "ASC"
    },
    "chef.automate.api.compliance.profiles.v1.Ref": {
      "type": "object",
      "properties": {
        "url": {
          "type": "string"
        },
        "ref": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Result": {
      "type": "object",
      "properties": {
        "status": {
          "type": "string"
        },
        "code_desc": {
          "type": "string"
        },
        "run_time": {
          "type": "number",
          "format": "float"
        },
        "start_time": {
          "type": "string"
        },
        "message": {
          "type": "string"
        },
        "skip_message": {
          "type": "string"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.ResultSummary": {
      "type": "object",
      "properties": {
        "valid": {
          "type": "boolean",
          "format": "boolean"
        },
        "timestamp": {
          "type": "string"
        },
        "location": {
          "type": "string"
        },
        "controls": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.SourceLocation": {
      "type": "object",
      "properties": {
        "ref": {
          "type": "string"
        },
        "line": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "chef.automate.api.compliance.profiles.v1.Support": {
      "type": "object",
      "properties": {
        "os_name": {
          "type": "string"
        },
        "os_family": {
          "type": "string"
        },
        "release": {
          "type": "string"
        },
        "inspec_version": {
          "type": "string"
        },
        "platform": {
          "type": "string"
        }
      }
    },
    "google.protobuf.Any": {
      "type": "object",
      "properties": {
        "type_url": {
          "type": "string"
        },
        "value": {
          "type": "string",
          "format": "byte"
        }
      }
    },
    "grpc.gateway.runtime.StreamError": {
      "type": "object",
      "properties": {
        "grpc_code": {
          "type": "integer",
          "format": "int32"
        },
        "http_code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "http_status": {
          "type": "string"
        },
        "details": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/google.protobuf.Any"
          }
        }
      }
    }
  },
  "x-stream-definitions": {
    "chef.automate.api.compliance.profiles.v1.ProfileData": {
      "type": "object",
      "properties": {
        "result": {
          "$ref": "#/definitions/chef.automate.api.compliance.profiles.v1.ProfileData"
        },
        "error": {
          "$ref": "#/definitions/grpc.gateway.runtime.StreamError"
        }
      },
      "title": "Stream result of chef.automate.api.compliance.profiles.v1.ProfileData"
    }
  }
}
`)
}
