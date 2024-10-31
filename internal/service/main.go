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

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()
	s.parser.Parse()

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
	if err := newService(cfg, srv).run(); err != nil {
		panic(err)
	}
}
