# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY *.* ./

RUN go mod download

RUN go build -o /go-grpc-time-svc

EXPOSE 50051

CMD [ "/go-grpc-time-svc" ]