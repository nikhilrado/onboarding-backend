// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "WSO Devs",
            "url": "https://wso.williams.edu/",
            "email": "wso-dev@wso.williams.edu"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/admin/edit/": {
            "post": {
                "security": [
                    {
                        "BasicAuth": []
                    }
                ],
                "description": "Parse the input body json, add/edit it in fakedatabase, and then write the responses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "edits the currently logged in user in fakedatabse",
                "parameters": [
                    {
                        "description": "user value to edit",
                        "name": "equest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.EditResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.EditResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{name}": {
            "get": {
                "description": "Parse the input path param name, query it in fakedatabase, and then write the responses based on whether user exists in database",
                "produces": [
                    "application/json"
                ],
                "summary": "lists the value associated with user from fakedatabase",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name of the user",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/model.UserInfo"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "sanity check, return \"Hello World\"",
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.EditResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "model.UserInfo": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string",
                    "example": "Blue"
                },
                "user": {
                    "type": "string",
                    "example": "Ye"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "WSO Backend Onboarding Project",
	Description:      "This is a sample project for onboarding WSO members.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
