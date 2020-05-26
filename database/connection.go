package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN : content a mongoDB connection*/
var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://root:-solrac97G@golang-course-olaby.mongodb.net/test?retryWrites=true&w=majority")

/*ConnectDB : Create a connection to mongoDB and return the connection*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("Connect successfully")
	return client
}

/*ConnectionOK : Check the connection and return true or false */
func ConnectionOK() bool {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
