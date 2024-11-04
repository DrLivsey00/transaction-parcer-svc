package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/requests"
	"github.com/DrLivsey00/transaction-parcer-svc/resources"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func GetTransfers(w http.ResponseWriter, r *http.Request) {
	services := Service(r)
	logger := Log(r)
	cfg := GetConfig(r)

	request, err := requests.NewTransferRequest(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data": err,
		})...)
		return
	}
	transfers, pages, err := services.GetTransfers(request)
	logger.Infof("pages:%d", pages)

	//validating pages
	if *request.Page > pages {
		ape.RenderErr(w, problems.BadRequest(validation.Errors{
			"data": fmt.Errorf("page doesn`t exists"),
		})...)
		return
	}
	if err != nil {
		logger.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	pageSize := *request.PageSize
	offset := (*request.Page - 1) * pageSize
	end := offset + pageSize
	if end > len(transfers) {
		end = len(transfers)
	}

	pagedTransfers := transfers[offset:end]

	nextPage := *request.Page + 1
	if nextPage > pages {
		nextPage = pages
	}

	//Building links
	baseURL := fmt.Sprintf("%s/integrations/transac-parser-svc/transfers", cfg.Custom().DomainName)

	next, self, last, err := buildUrl(request, baseURL, nextPage, pages)
	if err != nil {
		logger.Error(err)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	data := make([]resources.TransferData, len(pagedTransfers))
	for i, transfer := range pagedTransfers {
		data[i] = resources.TransferData{
			Id:         transfer.Id,
			Attributes: transfer,
		}
	}

	ape.Render(w, resources.TransferResponce{
		Links: resources.TransferLinks{
			Next: next,
			Self: self,
			Last: last,
		},
		Data: data,
	})

}
func buildUrl(request requests.TransferRequest, baseURL string, nextPage, lastPage int) (string, string, string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", "", "", err
	}

	buildParams := func(page int) url.Values {
		params := url.Values{}
		if len(request.FromAdresses) > 0 {
			for _, address := range request.FromAdresses {
				params.Add("filter[from]", address)
			}
		}
		if len(request.ToAdresses) > 0 {
			for _, address := range request.ToAdresses {
				params.Add("filter[to]", address)
			}
		}
		if len(request.Counterparty) > 0 {
			for _, party := range request.Counterparty {
				params.Add("filter[counterparty]", party)
			}
		}
		params.Add("filter[page_size]", fmt.Sprintf("%d", *request.PageSize))
		params.Add("filter[offset]", fmt.Sprintf("%d", page))
		return params
	}

	nextUrl := *base
	selfUrl := *base
	lastUrl := *base
	nextUrl.RawQuery = buildParams(nextPage).Encode()
	selfUrl.RawQuery = buildParams(*request.Page).Encode()
	lastUrl.RawQuery = buildParams(lastPage).Encode()

	return nextUrl.String(), selfUrl.String(), lastUrl.String(), nil
}
