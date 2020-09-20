package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
)

type ServerTls struct {
	*http.Server
}

var cfgTls = true

const port = ":8080"

func (s *ServerTls) addr() string {
	s.Addr = port
	return s.Addr
}
func (s *ServerTls) Start(c context.Context, h http.Handler) error {
	s.Handler = h

	fmt.Printf("Create a sample server, port %s , with tls : %v", s.addr(), s.TlsConfig())

	if cfgTls {
		fmt.Println("https")
		err := s.ListenAndServeTLS("", "")
		if err != nil {
			return err
		}
	}
	fmt.Println("http")
	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerTls) TlsConfig() bool {
	wd , err := os.Getwd()
	fmt.Println(wd)
	if err != nil {
		cfgTls = false
		return false
		}


	cert, err := tls.LoadX509KeyPair(fmt.Sprintf("%v/tls/localhost.crt", wd), fmt.Sprintf("%v/tls/localhost.key", wd))
	if err != nil {
		fmt.Println(err)
		cfgTls = false
		return false
	}
	s.Server.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	return true

}

//NewClient return a new instance of client
func NewServerTLS() *ServerTls {
	return &ServerTls{
		Server: &http.Server{},
	}
}
