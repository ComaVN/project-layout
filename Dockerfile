FROM golang:alpine

RUN apk update \
    && apk --no-cache upgrade \
    && apk --no-cache add \
        git \
;

RUN go get -d \
    github.com/golang/glog \
;

WORKDIR /go/src/project-layout

COPY cmd cmd
COPY internal internal
RUN ls -la /go/src
RUN go build -o myapp ./cmd/myapp/main.go
RUN ./myapp -logtostderr
