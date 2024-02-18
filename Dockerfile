FROM golang:1.22.0-bookworm AS build-stage
RUN mkdir -p /build/static/
WORKDIR /build
COPY static/ /build/static
COPY go.mod go.sum main.go /build/
RUN go build -o main
CMD [ "./main" ]
