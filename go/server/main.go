package main

import (
	"strings"
	"log"
	"image"
	"net"
	"net/http"
	_ "image/png"
	
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/gusga/grpc-5rabbits-talks/imagexample"
	"github.com/otiai10/gosseract"
)

const (
	port = ":50051"
)


type server struct{}

func (s *server) ResolveCaptcha(ctx context.Context, in *pb.ImageRequest) (*pb.ImageResponse, error) {
	client, err := gosseract.NewClient()
	pngImage := getImage(in.GetUrl())
	log.Println("Resolving captcha")
	captcha, err := client.Image(pngImage).Out()
	captcha = strings.TrimSpace(captcha)
	if err != nil {
		log.Fatalf("Error resolving captcha %v\n", err)
	}
	log.Printf("Captcha resolve: %s\n", captcha)
	return &pb.ImageResponse{Captcha: captcha, Language: "Go"}, nil
}

func getImage(url string) image.Image {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error downloading the image %v\n", err)
	}
	log.Println("Getting image")
	defer resp.Body.Close()
	imageToResolve, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Fatalf("Error decoding image %v\n", err)
	}
	log.Println("Decoding image")
	return imageToResolve
}

func main() {
	
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}
	log.Printf("Listening on port %s\n", port)
	s := grpc.NewServer()
	pb.RegisterImageCaptchaServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}

}