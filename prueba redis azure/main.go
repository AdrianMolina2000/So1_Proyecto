package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

const keyPrefix = "user:"

type userHandler struct {
	client *redis.Client
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
	router.HandleFunc("/users/", uh.createUser).Methods(http.MethodPost)
	router.HandleFunc("/users/{userid}", uh.getUser).Methods(http.MethodGet)

	server := http.Server{Addr: ":8080", Handler: router}
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

func (uh userHandler) createUser(rw http.ResponseWriter, r *http.Request) {
	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	var u map[string]interface{}
	err = json.Unmarshal([]byte(payload), &u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	resultString := string(payload)

	//userid := u["id"].(string)
	_, err = uh.client.LPush(r.Context(), "prueba2", resultString).Result()
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (uh userHandler) getUser(rw http.ResponseWriter, r *http.Request) {
	//userid := mux.Vars(r)["userid"]
	info, err := uh.client.LRange(r.Context(), "prueba2", 0, 10000).Result()
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
