package model

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Responses []Response

type MessageResponse struct {
	Data     interface{} `json:"data"`
	Errors   Responses   `json:"errors"`
	Messages Responses   `json:"messages"`
}

/*
 Estructura de los mensajes de respuesta:
{
	"data" : {un objeto o array de objetos},
	"errors" : [
		{"code": "unexpected", "message": "algo pasó"},
		{"code": "not_found", "message": "no se encontró"}
	],
	"messages": [
		{"code": "ok", "message": "ok"},
		{"code": "record_created", "message": "registro creado"}
	]
}
*/
