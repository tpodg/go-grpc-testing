package main

import (
	"log"
	"net/http"
)

type handler struct {
	router *http.ServeMux
}

func newHandler() handler {
	h := handler{
		router: http.NewServeMux(),
	}

	h.router.HandleFunc("/rest", h.sendRest)

	return h
}

func (h handler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	h.router.ServeHTTP(rw, req)
}

func (h handler) sendRest(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving %s /rest", r.Method)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

func serveRest(s *http.Server) {
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("failed to listen and serve", err)
	}
}
