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
        "/tasks/start/{id}": {
            "post": {
                "description": "Starts a task for a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Start Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routers.TaskID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task started successfully",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID, Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "409": {
                        "description": "Task already started",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to start task",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    }
                }
            }
        },
        "/tasks/stop/{id}": {
            "post": {
                "description": "Stops a task for a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Stop Task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Task ID",
                        "name": "task_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routers.TaskID"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task stopped successfully",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID, Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "409": {
                        "description": "Task not started",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to stop task",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    }
                }
            }
        },
        "/tasks/{id}": {
            "get": {
                "description": "Retrieves a list of tasks for a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tasks"
                ],
                "summary": "Get Tasks",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of tasks",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/database.Task"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "404": {
                        "description": "Tasks not found",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to get tasks",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieves a list of users with optional filters",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get Users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Passport Series",
                        "name": "passport_serie",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Passport Number",
                        "name": "passport_number",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Last Name",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "First Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "$ref": "#/definitions/database.User"
                        }
                    },
                    "400": {
                        "description": "Invalid limit, Invalid page",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "404": {
                        "description": "Users not found",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to get users",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new user to the database using passport series and number",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Add a new user",
                "parameters": [
                    {
                        "description": "Passport number",
                        "name": "passportNumber",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/routers.Passport"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User added successfully",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload, Invalid passport number",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to fetch user data from external API, Failed to add user to the database",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "put": {
                "description": "Updates a user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User data to update",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated successfully",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "400": {
                        "description": "No fields to update, Error print",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to update user",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    },
                    "500": {
                        "description": "Failed to delete user",
                        "schema": {
                            "$ref": "#/definitions/routers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "database.Task": {
            "type": "object",
            "properties": {
                "duration": {
                    "type": "string"
                },
                "end_time": {
                    "type": "string"
                },
                "start_time": {
                    "type": "string"
                },
                "task_id": {
                    "type": "integer"
                }
            }
        },
        "database.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "passport_number": {
                    "type": "string"
                },
                "passport_serie": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                }
            }
        },
        "routers.Passport": {
            "type": "object",
            "properties": {
                "passportNumber": {
                    "type": "string"
                }
            }
        },
        "routers.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "routers.TaskID": {
            "type": "object",
            "properties": {
                "task_id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Time tracker",
	Description:      "This is an example of a time tracking API..",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
