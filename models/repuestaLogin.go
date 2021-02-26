package models

//RespuestaLogin tiene el token que se devuelve del login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
