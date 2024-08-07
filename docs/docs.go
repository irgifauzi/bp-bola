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
            "url": "https://github.com/irgifauzi",
            "email": "714220035@std.ulbi.ac.id"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/delete/{id}": {
            "delete": {
                "description": "Hapus data pemain.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pemain"
                ],
                "summary": "Delete data pemain.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/insert": {
            "post": {
                "description": "Input data pemain.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pemain"
                ],
                "summary": "Insert data pemain.",
                "parameters": [
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqPemain"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Pemain"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/pemain": {
            "get": {
                "description": "Mengambil semua data pemain.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pemain"
                ],
                "summary": "Get All Data Pemain.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Pemain"
                        }
                    }
                }
            }
        },
        "/pemain/{id}": {
            "get": {
                "description": "Ambil per ID data pemain.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pemain"
                ],
                "summary": "Get By ID Data Pemain.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Pemain"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/update/{id}": {
            "put": {
                "description": "Ubah data pemain.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pemain"
                ],
                "summary": "Update data pemain.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Masukan ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ReqPemain"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.Pemain"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.Pemain": {
            "type": "object",
            "properties": {
                "_id": {
                    "type": "string",
                    "example": "12345"
                },
                "berat": {
                    "type": "number",
                    "example": 80
                },
                "nama_pemain": {
                    "type": "string",
                    "example": "haaland"
                },
                "negara": {
                    "type": "string",
                    "example": "Denmark"
                },
                "no_punggung": {
                    "type": "integer",
                    "example": 9
                },
                "posisi": {
                    "type": "string",
                    "example": "Striker"
                },
                "tanggal_lahir": {
                    "type": "string",
                    "example": "09-06-1985"
                },
                "tim": {
                    "$ref": "#/definitions/controller.ReqClub"
                },
                "tinggi": {
                    "type": "number",
                    "example": 192
                }
            }
        },
        "controller.ReqClub": {
            "type": "object",
            "properties": {
                "jumlah_pemain": {
                    "type": "integer",
                    "example": 11
                },
                "liga": {
                    "type": "string",
                    "example": "La liga"
                },
                "logo": {
                    "type": "string",
                    "example": "https://upload.wikimedia.org/wikipedia/id/4/43/Al-Nassr_fc.png"
                },
                "manajer": {
                    "type": "string",
                    "example": "Jajang"
                },
                "nama_club": {
                    "type": "string",
                    "example": "Real Madrid"
                },
                "stadion": {
                    "type": "string",
                    "example": "Si Jalak"
                },
                "tahun_berdiri": {
                    "type": "integer",
                    "example": 1985
                }
            }
        },
        "controller.ReqPemain": {
            "type": "object",
            "properties": {
                "berat": {
                    "type": "number",
                    "example": 80
                },
                "nama_pemain": {
                    "type": "string",
                    "example": "haaland"
                },
                "negara": {
                    "type": "string",
                    "example": "Denmark"
                },
                "no_punggung": {
                    "type": "integer",
                    "example": 9
                },
                "posisi": {
                    "type": "string",
                    "example": "Striker"
                },
                "tanggal_lahir": {
                    "type": "string",
                    "example": "09-06-1985"
                },
                "tim": {
                    "$ref": "#/definitions/controller.ReqClub"
                },
                "tinggi": {
                    "type": "number",
                    "example": 192
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "uas-bola-919c3c7dac41.herokuapp.com",
	BasePath:         "/",
	Schemes:          []string{"https", "http"},
	Title:            "TES SWAGGER ULBI",
	Description:      "This is a sample swagger for Fiber",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
