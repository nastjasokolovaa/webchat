package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	addr := ":80"
	srv := &http.Server{
		Addr: addr,
		Handler: &server{
			mu:           &sync.RWMutex{},
			messageStore: make(map[dialogID][]message, 128),
		},
	}

	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		log.Printf("server is starting at %s", addr)
		if err := srv.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	log.Println("server stopped with signal", <-stopChan)
}

type dialogID = int64

type server struct {
	mu            *sync.RWMutex
	messageStore  map[dialogID][]message
	nextMessageID int64
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		s.index(w, req)
	case "/send":
		s.send(w, req)
	case "/messages":
		s.messages(w, req)
	case "/dialogs":
	default:
		http.NotFound(w, req)
	}
}
