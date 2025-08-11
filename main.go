package main

import "challenge/cmd"

var (
	smsWorkerPoolSize   = int8(4)
	// depending on your RAM you can increase this
	smsWorkerBufferSize = int64(1024)
	serverPort          = ":8080"
)

func main() {
	cmd.Serve(serverPort, smsWorkerBufferSize, smsWorkerPoolSize)
}
