package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Team1 string `json:"team1,omitempty"`
	Team2 string `json:"team2,omitempty"`
	Score string `json:"score,omitempty"`
	Phase int    `json:"phase,omitempty"`
}

type Resultados struct {
	Results []User `json:"results,omitempty"`
	Record  int    `json:"record,omitempty"`
}

var client *mongo.Client

// CORS
func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
func createResult(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var usr User
	json.NewDecoder((request.Body)).Decode(&usr)
	colletion := client.Database("goDB").Collection("Result")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := colletion.InsertOne(ctx, usr)
	json.NewEncoder(response).Encode(result)
}

func getResult(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var listUsr []User
	var r Resultados
	colletion := client.Database("goDB").Collection("Result")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}}).SetLimit(10)
	cursor, err := colletion.Find(ctx, bson.D{}, opts)
	results, _ := colletion.CountDocuments(ctx, bson.D{})
	if err != nil {
		fmt.Println(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var usr User
		cursor.Decode(&usr)
		listUsr = append(listUsr, usr)
	}
	r.Results = listUsr
	r.Record = int(results)
	json.NewEncoder(response).Encode(r)
}
func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://sopes-pf:vZfjeNrh74ozmNuaeZDf5Mqg3TWIfgs8bMM6Y10nkPTJqLDMj1rujRF3vaOiOHn3iXEQ2SzKnq8D0FxapGjqiQ%3D%3D@sopes-pf.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&maxIdleTimeMS=120000&appName=@sopes-pf@"))
	router := mux.NewRouter()
	enableCORS(router)
	router.HandleFunc("/create", createResult).Methods("POST")
	router.HandleFunc("/get", getResult).Methods("GET")
	fmt.Println("Server on port", 8000)
	http.ListenAndServe(":8000", router)
	defer client.Disconnect(ctx)
}
