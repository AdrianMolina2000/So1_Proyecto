package main

import (
	b "bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Info struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase int    `json:"phase"`
}

func consumer() {
	conn, _ := kafka.DialLeader(context.Background(), "tcp", "kafka-cluster-kafka-bootstrap", "my-topic", 0)
	conn.SetWriteDeadline(time.Now().Add(time.Second * 3))

	collection := client.Database("goDB").Collection("Result")
	ctx := context.TODO()

	for {
		msg, err := conn.ReadMessage(1e6)
		http.Post("http://3.144.197.243:3030/kafka", "application/json", b.NewBuffer([]byte("{\"HolaMundo\":\""+string(msg.Value)+"\"}")))
		if err != nil {
			http.Post("http://3.144.197.243:3030/kafka", "application/json", b.NewBuffer([]byte("{\"HolaMundo\":\""+err.Error()+"\"}")))
			conn, _ = kafka.DialLeader(context.Background(), "tcp", "kafka-cluster-kafka-bootstrap", "my-topic", 0)
			conn.SetWriteDeadline(time.Now().Add(time.Second * 3))
			continue
		}

		info := Info{}
		json.Unmarshal(msg.Value, &info)

		_, err = collection.InsertOne(ctx, info)
		if err != nil {
			http.Post("http://3.144.197.243:3030/kafka", "application/json", b.NewBuffer([]byte("{\"HolaMundo\":\""+err.Error()+"\"}")))
			client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://sopes-pf:vZfjeNrh74ozmNuaeZDf5Mqg3TWIfgs8bMM6Y10nkPTJqLDMj1rujRF3vaOiOHn3iXEQ2SzKnq8D0FxapGjqiQ%3D%3D@sopes-pf.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&maxIdleTimeMS=120000&appName=@sopes-pf@"))
			continue
		}

		_, err = http.PostForm("http://localhost:2000/team1/"+info.Team1+"/team2/"+info.Team2+"/score/"+info.Score+"/phase/"+strconv.Itoa(info.Phase), url.Values{})

		if err != nil {
			http.Post("http://3.144.197.243:3030/kafka", "application/json", b.NewBuffer([]byte("{\"HolaMundo\":\""+err.Error()+"\"}")))
			continue
		}

		http.Post("http://3.144.197.243:3030/kafka", "application/json", b.NewBuffer([]byte("{\"HolaMundo\":\"insertado correctamente\"}")))
	}
}

func main() {
	http.Post("http://3.144.197.243:3030/kafka", "application/json", b.NewBuffer([]byte("{\"Hola\":11}")))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://sopes-pf:vZfjeNrh74ozmNuaeZDf5Mqg3TWIfgs8bMM6Y10nkPTJqLDMj1rujRF3vaOiOHn3iXEQ2SzKnq8D0FxapGjqiQ%3D%3D@sopes-pf.mongo.cosmos.azure.com:10255/?ssl=true&retrywrites=false&maxIdleTimeMS=120000&appName=@sopes-pf@"))
	defer client.Disconnect(ctx)

	consumer()
}
