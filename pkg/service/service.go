package service

import (
	"log"

	"challenge/internal/background"
	"challenge/pkg/sms"
)

type SmsService struct {
	queue *background.Queue[sms.SMS]
}

// Sends SMS in the background
func (s *SmsService) SendSMS(msg sms.SMS) {
	s.queue.Enq(msg)
	log.Printf("SMS enqueued %v\n", msg.Message)
}

// Worker that sends SMS
func (s *SmsService) SmsSenderWorker() {
	for {
		msg, ok := s.queue.Deq()
		if !ok {
			log.Print("Not able to get message from the queue")
		}
		go msg.Send()
	}
}

func (s *SmsService) StartWorkerPool(size int8) {
	for range make([]int, size) {
		go s.SmsSenderWorker()
	}
}
