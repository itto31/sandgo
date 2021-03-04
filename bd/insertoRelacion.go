package bd

import (
	"context"
	"time"

	"github.com/itto31/sandgo/models"
)

//InsertarRelacion graba la relacion en la BD
func InsertarRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoC.Database("sandgo")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
