package service

import (
	"fmt"
	"log"
	"os"
	mediapb "v-shi/pkg/pb/media/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type MediaService struct {
	Client mediapb.MediaClient
}

func NewMediaServer() *MediaService {
	svc := &MediaService{}
	svc.connectMediaRPC()

	return svc
}

func (s *MediaService) connectMediaRPC() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	log.Println(os.Getenv("APP_MEDIA_RPC_PORT"), "KKK")
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%v", os.Getenv("APP_MEDIA_RPC_PORT")), opts...)
	if err != nil {
		panic(err)
	}

	s.Client = mediapb.NewMediaClient(conn)
}
