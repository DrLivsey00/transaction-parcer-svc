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

	var nextPage int
	if *request.Page+1 <= pages {
		nextPage = *request.Page + 1
	} else {
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

	data := make([]resources.TransferData, len(transfers))
	for i, transfer := range transfers {
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
	nextUrl, err := url.Parse(baseURL)
	if err != nil {
		return "", "", "", err
	}
	selfUrl := nextUrl
	lastUrl := nextUrl
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
	nextParams := params
	selfParams := params
	lastParams := params

	nextParams.Set("filter[offset]", fmt.Sprintf("%d", nextPage))
	selfParams.Set("filter[offset]", fmt.Sprintf("%d", *request.Page))
	lastParams.Set("filter[offset]", fmt.Sprintf("%d", lastPage))

	nextUrl.RawQuery = nextParams.Encode()
	selfUrl.RawQuery = selfParams.Encode()
	lastUrl.RawQuery = selfParams.Encode()

	return nextUrl.String(), selfUrl.String(), lastUrl.String(), nil
}
