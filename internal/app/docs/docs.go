// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/Drakenchef",
            "email": "drakenchef@gmail.com"
        },
        "license": {
            "name": "AS IS (NO WARRANTY)"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/planets": {
            "get": {
                "description": "Get a list of planets with optional filtering by planet name.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "planets"
                ],
                "summary": "Get a list of planets",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Planet name for filtering",
                        "name": "Марс",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ds.Planet"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ds.Planet": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "is_delete": {
                    "type": "boolean"
                },
                "radius": {
                    "type": "number"
                },
                "distance": {
                    "type": "number"
                },
                "gravity": {
                    "type": "number"
                },
				"type": {
					"type": "string"
				}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "AMS",
	Description:      "AMS flights",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
