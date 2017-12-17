# install vendors
install:
	glide install

# build and start the application
run: build
	./bin/server

# build the application and put binary in bin directory
build:
	go build -o ./bin/server
