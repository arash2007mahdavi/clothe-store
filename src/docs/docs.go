// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "this welcome message is for testing api",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Welcome"
                ],
                "summary": "welcome to store",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/clothes": {
            "get": {
                "description": "get clothes information (price and amount)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clothes"
                ],
                "summary": "get clothes",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/clothes/buy": {
            "post": {
                "description": "buy clothe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Clothes"
                ],
                "summary": "buy clothe",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/profile/charge/wallet": {
            "post": {
                "description": "charge wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "charge wallet",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/profile/new": {
            "post": {
                "description": "create new profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "create new profile",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/profile/see": {
            "get": {
                "description": "watch your profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "watch profile",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        },
        "/profile/see/all": {
            "get": {
                "description": "watch all of profiles",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "watch profiles",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/responses.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "responses.Response": {
            "type": "object",
            "properties": {
                "Error": {
                    "type": "string"
                },
                "Response": {},
                "Status": {
                    "type": "boolean"
                },
                "StatusCode": {
                    "type": "integer"
                },
                "ValidationError": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/validation.Validationerror"
                    }
                }
            }
        },
        "validation.Validationerror": {
            "type": "object",
            "properties": {
                "property": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
