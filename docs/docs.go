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
        "/cars": {
            "get": {
                "description": "GET information about a car by its registration number",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "GET car by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Registration Number",
                        "name": "regNum",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Mark of car",
                        "name": "mark",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "model of car",
                        "name": "model",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Year of car",
                        "name": "year",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Lower limit of car definition",
                        "name": "lyear",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Top limit of car definition",
                        "name": "tyear",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Owner name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Owner surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Owner patronymic",
                        "name": "patronymic",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page ",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of elements per page",
                        "name": "pagesize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Information of cars by using filter",
                        "schema": {
                            "$ref": "#/definitions/service.SuccessGetResponse"
                        }
                    },
                    "400": {
                        "description": "Incorrect filtering attributes",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Post information about a car by its registration number by using external API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Post car by id in external API",
                "parameters": [
                    {
                        "description": "Request data",
                        "name": "regNums",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/service.ReqBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "All cars successfully added or some cars were added, but some registration numbers were invalid.",
                        "schema": {
                            "$ref": "#/definitions/service.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - All provided registration numbers were invalid."
                    }
                }
            }
        },
        "/cars/{regnum}": {
            "delete": {
                "description": "Delete information about a car by its registration number",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Delete car id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car Registration Number",
                        "name": "regnum",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Successful deletion, no content returned"
                    },
                    "404": {
                        "description": "Resource not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            },
            "patch": {
                "description": "Update information about a car by its registration number",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "cars"
                ],
                "summary": "Patch car by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Car Registration Number",
                        "name": "regnum",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Patch data",
                        "name": "jsonfile",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.Patch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Patch request was successful"
                    },
                    "400": {
                        "description": "Bad Request - wrong JSON format or update request failed"
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Catalog": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/models.Person"
                },
                "regNum": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.Patch": {
            "type": "object",
            "properties": {
                "mark": {
                    "type": "string"
                },
                "model": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "patronymic": {
                    "type": "string"
                },
                "regNum": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "models.Person": {
            "type": "object",
            "properties": {
                "name": {
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
        "service.ReqBody": {
            "type": "object",
            "properties": {
                "regNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "service.SuccessGetResponse": {
            "type": "object",
            "properties": {
                "cars": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Catalog"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "service.SuccessResponse": {
            "type": "object",
            "properties": {
                "badRegNums": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "errorResponse": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9092",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Car info API",
	Description:      "Car info service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
