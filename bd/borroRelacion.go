package bd

import (
	"context"
	"time"

	"github.com/itto31/sandgo/models"
)

//BorroRelacion borra la relacion en la BD
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoC.Database("sandgo")
	col := bd.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil

}
