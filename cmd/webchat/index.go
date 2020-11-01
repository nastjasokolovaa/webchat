package main

import (
	"fmt"
	"net/http"
)

func (s *server) index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello, world")
}
