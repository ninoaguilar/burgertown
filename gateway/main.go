package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	common "github.com/ninoaguilar/burgertown/common"
	pb "github.com/ninoaguilar/burgertown/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:2000"
)

func main() {
	conn, error := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if error != nil {
		log.Fatalf("Failed to dial server: %v", error)
	}
	defer conn.Close()

	log.Println("Dialing order service at ", orderServiceAddr)

	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http service")
	}
}
