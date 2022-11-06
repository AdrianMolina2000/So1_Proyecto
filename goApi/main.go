package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const keyPrefix = "user:"

var clientM *mongo.Client

type userHandler struct {
	client *redis.Client
}

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

func main() {

	//CONEXION REDIS
	redisHost := "redisPF.redis.cache.windows.net:6380"
	redisPassword := "fnU7xijFdWUwB0Cms1RlzqtuAI9D6ygGbAzCaKpFSlo="

	op := &redis.Options{Addr: redisHost, Password: redisPassword, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12}, WriteTimeout: 5 * time.Second}
	client := redis.NewClient(op)

	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("failed to connect with redis instance at %s - %v", redisHost, err)
	}

	uh := userHandler{client: client}

	router := mux.NewRouter()
	enableCORS(router)
	router.HandleFunc("/getPartidosRedis/", uh.getPartidos).Methods(http.MethodGet)

	//CONEXION MONGODB
	clientM, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://sopes-pf:vZfjeNrh74ozmNuaeZDf5Mqg3TWIfgs8bMM6Y10nkPTJqLDMj1rujRF3vaOiOHn3iXEQ2SzKnq8D0FxapGjqiQ%3D%3D@sopes-pf.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&maxIdleTimeMS=120000&appName=@sopes-pf@"))
	router.HandleFunc("/getPartidosMongo/", getResult).Methods("GET")

	//SERVIDOR
	server := http.Server{Addr: ":8000", Handler: router}
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("press ctrl+c to shutdown")
		<-exit
		if client != nil {
			err := client.Close()
			if err != nil {
				log.Println("failed to close redis", err)
			}
		}
		server.Shutdown(context.Background())
	}()

	log.Fatal(server.ListenAndServe())
	log.Println("application stopped")
}

func (uh userHandler) getPartidos(rw http.ResponseWriter, r *http.Request) {
	//userid := mux.Vars(r)["userid"]
	info, err := uh.client.LRange(r.Context(), "partidos", 0, 1000).Result()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(info) == 0 {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	rw.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(rw).Encode(info)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		rw.Header().Del("Content-Type")
	}
}

func getResult(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var listUsr []User
	var r Resultados
	colletion := clientM.Database("goDB").Collection("Result")
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
