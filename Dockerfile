FROM golang:alpine

WORKDIR /go/src/myproject

COPY cmd cmd
COPY internal /go/src/
RUN ls -la /go/src
RUN go build -o myapp ./cmd/myapp/main.go
RUN ./myapp
