package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/davidlux123/gRPC-service/src/producers"
	pb "github.com/davidlux123/gRPC-service/src/proto"
	"google.golang.org/grpc"
	
)

type server struct {
	pb.UnimplementedIngressGameServer
}

func (s *server) SendResultGame(ctx context.Context, in *pb.GameRequest) (*pb.GameReply, error) {
	log.Printf("id: %v, players: %v", in.GetGameId(), in.GetPlayers())
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	resp := producers.SaveToKafka(in.GetGameId(), in.GetPlayers(), fecha)
	//resp := "{" + `"id":` + strconv.Itoa(int(in.GetGameId())) + "," + `"players":` + strconv.Itoa(int(in.GetPlayers())) + "," + `"date_game":` + fecha + "}"
	return &pb.GameReply{Response_Game: resp}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterIngressGameServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
