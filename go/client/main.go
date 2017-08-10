package main

import (
	"log"
	"flag"
	//"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/gusga/grpc-5rabbits-talks/imagexample"

)
const (
	address = "localhost:50051"
)

func main() {
	var imageURL = flag.String("i", "", "Image url to resolve")
	flag.Parse()

	if len(*imageURL) == 0 {
		log.Fatal("Must enter a image url")
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	grpcClient := pb.NewImageCaptchaServiceClient(conn)
	response, err := grpcClient.ResolveCaptcha(context.Background(), &pb.ImageRequest{Url: *imageURL})
	if err != nil {
		log.Fatalf("Can resolve the captcha %v\n", err)
	}
	log.Printf("Captcha resolved %s in %s", response.GetCaptcha(), response.GetLanguage())
}	