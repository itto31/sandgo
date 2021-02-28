package bd

import (
	"context"
	"time"

	"github.com/itto31/sandgo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ModificoRegistro permite modificar el perfill del usuario
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("sandgo")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	if len(u.SitioWeb) > 0 {
		registro["sitioweb"] = u.SitioWeb
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}
