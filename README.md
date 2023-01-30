# go-grpc-time-svc
	Clone or download the zip folder from green Code dropdown

# Command to run the RabbitMQ server
	docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
  
# Command to run the service
	Open the terminal
	Go to the go-grpc-time-svc directory
	Run the below command
		go run consumer.go
		
# To call the CurrTime Api
	Hit the below curl command
		curl --location --request GET 'http://localhost:3000/time/currTime'
