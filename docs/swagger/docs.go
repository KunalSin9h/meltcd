// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/meltred/meltcd/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "General"
                ],
                "summary": "Check server status",
                "responses": {}
            }
        },
        "/apps": {
            "get": {
                "tags": [
                    "Apps"
                ],
                "summary": "Get a list all applications created",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.AppList"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Apps"
                ],
                "summary": "Update an application",
                "parameters": [
                    {
                        "description": "Application body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/application.Application"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Apps"
                ],
                "summary": "Create a new application",
                "parameters": [
                    {
                        "description": "Application body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/application.Application"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            }
        },
        "/apps/{app_name}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Apps"
                ],
                "summary": "Get details of an application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application name",
                        "name": "app_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/application.Application"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "Apps"
                ],
                "summary": "Remove an application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application name",
                        "name": "app_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            }
        },
        "/apps/{app_name}/recreate": {
            "post": {
                "tags": [
                    "Apps"
                ],
                "summary": "Recreate application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application name",
                        "name": "app_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            }
        },
        "/apps/{app_name}/refresh": {
            "post": {
                "tags": [
                    "Apps"
                ],
                "summary": "Refresh/Synchronize an application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application name",
                        "name": "app_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "tags": [
                    "General"
                ],
                "summary": "Login user",
                "responses": {
                    "302": {
                        "description": "Found"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        },
        "/repo": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Repo"
                ],
                "summary": "Get a list all repositories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/repo.ListData"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Repo"
                ],
                "summary": "Update a repository",
                "parameters": [
                    {
                        "description": "Repository details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repo.PrivateRepoDetails"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Repo"
                ],
                "summary": "Add a new repository",
                "parameters": [
                    {
                        "description": "Repository details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repo.PrivateRepoDetails"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Repo"
                ],
                "summary": "Remove a repository",
                "parameters": [
                    {
                        "description": "Repository url",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/repo.RemovePayload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.GlobalResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "tags": [
                    "Users"
                ],
                "summary": "Get all the users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.AllUsers"
                        }
                    }
                }
            }
        },
        "/users/current": {
            "get": {
                "tags": [
                    "Users"
                ],
                "summary": "Get username of current logged-in user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized"
                    }
                }
            }
        }
    },
    "definitions": {
        "app.GlobalResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "application.Application": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "health": {
                    "$ref": "#/definitions/application.Health"
                },
                "health_status": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_synced_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "refresh_timer": {
                    "description": "Timer to check for Sync format of \"3m50s\"",
                    "type": "string"
                },
                "source": {
                    "$ref": "#/definitions/application.Source"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "application.Health": {
            "type": "integer",
            "enum": [
                0,
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "Healthy",
                "Progressing",
                "Degraded",
                "Suspended"
            ]
        },
        "application.Source": {
            "type": "object",
            "properties": {
                "path": {
                    "type": "string"
                },
                "repoURL": {
                    "type": "string"
                },
                "targetRevision": {
                    "type": "string"
                }
            }
        },
        "auth.AllUsers": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/auth.User"
                    }
                }
            }
        },
        "auth.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "lastLoggedIn": {
                    "type": "string"
                },
                "passwordHash": {
                    "description": "hash passwords",
                    "type": "string"
                },
                "rol": {
                    "$ref": "#/definitions/auth.UserRole"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.UserRole": {
            "type": "string",
            "enum": [
                "admin",
                "general"
            ],
            "x-enum-varnames": [
                "Admin",
                "General"
            ]
        },
        "core.AppList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/core.AppStatus"
                    }
                }
            }
        },
        "core.AppStatus": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "health": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_synced_at": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "repo.ListData": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "repo.PrivateRepoDetails": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "repo.RemovePayload": {
            "type": "object",
            "properties": {
                "repo": {
                    "type": "string"
                }
            }
        }
    },
    "externalDocs": {
        "description": "Meltcd Docs",
        "url": "https://cd.meltred.tech/docs"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.6",
	Host:             "localhost:11771",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Meltcd API",
	Description:      "Argo-cd like GitDevOps Continuous Development platform for docker swarm.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
