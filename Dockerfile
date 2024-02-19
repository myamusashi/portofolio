FROM golang:1.22-bullseye as builder

RUN mkdir -p /build/
WORKDIR /build

COPY src/go.* /build/
RUN go mod download

COPY src/main.go /build/
COPY src/static/ /build/static/
RUN go build -o main

CMD [ "./main" ]
