package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/govindyagyasaini/mongodbapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "your-db-name"
const colName = "your-collection-name"
const connectionURI = "your-connection-uri"

var collection *mongo.Collection

// Connect to database
func init() {

	//create client options instance
	//Client creates a new client options instance
	//It is a constructor in mongo package
	//This client is not yet connected to the database; it's just the settings for how the connection will be made.
	//ApplyURI method configures the client to use the MongoDB URI
	clientOption := options.Client().ApplyURI(connectionURI)

	//Connect method establishes the mongoDB connection using the client configuration object created earlier
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal("Error while connecting to MongoDB:", err)
	}

	//Check if Mongo DB connection is successful by doing a ping request
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	} else {
		fmt.Println("Connected to MongoDB successfully.")
	}

	//Access/create collection inside database in mongoDB
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection reference is ready.")

}

func insertMovie(movie model.PrimeVideo) {
	fmt.Println("Inserting movie into mongodb collection")
	result, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal("Insertion of document failed in mongodb", err)
	}
	fmt.Printf("Inserted document with _id %v\n", result.InsertedID)

}

func updateMovie(movieID string) {
	fmt.Println("Updating movie as watched")

	// Convert movieID to ObjectID, handle error if it fails
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal("Invalid movieID:", err)
		return
	}
	fmt.Println("Movie ID:", movieID)

	filter := bson.M{"_id": id} //
	update := bson.M{"$set": bson.M{"watched": true}}
	fmt.Println(id)

	result, err := collection.UpdateOne(context.Background(), filter, update) //
	if err != nil {
		log.Fatal("Error occurred while updating the movie.", err)
	}

	// Check if the movie was actually updated
	if result.ModifiedCount == 0 {
		fmt.Println("No movie found with that ID to update.")
	} else {
		fmt.Printf("Modified count is %v\n", result.ModifiedCount)
	}
}

func deleteMovie(movieID string) {
	fmt.Println("Deleting movie from mongodb")

	// Convert movieID to ObjectID, handle error if it fails
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal("Invalid movieID:", err)
		return
	}

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal("Error while deleting movie:", err)
	}

	// Check if the movie was actually deleted
	if result.DeletedCount == 0 {
		fmt.Println("No movie found with that ID to delete.")
	} else {
		fmt.Println("Deleted count is:", result.DeletedCount)
	}
}

func deleteAllMovies() {
	fmt.Println("Deleting all movies from mongodb")
	result, err := collection.DeleteMany(context.TODO(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal("Error while deleting movie", err)
	}
	fmt.Println("Deleted all movies and delete count is: ", result.DeletedCount)

}

func getAllMovies() []primitive.M {
	fmt.Println("Getting all movies from mongodb")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Error occured while getting the movies.", err)
	}
	defer cursor.Close(context.TODO())
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, primitive.M(movie))
	}
	return movies
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("<h1>Welcome to MongoDB API in GoLang</h1>")

}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "GET")
	result := getAllMovies()
	json.NewEncoder(w).Encode(result)

}

func InsertMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.PrimeVideo
	json.NewDecoder(r.Body).Decode(&movie)
	insertMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Marking a movie as watched")
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteAllMovies()
	json.NewEncoder(w).Encode("All movies deleted")

}
