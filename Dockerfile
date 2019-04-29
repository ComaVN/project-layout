FROM golang:alpine

WORKDIR /go/src/project-layout

COPY cmd cmd
COPY internal internal
RUN ls -la /go/src
RUN go build -o myapp ./cmd/myapp/main.go
RUN ./myapp
