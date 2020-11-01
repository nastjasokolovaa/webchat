package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type message struct {
	ID       int64     `json:"id"`
	UserName string    `json:"user_name"`
	Message  string    `json:"message"`
	Time     time.Time `json:"time"`
}

type dialog struct {
	ID            int64 `json:"id"`
	OwnerID       int64 `json:"owner_id"`
	ParticipantID int64 `json:"participant_id"`
}

func (s *server) messages(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	dialogID, err := strconv.ParseInt(req.URL.Query().Get("dialog_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}

	s.mu.RLock()
	encoder.Encode(s.messageStore[dialogID])
	s.mu.RUnlock()
}

func (s *server) send(w http.ResponseWriter, req *http.Request) {
	encoder := json.NewEncoder(w)
	dialogID, err := strconv.ParseInt(req.URL.Query().Get("dialog_id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err.Error())
		return
	}
	nextMessage := req.URL.Query().Get("message")
	if len(nextMessage) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode("message is empty")
		return
	}

	s.mu.Lock()
	s.nextMessageID++
	s.messageStore[dialogID] = append(s.messageStore[dialogID], message{
		ID:       s.nextMessageID,
		UserName: "www",
		Message:  nextMessage,
		Time:     time.Now(),
	})
	s.mu.Unlock()

	s.mu.RLock()
	encoder.Encode(s.messageStore[dialogID])
	s.mu.RUnlock()
}
