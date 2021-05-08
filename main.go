package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"github.com/imrushi/restapi/helper"
	"github.com/imrushi/restapi/models"
)

//Connection mongoDB with helper class
var collection = helper.ConnectDB()

func createForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")

	var form models.Form

	//we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&form)

	//insert our form model
	result, err := collection.InsertOne(context.TODO(), form)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}



func main() {
	//Init Router
	r := mux.NewRouter()
	//arrange our route
	r.HandleFunc("/api/form", createForm).Methods("POST")


	//cors enable
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:3000"},
        AllowCredentials: true,
    })

	handler := c.Handler(r)

	//set out port address
	log.Fatal(http.ListenAndServe(":8000", handler))
	fmt.Print("Server is running on port 8000")
}
