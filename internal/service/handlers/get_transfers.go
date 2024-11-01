package handlers

import (
	"net/http"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetTransfers(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewTransferRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	ape.Render(w, request)

}
