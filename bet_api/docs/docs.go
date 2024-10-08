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
        "/apuestas": {
            "get": {
                "description": "Obtiene todas las apuestas realizadas en el casino",
                "produces": [
                    "application/json"
                ],
                "summary": "Listar apuestas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Apuesta"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Registra una nueva apuesta en el sistema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Crear una apuesta",
                "parameters": [
                    {
                        "description": "Datos de la apuesta a registrar",
                        "name": "apuesta",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Apuesta"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.Apuesta"
                        }
                    },
                    "400": {
                        "description": "Error al procesar la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/juegos": {
            "get": {
                "description": "Obtiene todos los juegos disponibles en el casino",
                "produces": [
                    "application/json"
                ],
                "summary": "Listar juegos",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/main.Juego"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Crea un nuevo juego en el casino",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Crear un juego",
                "parameters": [
                    {
                        "description": "Datos del juego",
                        "name": "juego",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Juego"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/main.Juego"
                        }
                    },
                    "400": {
                        "description": "Error al procesar la solicitud",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/juegos/{id}": {
            "delete": {
                "description": "Elimina un juego por su ID",
                "summary": "Eliminar un juego",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID del juego a eliminar",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "El juego ha sido eliminado"
                    },
                    "404": {
                        "description": "Juego no encontrado",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Apuesta": {
            "type": "object",
            "properties": {
                "fecha": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "juego_id": {
                    "type": "integer"
                },
                "monto": {
                    "type": "number"
                },
                "resultado": {
                    "type": "string"
                },
                "usuario_id": {
                    "type": "integer"
                }
            }
        },
        "main.Juego": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "tipo": {
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
