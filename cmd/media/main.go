package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"os/signal"
	"syscall"

	_ "v-shi/conf"

	"v-shi/cmd/media/server"

	_ "v-shi/pkg/miio"

	mediapb "v-shi/pkg/pb/media/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	rpcPort := os.Getenv("APP_MEDIA_RPC_PORT")
	gs := grpc.NewServer()

	// create instance of the GRPC Server
	media := server.NewMediaServer()

	// register instance of the media
	mediapb.RegisterMediaServer(gs, media)

	//register the reflection service which allow clients to check the service
	// use postman or evans
	reflection.Register(gs)

	// create a TCP socket for inbound connection
	addr := fmt.Sprintf(":%s", rpcPort)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	//listen for request
	go func() {
		log.Println("gRPC server is listening on : ", rpcPort)
		if err := gs.Serve(l); err != nil {
			log.Println("gRPC server failed to initialized on port : ", rpcPort)
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-c

	gs.Stop()
}
