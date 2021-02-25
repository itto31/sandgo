package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoC es el objeto de conexion
var MongoC = conectarBD()
var clienteOptions = options.Client().ApplyURI("mongodb+srv://root:1234@sandgo.sp7l2.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//Conexion para la base de datos
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa")
	return client
}

//ChequeoConnection comprueba la conexion
func ChequeoConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1

}
