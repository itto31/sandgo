package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DevuelvoTweetsSeguidores es la estructura con la devolveremos los tweets
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Usuario           string             `bson:"usuarioid" jsos:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuarioid" json:"userId,omitempty"`
	Tweet             struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
		ID      string
	}
}
