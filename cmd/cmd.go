package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"challenge/pkg/service"
	"challenge/pkg/sms"
)

func Serve(port string, smsBufferSize int64, workers int8) {
	s := service.NewSmsService(smsBufferSize)
	go s.StartWorkerPool(workers)
	http.HandleFunc("/", makeSMSHandler(&s))
	fmt.Printf("Server is running on %s\n", port)
	http.ListenAndServe(port, nil)
}

func makeSMSHandler(s *service.SmsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		data, err := parseSMSRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.SendSMS(*data)
	}
}

func parseSMSRequest(r *http.Request) (*sms.SMS, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body")
	}
	defer r.Body.Close()

	var data sms.SMS
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("invalid JSON")
	}
	return &data, nil
}
