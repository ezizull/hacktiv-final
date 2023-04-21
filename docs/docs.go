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
        "/auth/access-token": {
            "post": {
                "description": "Auth user by email and password",
                "tags": [
                    "auth"
                ],
                "summary": "GetAccessTokenByRefreshToken UserName",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.AccessTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.SecurityAuthenticatedUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "Auth user by email and password",
                "tags": [
                    "auth"
                ],
                "summary": "Login UserName",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.SecurityAuthenticatedUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    }
                }
            }
        },
        "/book": {
            "get": {
                "description": "Get all Books on the system",
                "tags": [
                    "book"
                ],
                "summary": "Get all Books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/book.PaginationResultBook"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
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
                    "book"
                ],
                "summary": "Create New Book",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book.NewBookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/hacktiv_final-project_domain_book.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    }
                }
            }
        },
        "/book/{book_id}": {
            "get": {
                "tags": [
                    "book"
                ],
                "summary": "Get books by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of book",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/hacktiv_final-project_domain_book.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all Users on the system",
                "tags": [
                    "user"
                ],
                "summary": "Get all Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/user.ResponseUser"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create new user on the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create New UserName",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.NewUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.ResponseUser"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    }
                }
            }
        },
        "/user/{user_id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get Users by ID on the system",
                "tags": [
                    "user"
                ],
                "summary": "Get users by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of user",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.AccessTokenRequest": {
            "type": "object",
            "required": [
                "refreshToken"
            ],
            "properties": {
                "refreshToken": {
                    "type": "string",
                    "example": "badbunybabybebe"
                }
            }
        },
        "auth.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "mail@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "Password123"
                }
            }
        },
        "book.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "book.NewBookRequest": {
            "type": "object",
            "required": [
                "description",
                "title"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Something"
                },
                "title": {
                    "type": "string",
                    "example": "Book"
                }
            }
        },
        "book.PaginationResultBook": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/hacktiv_final-project_domain_book.Book"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "nextCursor": {
                    "type": "integer"
                },
                "numPages": {
                    "type": "integer"
                },
                "prevCursor": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "controllers.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "hacktiv_final-project_domain_book.Book": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "example": "book description"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "example": "book title"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "user.DataSecurityAuthenticated": {
            "type": "object",
            "properties": {
                "expirationAccessDateTime": {
                    "type": "string",
                    "example": "2023-02-02T21:03:53.196419-06:00"
                },
                "expirationRefreshDateTime": {
                    "type": "string",
                    "example": "2023-02-03T06:53:53.196419-06:00"
                },
                "jwtAccessToken": {
                    "type": "string",
                    "example": "SomeAccessToken"
                },
                "jwtRefreshToken": {
                    "type": "string",
                    "example": "SomeRefreshToken"
                }
            }
        },
        "user.DataUserAuthenticated": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "some@mail.com"
                },
                "id": {
                    "type": "integer",
                    "example": 123
                },
                "role": {
                    "$ref": "#/definitions/user.Role"
                },
                "role_id": {
                    "type": "string",
                    "example": "admin"
                },
                "userName": {
                    "type": "string",
                    "example": "UserName"
                }
            }
        },
        "user.NewUser": {
            "type": "object",
            "required": [
                "age",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 1
                },
                "email": {
                    "type": "string",
                    "example": "mail@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "Password123"
                },
                "role_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "example": "someUser"
                }
            }
        },
        "user.ResponseUser": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string",
                    "example": "2021-02-24 20:19:39"
                },
                "email": {
                    "type": "string",
                    "example": "some@mail.com"
                },
                "firstName": {
                    "type": "string",
                    "example": "John"
                },
                "id": {
                    "type": "integer",
                    "example": 1099
                },
                "lastName": {
                    "type": "string",
                    "example": "Doe"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2021-02-24 20:19:39"
                },
                "user": {
                    "type": "string",
                    "example": "BossonH"
                }
            }
        },
        "user.Role": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "user.SecurityAuthenticatedUser": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/user.DataUserAuthenticated"
                },
                "security": {
                    "$ref": "#/definitions/user.DataSecurityAuthenticated"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
