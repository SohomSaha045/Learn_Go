package controller

import (
	// Import godotenv
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SohomSaha045/mongoAPI/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "Netflix"
const colName = "watchlist"

var collection *mongo.Collection

//connect with mongoDB

func init() {

	//client
	URL := goDotEnvVariable("URL")
	clientOption := options.Client().ApplyURI(URL)
	//connect
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection Succesfull")
	collection = client.Database(dbName).Collection(colName)
	//collection reference is ready
	fmt.Println("Collection reference is ready.")
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

//MONGODB helpers -file

// insert a record
func insertOne(movie model.Netflix) {
	res, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data Inserted:", res.InsertedID)
}

//update a record

func updateOne(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated count", res.ModifiedCount)
}

//delete a record

func deleteOne(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Delete count", res.DeletedCount)
}

func deleteAll() {
	res, err := collection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All Deleted count:", res.DeletedCount)
}

// get All
func getAll() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	//cursor is special type
	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cursor.Close(context.Background())
	return movies
}

//get One

// func getOne(movieID string) {
// 	var movie model.Netflix
// 	id, _ := primitive.ObjectIDFromHex(movieID)
// 	filter := bson.M{"_id": id}
// 	err := collection.FindOne(context.Background(), filter).Decode(&movie)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Movie Found is", movie)

// }

// Actual Controllers - file
func Get_All_Movies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAll()
	json.NewEncoder(w).Encode(allMovies)
}

func Create_a_Movie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOne(movie)
	json.NewEncoder(w).Encode(movie)

}
func Mark_As_Watched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func Delete_A_Movie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	params := mux.Vars(r)
	deleteOne(params["id"])
	var res string = "Items Deleted with id:"+params["id"]
	json.NewEncoder(w).Encode(res)
}
func Delete_All_Movies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAll()
	
	json.NewEncoder(w).Encode("All movies Deleted.")
}