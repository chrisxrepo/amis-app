package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

func main() {
	server := NewServer()
	server.Run()
}

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	router := http.NewServeMux()
	s := &Server{router: router}
	s.initRoute()

	return s
}

func (s *Server) Run() {
	l, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}

	fmt.Println("listen http://localhost:9999")
	if err := http.Serve(l, &httpHandle{router: s.router}); err != nil {
		panic(err)
	}
}

type httpHandle struct {
	router *http.ServeMux
}

func (h httpHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api/v1") {
		h.router.ServeHTTP(w, r)
		return
	}

	http.ServeFile(w, r, "ui"+r.URL.Path)

}

func (s *Server) initRoute() {
	// TODO
}
