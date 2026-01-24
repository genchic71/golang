package main

import "fmt"

// zad12
type Server struct {
	Port int
	Host string
	Tls  bool
}

type Serveroption func(*Server)

func WithPort(port int) Serveroption {
	return func(server *Server) {
		server.Port = port
	}
}
func WithHost(host string) Serveroption {
	return func(server *Server) {
		server.Host = host
	}
}
func WihtTls(tls bool) Serveroption {
	return func(server *Server) {
		server.Tls = tls
	}
}

func Newserver(opt ...Serveroption) *Server {
	s := &Server{
		Port: 8080,
		Host: "localhost",
		Tls:  false,
	}
	for _, o := range opt {
		o(s)
	}
	return s
}

func main() {
	s := Newserver(WithPort(7000))
	fmt.Println(s)
}
