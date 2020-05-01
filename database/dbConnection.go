package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN : content a mongoDB conection*/
var MongoCN = ConectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://root:-solrac97G@golang-course-olaby.mongodb.net/test?retryWrites=true&w=majority")

/*ConectDB : Create a conection to mongoDB and return the conection*/
func ConectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Conect successfull")
	return client
}

/*ConnectionOK : Check the conection and return true or false */
func ConnectionOK() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
