package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Привет Мир из сервера!"))
}

func main() {
	s := &server{
		addr: ":8080",
	}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(err)
	}
}
