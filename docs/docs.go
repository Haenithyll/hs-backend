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
        "/api/facets": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Returns facets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Facets"
                ],
                "summary": "Get facets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetFacetResponseItem"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new facet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Facets"
                ],
                "summary": "Create facet",
                "parameters": [
                    {
                        "description": "Facet",
                        "name": "facet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateFacetInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateFacetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/facets/{facetId}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes a facet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Facets"
                ],
                "summary": "Delete facet",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Facet ID",
                        "name": "facetId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updates a facet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Facets"
                ],
                "summary": "Update facet",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Facet ID",
                        "name": "facetId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Facet",
                        "name": "facet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateFacetInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateFacetResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/prisms": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new prism",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Prisms"
                ],
                "summary": "Create prism",
                "parameters": [
                    {
                        "description": "Prism",
                        "name": "prism",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePrismInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreatePrismResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Returns user info by email query param",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email address",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetUserByEmailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/communication-services": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Returns user communication services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Communication Services"
                ],
                "summary": "Get user communication services",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.GetUserCommunicationServicesResponseItem"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Creates a new user communication service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Communication Services"
                ],
                "summary": "Create user communication service",
                "parameters": [
                    {
                        "description": "User Communication Service",
                        "name": "userCommunicationService",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserCommunicationServiceInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserCommunicationServiceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/communication-services/{userCommunicationServiceId}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes a user communication service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Communication Services"
                ],
                "summary": "Delete user communication service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Communication Service ID",
                        "name": "userCommunicationServiceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Updates a user communication service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User Communication Services"
                ],
                "summary": "Update user communication service",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User Communication Service ID",
                        "name": "userCommunicationServiceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Communication Service",
                        "name": "userCommunicationService",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserCommunicationServiceInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserCommunicationServiceResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/users/me": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Returns user me",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user me",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetUserMeResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/error.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateFacetEnrichedConfig": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreateFacetEnrichedConfigItem"
                    }
                }
            }
        },
        "dto.CreateFacetEnrichedConfigItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/enum.CommunicationService"
                },
                "status": {
                    "$ref": "#/definitions/enum.FacetStatus"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.CreateFacetInput": {
            "type": "object",
            "required": [
                "color",
                "configuration",
                "privateLabel",
                "publicLabel"
            ],
            "properties": {
                "color": {
                    "type": "string"
                },
                "configuration": {
                    "$ref": "#/definitions/json.FacetConfig"
                },
                "privateLabel": {
                    "type": "string"
                },
                "publicLabel": {
                    "type": "string"
                }
            }
        },
        "dto.CreateFacetResponse": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "configuration": {
                    "$ref": "#/definitions/dto.CreateFacetEnrichedConfig"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "privateLabel": {
                    "type": "string"
                },
                "publicLabel": {
                    "type": "string"
                }
            }
        },
        "dto.CreatePrismEnrichedConfig": {
            "type": "object",
            "properties": {
                "base": {
                    "$ref": "#/definitions/dto.CreatePrismFacet"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.CreatePrismEnrichedUserItem"
                    }
                }
            }
        },
        "dto.CreatePrismEnrichedUserItem": {
            "type": "object",
            "properties": {
                "facet": {
                    "$ref": "#/definitions/dto.CreatePrismFacet"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "dto.CreatePrismFacet": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "privateLabel": {
                    "type": "string"
                },
                "publicLabel": {
                    "type": "string"
                }
            }
        },
        "dto.CreatePrismInput": {
            "type": "object",
            "required": [
                "configuration",
                "name"
            ],
            "properties": {
                "configuration": {
                    "$ref": "#/definitions/json.PrismConfig"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CreatePrismResponse": {
            "type": "object",
            "properties": {
                "configuration": {
                    "$ref": "#/definitions/dto.CreatePrismEnrichedConfig"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserCommunicationServiceInput": {
            "type": "object",
            "required": [
                "name",
                "service",
                "value"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/enum.CommunicationService"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserCommunicationServiceResponse": {
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
                "service": {
                    "$ref": "#/definitions/enum.CommunicationService"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.GetFacetEnrichedConfig": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.GetFacetEnrichedConfigItem"
                    }
                }
            }
        },
        "dto.GetFacetEnrichedConfigItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/enum.CommunicationService"
                },
                "status": {
                    "$ref": "#/definitions/enum.FacetStatus"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.GetFacetResponseItem": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "configuration": {
                    "$ref": "#/definitions/dto.GetFacetEnrichedConfig"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "privateLabel": {
                    "type": "string"
                },
                "publicLabel": {
                    "type": "string"
                }
            }
        },
        "dto.GetUserByEmailResponse": {
            "type": "object",
            "properties": {
                "avatarUrl": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "dto.GetUserCommunicationServicesResponseItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/enum.CommunicationService"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.GetUserMeResponse": {
            "type": "object",
            "properties": {
                "avatarUrl": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateFacetEnrichedConfig": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.UpdateFacetEnrichedConfigItem"
                    }
                }
            }
        },
        "dto.UpdateFacetEnrichedConfigItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/enum.CommunicationService"
                },
                "status": {
                    "$ref": "#/definitions/enum.FacetStatus"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateFacetInput": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "configuration": {
                    "$ref": "#/definitions/json.FacetConfig"
                },
                "privateLabel": {
                    "type": "string"
                },
                "publicLabel": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateFacetResponse": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "configuration": {
                    "$ref": "#/definitions/dto.UpdateFacetEnrichedConfig"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "privateLabel": {
                    "type": "string"
                },
                "publicLabel": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserCommunicationServiceInput": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/enum.CommunicationService"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserCommunicationServiceResponse": {
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
                "service": {
                    "$ref": "#/definitions/enum.CommunicationService"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "enum.CommunicationService": {
            "type": "string",
            "enum": [
                "faceToFace",
                "phone",
                "message",
                "email",
                "discord",
                "microsoftTeams"
            ],
            "x-enum-varnames": [
                "FaceToFace",
                "Phone",
                "Message",
                "Email",
                "Discord",
                "MicrosoftTeams"
            ]
        },
        "enum.FacetStatus": {
            "type": "string",
            "enum": [
                "available",
                "emergencyOnly"
            ],
            "x-enum-varnames": [
                "Available",
                "EmergencyOnly"
            ]
        },
        "error.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "json.FacetConfig": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/json.FacetConfigItem"
                    }
                }
            }
        },
        "json.FacetConfigItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/enum.FacetStatus"
                }
            }
        },
        "json.PrismConfig": {
            "type": "object",
            "properties": {
                "base": {
                    "type": "integer"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/json.PrismConfigUserItem"
                    }
                }
            }
        },
        "json.PrismConfigUserItem": {
            "type": "object",
            "properties": {
                "facetId": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and your token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "HS Backend API",
	Description:      "This is the backend API for the HS project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
