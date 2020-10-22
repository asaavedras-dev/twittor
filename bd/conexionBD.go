package bd

import (
	"context"
	"log"

	"go.mongo.org/mongo-driver/mongo"
	"go.mongo.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la BDD*/
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:<password>@twittor.kphff.mongodb.net/<dbname>?retryWrites=true&w=majority")

/*ConectarBD es la funcion que me permite conectar la BDD*/
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión Exitosa con la BDD")
}

/*ChequeoConnection es el Ping a la BDD*/
func ChequeoConnection() int {
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
