appname := server
image_name := gusga/5rabbits-talk-image-server

build:
	mkdir build/ && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/$(appname)

docker:
	docker build -t $(image_name) .

push:
	docker push $(image_name)

clean:
	\rm -rf build/

.PHONY: clean build docker
