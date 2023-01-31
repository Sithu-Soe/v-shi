package server

import (
	"io"
	"log"
	"net/http"
	"os"
	mediapb "v-shi/pkg/pb/media/v1"
	"v-shi/pkg/utils"
)

type MediaServer struct {
	mediapb.UnimplementedMediaServer
}

func NewMediaServer() *MediaServer {
	return &MediaServer{}
}

func (m *MediaServer) CreateMedia(stream mediapb.Media_CreateMediaServer) error {
	var filename string
	var fullPath string

	var fo *os.File
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	for {
		req, err := stream.Recv()

		// At the start of streaming, create file
		if len(filename) == 0 {
			// open saved file
			filename = utils.GenerateUniqueCode(req.GetPrefix()) + req.GetExt()
			fullPath = req.GetFilePath() + filename

			log.Println(http.DetectContentType(req.GetData()))

			fo, err = os.Create(filename)
			if err != nil {
				return err
			}
		}

		// read by chunk
		n := len(req.GetData())

		if err == io.EOF || len(req.GetData()) == 0 {
			// we've reached end of request receiving.
			// we have finished reading the client stream
			res := &mediapb.MediaCreateResponse{
				Filename: filename,
			}

			return stream.SendAndClose(res)
		}

		if err != nil {
			log.Printf("Error while receiving requests: %v", err)
			return err
		}

		log.Println("HERE", fullPath)

		// write a chunk to a file
		if _, err := fo.Write(req.GetData()[:n]); err != nil {
			return err
		}
	}

}
