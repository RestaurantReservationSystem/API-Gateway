package main

import (
	"api_get_way/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn1, err := grpc.NewClient(":8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error")
	}
	conn2, err := grpc.NewClient(":8083", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error")
	}
	conn3, err := grpc.NewClient(":8089", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error")
	}
	router := api.RouterApi(conn1, conn2, conn3)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal("error is pai get way connection port")
	}

}
