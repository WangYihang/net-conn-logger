package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"errors"

	"github.com/WangYihang/go-net-conn-logger/pkg/listener"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

// The following code is shamelessly stolen from go-smtp official repo [1]
// [1] https://github.com/emersion/go-smtp/blob/master/example_server_test.go

// The Backend implements SMTP server methods.
type Backend struct{}

// NewSession is called after client greeting (EHLO, HELO).
func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

// A Session is returned after successful login.
type Session struct{}

// AuthMechanisms returns a slice of available auth mechanisms; only PLAIN is
// supported in this example.
func (s *Session) AuthMechanisms() []string {
	return []string{sasl.Plain}
}

// Auth is the handler for supported authenticators.
func (s *Session) Auth(mech string) (sasl.Server, error) {
	return sasl.NewPlainServer(func(identity, username, password string) error {
		if username != "username" || password != "password" {
			return errors.New("invalid username or password")
		}
		return nil
	}), nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := io.ReadAll(r); err != nil {
		return err
	} else {
		log.Println("Data:", string(b))
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func main() {
	domain := "localhost"
	addr := "localhost:1025"

	// Create a new listener
	l, err := listener.NewListener(addr)
	if err != nil {
		panic(err)
	}

	// Start the log consumer
	go func() {
		for logEntry := range l.LogChan() {
			data, _ := json.Marshal(logEntry)
			fmt.Println(string(data))
		}
	}()

	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = addr
	s.Domain = domain
	s.WriteTimeout = 10 * time.Second
	s.ReadTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Addr)
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
