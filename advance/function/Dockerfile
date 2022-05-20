# Use the Python base image
FROM golang:1.17-buster as builder

# Set a working directory
WORKDIR /func

# Copy all the files from the local directory into the container
COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

RUN go build -v -o server

RUN export FUNCTION_TARGET=HelloWorld

EXPOSE 8080

# Install the Functions Framework
RUN go run cmd/main.go
#CMD ["/cmd/main"]