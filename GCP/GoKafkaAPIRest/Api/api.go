package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

type Info struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase int    `json:"phase"`
}

func Input(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	info := Info{}

	json.Unmarshal(body, &info)

	message, _ := json.Marshal(&info)
	conn, err := kafka.DialLeader(context.Background(), "tcp", "kafka-cluster-kafka-bootstrap", "my-topic", 0)

	if err != nil {
		fmt.Fprint(w, "1. ")
		fmt.Fprintln(w, err)
		return
	}

	if conn.SetWriteDeadline(time.Now().Add(time.Second*10)) != nil {
		fmt.Fprint(w, "2. ")
		fmt.Fprintln(w, err)
		return
	}
	_, err = conn.WriteMessages(kafka.Message{Value: message})
	if err != nil {
		fmt.Fprint(w, "3. ")
		fmt.Fprintln(w, err)
		return
	}

	fmt.Fprintln(w, string(message))
}

func NewApi() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/input", Input).Methods("POST")
	log.Println("RestAPI up")
	log.Fatal(http.ListenAndServe(":3030", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(r)))
}
