## Running
Make sure you have a docker daemon running and go installed.
Then, in your terminal run the following in the project's root directory:
1. `docker-compose up` to run the RabbitMQ server and the Go consumer.
2. `go run ./cmd/producer/main.go` to run the Go producer which will send 100 messages to the RabbitMQ server.

The `Consumer` will consume the messages and print them to the terminal.
The time to process each message individually is 1 second.
Thanks to concurrency implemented on `internal/mqlistener/listener.go`, the `Consumer` will process 5 messages at a time.