package models

//Tweet captura el body, el mensaje que nos llego
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
