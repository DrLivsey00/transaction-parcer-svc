package requests

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DrLivsey00/transaction-parcer-svc/resources"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/urlval"
)

func ParseQueryParams(r *http.Request) (string, error) {
	params := chi.URLParam(r, "txHash")
	if params == "" {
		return "", errors.New("empty params")
	}
	return params, nil

}

type TransferRequest struct {
	FilterType []resources.FilterType `filter:"filter"`
}

func NewTransferRequest(r *http.Request) (TransferRequest, error) {
	request := TransferRequest{}
	err := urlval.DecodeSilently(r.URL.Query(), &request)
	if err != nil {
		return request, fmt.Errorf("invalid request params: %s", err.Error())
	}
	return request, nil
}
