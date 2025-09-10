package server

import (
	"errors"
	"io"
	"log"
	"time"

	"github.com/alvinfadli/cnd/apps/smtp-server/internal/config"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
)

type Backend struct{}

func (bkd *Backend) NewSession(c *smtp.Conn) (smtp.Session, error) {
	return &Session{}, nil
}

type Session struct {
	auth bool
}

func (s *Session) AuthMechanisms() []string {
	return []string{sasl.Plain}
}

func (s *Session) Auth(mech string) (sasl.Server, error) {
	return sasl.NewPlainServer(func(identity, username, password string) error {
		if username != "username" || password != "password" {
			return errors.New("Invalid username or password")
		}
		s.auth = true
		return nil
	}), nil
}

func (s *Session) Mail(from string, opts *smtp.MailOptions) error {
	if !s.auth {
		return smtp.ErrAuthRequired
	}
	log.Println("Mail from:", from)
	return nil
}

func (s *Session) Rcpt(to string, opts *smtp.RcptOptions) error {
	if !s.auth {
		return smtp.ErrAuthRequired
	}
	log.Println("Rcpt to:", to)
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if !s.auth {
		return smtp.ErrAuthRequired
	}
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

func Start(cfg *config.Config) {
	be := &Backend{}

	s := smtp.NewServer(be)

	ioTimeout := time.Duration(10) * time.Second

	s.Addr = cfg.Domain + ":" + cfg.Port
	s.Domain = cfg.Domain
	s.WriteTimeout = ioTimeout
	s.ReadTimeout = ioTimeout
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = cfg.MaxRecipients
	s.AllowInsecureAuth = cfg.AllowInsecureAuth

	log.Println("Starting server at", s.Addr, "with domain", cfg.Domain)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}