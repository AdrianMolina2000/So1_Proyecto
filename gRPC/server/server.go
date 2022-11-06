package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	pb "servergRPC/proto"
	"time"

	"github.com/go-redis/redis/v8"

	// "github.com/gomodule/redigo/redis"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedAddDataServer
}

type Partido struct {
	Team1 string `json:"team1"`
	Team2 string `json:"team2"`
	Score string `json:"score"`
	Phase string `json:"phase"`
}

// SayHello implements helloworld.GreeterServer
func (s *server) AgregarData(ctx context.Context, in *pb.RequestData) (*pb.ResponseData, error) {
	team1 := in.GetTeam1()
	team2 := in.GetTeam2()
	score := in.GetScore()
	phase := in.GetPhase()

	output, err := json.Marshal(Partido{Team1: team1, Team2: team2, Score: score, Phase: phase})
	if err != nil {
		fmt.Println(err)
	}

	redisHost := "redisPF.redis.cache.windows.net:6380"
	redisPassword := "fnU7xijFdWUwB0Cms1RlzqtuAI9D6ygGbAzCaKpFSlo="

	op := &redis.Options{Addr: redisHost, Password: redisPassword, TLSConfig: &tls.Config{MinVersion: tls.VersionTLS12}, WriteTimeout: 5 * time.Second}
	client := redis.NewClient(op)

	err = client.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("failed to connect with redis instance at %s - %v", redisHost, err)
	}

	_, err = client.LPush(ctx, "partidos", output).Result()
	if err != nil {
		fmt.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
	}

	// conn, err := redis.DialURL("redis://localhost:6379")
	// if err != nil {
	// 	fmt.Printf("ERROR: fail initializing the redis pool: %s", err.Error())
	// }

	// a, err := conn.Do("lpush", "partidos", output)
	// if err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println(a)
	// }

	return &pb.ResponseData{Respuesta: "Se ha insertado conrrectamente el partido"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddDataServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
