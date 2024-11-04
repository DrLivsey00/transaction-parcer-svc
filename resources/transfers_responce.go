package resources

type TransferResponce struct {
	Links TransferLinks  `json:"links"`
	Data  []TransferData `json:"data"`
}

type TransferLinks struct {
	Self string `json:"self"`
	Next string `json:"next"`
	Last string `json:"last"`
}

type TransferData struct {
	Id         int      `json:"id"`
	Attributes Transfer `json:"attributes"`
}
