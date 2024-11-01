package handlers

import (
	"net/http"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/requests"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func FindBySender(w http.ResponseWriter, r *http.Request) {
	params, err := requests.ParseQueryParams(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data": err,
		})...)
		return
	}
	srv := Service(r)
	if srv == nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	transfers, err := srv.GetTransferBySenderTx(params)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(transfers) == 0 {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, transfers)
}

func FindByreceiver(w http.ResponseWriter, r *http.Request) {
	params, err := requests.ParseQueryParams(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data": err,
		})...)
		return
	}
	srv := Service(r)
	if srv == nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	transfers, err := srv.GetTransferByReceiverTx(params)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if len(transfers) == 0 {
		ape.RenderErr(w, problems.NotFound())
		return
	}
	ape.Render(w, transfers)
}
