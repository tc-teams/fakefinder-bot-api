package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
)

type ServerTls struct {
	*http.Server
	*tls.Config
}

var cfgTls bool

const port = ":8080"

func (s *ServerTls) addr() string {
	s.Addr = port
	return s.Addr
}
func (s *ServerTls) Start(c context.Context, h http.Handler) error {
	s.Handler = h

	fmt.Printf("Create a sample server, port %s , with tls : %v", s.addr(), s.TlsConfig())

	if cfgTls {
		err := s.ListenAndServeTLS("", "")
		if err != nil {
			return err
		}
	}
	err := s.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *ServerTls) TlsConfig() bool {

	cert, err := tls.LoadX509KeyPair("/root/go-work/src/github.dxc.com/projects/fakefinder-bot-api/server/tls/localhost.crt", "/root/go-work/src/github.dxc.com/projects/fakefinder-bot-api/server/tls/localhost.key")
	if err != nil {
		return false
	}
	s.TLSConfig = &tls.Config{
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
