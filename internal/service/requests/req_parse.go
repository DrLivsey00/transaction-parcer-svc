package requests

import (
	"errors"
	"fmt"
	"net/http"

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
	FromAdresses []string `filter:"from"`
	ToAdresses   []string `filter:"to"`
	Counterparty []string `filter:"counterparty"`
}

func NewTransferRequest(r *http.Request) (TransferRequest, error) {
	request := TransferRequest{}
	err := urlval.DecodeSilently(r.URL.Query(), &request)
	if err != nil {
		return request, fmt.Errorf("invalid request params: %s", err.Error())
	}

	err = validateFilters(request.FromAdresses, request.ToAdresses, request.Counterparty)

	if err != nil {
		return request, fmt.Errorf("invalid request params: %s", err.Error())
	}

	return request, nil
}

func validateFilters(fromAdresses, toAdresses, counterPartyAddresses []string) error {
	var hasFrom, hasTo, hasCounterparty bool
	hasFrom = len(fromAdresses) > 0
	hasTo = len(toAdresses) > 0
	hasCounterparty = len(counterPartyAddresses) > 0
	if (hasTo && hasCounterparty && hasFrom) || (hasFrom && hasCounterparty) || (hasTo && hasCounterparty) {
		return errors.New("too many filters")
	}
	return nil
}
