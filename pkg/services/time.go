package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hritikkhetan/go-grpc-time-svc/pkg/pb"
	"github.com/streadway/amqp"
)

type Server struct {
	pb.TimeServiceServer
}

func throwError(err error) {
	fmt.Println(err)
	panic(err)
}

func (s *Server) CurrTime(ctx context.Context, req *pb.CurrTimeRequest) (*pb.CurrTimeResponse, error) {

	amqpConnection()

	return &pb.CurrTimeResponse{
		Status: http.StatusCreated,
	}, nil
}

func amqpConnection() {

	fmt.Println("Setting RabbitMQ server")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		throwError(err)
	}
	defer conn.Close()
	fmt.Println("Successfully connected to our RabbitMQ instance")

	ch, err := conn.Channel()
	if err != nil {
		throwError(err)
	}
	defer ch.Close()

	publishCurrTime(ch)
}

func publishCurrTime(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(
		"RabbitMQ Queue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		throwError(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"RabbitMQ Queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(time.Now().Format(time.RFC3339)),
		},
	)

	if err != nil {
		throwError(err)
	}

	fmt.Println("Successfully published message to queue")
}
