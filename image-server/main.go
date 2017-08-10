package main

import (
	"log"
	"net/http"
	"os"
	"io/ioutil"
)
func port() string{
	p := os.Getenv("PORT")
	if len(p) == 0 {
		p = "9000"
	}
	return p
}

func listImages(root string) {
	files, err := ioutil.ReadDir(root + "/images")
	if err != nil {
		log.Fatal(err)
	}
    for _, f := range files {
		log.Printf("Image: %s\n", f.Name())
	}
}

func main() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(root)
	listImages(root)
	http.Handle("/", http.FileServer(http.Dir(root + "/images")))
	log.Printf("Serving images on port %s\n", port())
	log.Fatal(http.ListenAndServe(":" + port(), nil))
}
