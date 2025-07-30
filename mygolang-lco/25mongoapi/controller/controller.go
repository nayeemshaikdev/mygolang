package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/nayeemshaikdev/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	//"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://nayeemshaikdev:Aydin@050722@aydin.nkr61vq.mongodb.net/?retryWrites=true&w=majority&appName=Aydin"

var dbName = "netflix"
var colName = "watchlist"

//Most Important
var collection *mongo.Collection

func init()  {
	//clientOption := options.Client().ApplyURI(connectionString)

	// Create a context with a timeout
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	//client, err := mongo.Connect(context.TODO(), clientOption)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	
	if err != nil {
		genericError(err)
	}

	fmt.Println("Mongo DB connection success")

	collection := client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready : ",collection)


}

// Mongodb controllers
func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie)
	genericError(err)

	fmt.Println("Inserted one movie in db with Id is : ", inserted.InsertedID)

}

func updateOneMovie(movieId string)  {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	genericError(err)

	fmt.Println("modified count : ",result.ModifiedCount)


}

func deleteOneMovie(movieId string)  {
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	genericError(err)

	fmt.Println("Movie got deleted with delete count : ",deleteCount)	
}

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.M{{}})

	genericError(err)

	fmt.Println("Number of movies delete : ",deleteResult.DeletedCount)
	return deleteResult.DeletedCount
	
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.M{{}})
	genericError(err)

	var movies []primitive.M

	for cur.Next(context.Background()){
		var movie bson.M

		err := cur.Decode(&movie)
		genericError(err)
		movies = append(movies, movie)	
	}
	defer cur.Close(context.Background())
	return movies
}

func GetAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()

	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix

	err := json.NewDecoder(r.Body).Decode(&movie)
	genericError(err)

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}

func genericError(err error){
	if err != nil{
		log.Fatal(err)
	}	
}