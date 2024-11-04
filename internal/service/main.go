package service

import (
	"net"
	"net/http"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/config"
	"github.com/DrLivsey00/transaction-parcer-svc/internal/parser"
	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/services"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
	services *services.Services
	parser   parser.Parser
}

func startParser(cfg config.Config, services *services.Services) {
	parser := parser.NewParser(cfg, services)
	parser.Start()
}
func startHTTP(cfg config.Config, srv *services.Services, errorChan chan<- error) {
	if err := newService(cfg, srv).run(); err != nil {
		errorChan <- err
	}
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()
	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config, srv *services.Services) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
		services: srv,
		parser:   parser.NewParser(cfg, srv),
	}
}

func Run(cfg config.Config, srv *services.Services) {
	errorChan := make(chan error)
	go startParser(cfg, srv)
	go startHTTP(cfg, srv, errorChan)
	for err := range errorChan {
		panic(err)
	}
}
