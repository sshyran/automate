package api

func init() {
	Swagger.Add("iam_v2beta_rules", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/iam/v2beta/rules.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/iam/v2beta/rules": {
      "get": {
        "operationId": "ListRules",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaListRulesResp"
            }
          }
        },
        "tags": [
          "Rules"
        ]
      },
      "post": {
        "operationId": "CreateRule",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaCreateRuleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/v2betaCreateRuleReq"
            }
          }
        ],
        "tags": [
          "Rules"
        ]
      }
    },
    "/iam/v2beta/rules/{id}": {
      "get": {
        "operationId": "GetRule",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaGetRuleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Rules"
        ]
      },
      "delete": {
        "operationId": "DeleteRule",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/v2betaDeleteRuleResp"
            }
          }
        },
        "parameters": [
          {
            "name": "id",
            "in": "path",
            "required": true,
            "type": "string"
          }
        ],
        "tags": [
          "Rules"
        ]
      }
    }
  },
  "definitions": {
    "v2betaCondition": {
      "type": "object",
      "properties": {
        "type": {
          "$ref": "#/definitions/v2betaConditionType"
        },
        "values": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "operator": {
          "$ref": "#/definitions/v2betaConditionOperator"
        }
      }
    },
    "v2betaConditionOperator": {
      "type": "string",
      "enum": [
        "CONDITION_OPERATOR_UNSET",
        "MEMBER_OF",
        "EQUALS"
      ],
      "default": "CONDITION_OPERATOR_UNSET"
    },
    "v2betaConditionType": {
      "type": "string",
      "enum": [
        "CONDITION_TYPE_UNSET",
        "CHEF_SERVERS",
        "CHEF_ORGS",
        "CHEF_ENVIRONMENTS",
        "ROLES",
        "CHEF_TAGS",
        "POLICY_GROUP",
        "POLICY_NAME"
      ],
      "default": "CONDITION_TYPE_UNSET"
    },
    "v2betaCreateRuleReq": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "project_id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "$ref": "#/definitions/v2betaRuleType"
        },
        "conditions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v2betaCondition"
          }
        }
      }
    },
    "v2betaCreateRuleResp": {
      "type": "object",
      "properties": {
        "rule": {
          "$ref": "#/definitions/v2betaRule"
        }
      }
    },
    "v2betaDeleteRuleResp": {
      "type": "object"
    },
    "v2betaGetRuleResp": {
      "type": "object",
      "properties": {
        "rule": {
          "$ref": "#/definitions/v2betaRule"
        }
      }
    },
    "v2betaListRulesResp": {
      "type": "object",
      "properties": {
        "rules": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v2betaRule"
          }
        }
      }
    },
    "v2betaRule": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        },
        "project_id": {
          "type": "string"
        },
        "name": {
          "type": "string"
        },
        "type": {
          "$ref": "#/definitions/v2betaRuleType"
        },
        "conditions": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/v2betaCondition"
          }
        }
      }
    },
    "v2betaRuleType": {
      "type": "string",
      "enum": [
        "RULE_TYPE_UNSET",
        "NODE",
        "EVENT"
      ],
      "default": "RULE_TYPE_UNSET"
    }
  }
}
`)
}