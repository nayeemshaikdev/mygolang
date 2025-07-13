package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

const connectionString = "mongodb+srv://nayeemshaikdev:Aydin@050722@aydin.nkr61vq.mongodb.net/?retryWrites=true&w=majority&appName=Aydin"

var dbName = "netflix"
var colName = "watchlist"

//Most Important
var collection *mongo.collection

func init()  {
	clientOption := options.client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	
	if err != nil {
		genericError(err)
	}

	fmt.Println("Mongo DB connection success")

	collection := client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready : ",collection)


}

func genericError(err error){
	log.Fatal(err)
}