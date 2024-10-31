package requests

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
)

func ParseQueryParams(r *http.Request) (string, error) {
	params := chi.URLParam(r, "txHash")
	if params == "" {
		return "", errors.New("empty params")
	}
	return params, nil

}
