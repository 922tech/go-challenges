package service

import (
	"testing"
	"time"

	"challenge/internal/background"
	"challenge/pkg/sms"
)

func TestSendSMS(t *testing.T) {
	q := background.NewQueue[sms.SMS](100)
	s := SmsService{queue: q}
	go s.StartWorkerPool(3)

	for _, i := range []string{"a", "b", "c", "d", "e", "f"} {
		s.SendSMS(sms.SMS{Message: i, Reciever: i})
	}
	time.Sleep(4 * time.Second)
	if q.Len() != 0 {
		t.Fail()
		t.Log("queue is not empty yet")
	}
}
