package service

import (
	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/handlers"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxService(s.services),
		),
	)
	r.Route("/integrations/transac-parser-svc", func(r chi.Router) {
		r.Get("/from/{txHash}", handlers.FindBySender)
		r.Get("/to/{txHash}", handlers.FindByreceiver)
		r.Get("/transfers", handlers.GetTransfers)
		// configure endpoints here
	})

	return r
}
