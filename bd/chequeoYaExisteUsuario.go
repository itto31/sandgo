package bd

import (
	"context"
	"time"

	"github.com/itto31/sandgo/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario recibe un email de parametro y revisa si ya esta registado
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("sandgo")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}
