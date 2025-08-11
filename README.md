# Background SMS Sender

A high-performance SMS sending service implemented in Go using buffered channels for background processing.

## Features

- Asynchronous SMS processing using worker pools
- Buffered queue system for message handling
- HTTP API endpoint for SMS submission
- Configurable worker pool and buffer sizes
- Built-in retry mechanism for failed SMS deliveries

## Architecture

The service consists of several key components:

- HTTP Server: Handles incoming SMS requests ([cmd/cmd.go](cmd/cmd.go))
- SMS: Handles the actual SMS sending logic ([pkg/sms/sms.go](pkg/sms/sms.go))
- SMS Service: Manages the queuing and processing of messages ([pkg/service/service.go](pkg/service/service.go))
- Background Queue: Generic queue implementation using channels ([internal/background/queue.go](internal/background/queue.go))

## Configuration

The following parameters can be configured in [main.go](main.go):

```go
smsWorkerPoolSize   = int8(4)      // Number of concurrent workers
smsWorkerBufferSize = int64(1024)  // Size of the message buffer
serverPort          = ":8080"      // HTTP server port
```

## API Usage

Send an SMS using a POST request:

```bash
curl -X POST http://localhost:8080 \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello", "reciever": "+1234567890"}'
```

## Running the Project

1. Clone the repository
2. Install dependencies:

```bash
go mod download
```

3. Start the server:

```bash
go run main.go
```

## Testing

Run the test suite:

```bash
go test ./...
```

## Performance Benchmark

Benchmark results with 1000 requests, 20 concurrent connections:

- Requests/sec: 16,194.85
- Average latency: 1.1ms
- 95th percentile: 3.3ms

Full benchmark details:

- 99% of requests complete within 7.7ms
- DNS+dialup average: 0.1ms
- Response wait time average: 1.0ms

### Details

```bash
$ hey -n 1000 -c 20 -m POST -d '{"message":"va
lue"}' -H "Content-Type: application/json" http:/
/localhost:8080

Summary:
  Total:        0.0617 secs
  Slowest:      0.0097 secs
  Fastest:      0.0001 secs
  Average:      0.0011 secs
  Requests/sec: 16194.8500


Response time histogram:
  0.000 [1]     |
  0.001 [622]   |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.002 [234]   |■■■■■■■■■■■■■■■
  0.003 [84]    |■■■■■
  0.004 [21]    |■
  0.005 [7]     |
  0.006 [8]     |■
  0.007 [6]     |
  0.008 [8]     |■
  0.009 [7]     |
  0.010 [2]     |


Latency distribution:
  10% in 0.0002 secs
  25% in 0.0003 secs
  50% in 0.0007 secs
  75% in 0.0015 secs
  90% in 0.0024 secs
  95% in 0.0033 secs
  99% in 0.0077 secs

Details (average, fastest, slowest):
  DNS+dialup:   0.0001 secs, 0.0001 secs, 0.0097 secs
  DNS-lookup:   0.0000 secs, 0.0000 secs, 0.0024 secs
  req write:    0.0000 secs, 0.0000 secs, 0.0016 secs
  resp wait:    0.0010 secs, 0.0001 secs, 0.0065 secs
  resp read:    0.0001 secs, 0.0000 secs, 0.0015 secs

Status code distribution:
  [200] 1000 responses

```
