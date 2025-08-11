package sms

import (
	"fmt"
	"log"
	"time"
)

type SMS struct {
	Message  string `json:"message"`
	Reciever string `json:"reciever"`
}

var counter int

func (s *SMS) Send() (bool,error) {
	log.Print("sending sms...")
	counter+=1
	fmt.Printf("counter %v\n", counter)
	if counter%4== 0 {
		time.Sleep(3 * time.Second)
		return false, fmt.Errorf("error in sending message to %s", s.Reciever)
	}
	log.Printf("SMS sent to %s successfully", s.Reciever)

	return true, nil
}
