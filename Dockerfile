FROM golang:1.11.5

# Move current project to a valid go path
COPY . /go/src/github.com/mvc-golang-final/
WORKDIR /go/src/github.com/mvc-golang-final/

# Install project dependencies
RUN go mod init
RUN go mod vendor

# Run app in production mode
EXPOSE 7070
CMD ["go", "run", "main.go"]
