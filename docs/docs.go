// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/health": {
            "get": {
                "description": "Checks the health status of the server and database connection",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "health"
                ],
                "summary": "Health endpoint",
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Could not establish a connection to the database",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/members": {
            "get": {
                "description": "Get all members with tags",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "members"
                ],
                "summary": "Get all members with tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Member"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new member",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "members"
                ],
                "summary": "Add a new member",
                "parameters": [
                    {
                        "description": "Member object that needs to be added",
                        "name": "member",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    }
                }
            }
        },
        "/members/{id}": {
            "get": {
                "description": "Get a member by ID with tags",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "members"
                ],
                "summary": "Get a member by ID with tags",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Member ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Member"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a member by ID",
                "tags": [
                    "members"
                ],
                "summary": "Delete a member",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Member ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Member deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Checks the availability of the server",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "Ping endpoint",
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/tags": {
            "get": {
                "description": "Gets all tags",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tags"
                ],
                "summary": "Get all tags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Tag"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Member": {
            "type": "object",
            "required": [
                "name",
                "type"
            ],
            "properties": {
                "contractDuration": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Tag"
                    }
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "EMPLOYEE",
                        "CONTRACTOR"
                    ]
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.Tag": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "members": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Member"
                    }
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
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
