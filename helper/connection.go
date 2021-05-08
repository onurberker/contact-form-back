package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	//set client opetions
	clientOptions := options.Client().ApplyURI("mongodb+srv://onur:2ogBrRJWtubxZMVl@cluster0.ztlgv.mongodb.net/Cluster0?retryWrites=true&w=majority")

	//connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Connected to MongoDB!")

	collection := client.Database("go_rest_api").Collection("forms")

	return collection
}

//ErrorResponse : This is error model
type ErrorResponse struct {
	StatusCode  int    `json:"status"`
	ErrorMesage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMesage: err.Error(),
		StatusCode:  http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
