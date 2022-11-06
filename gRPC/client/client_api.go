package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	pb "clientgRPC/proto"

	"google.golang.org/grpc"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API GO - gRPC Client!\n"))
}

func addPartidos(w http.ResponseWriter, r *http.Request) {
	team1 := mux.Vars(r)["team1"]
	team2 := mux.Vars(r)["team2"]
	score := mux.Vars(r)["score"]
	phase := mux.Vars(r)["phase"]

	/********************************** gRPC llamada al servidor ********************************/
	conn, err := grpc.Dial("20.121.182.112:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddDataClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := c.AgregarData(ctx, &pb.RequestData{
		Team1: team1,
		Team2: team2,
		Score: score,
		Phase: phase,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	/********************************** gRPC ********************************/

	/********************************** Respuesta ********************************/
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(struct {
		Mensaje string `json:"mensaje"`
	}{Mensaje: reply.GetRespuesta()})
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/", IndexHandler)
	router.HandleFunc("/team1/{team1}/team2/{team2}/score/{score}/phase/{phase}", addPartidos).Methods("POST")
	log.Println("Listening at port 2000")
	log.Fatal(http.ListenAndServe(":2000", router))
}
