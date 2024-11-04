package custom

type Custom struct {
	WssApiKey  string
	HttpApiKey string
	Contract   string
	DomainName string
}

func New(wssKey, httpKey, contract, domainName string) Custom {
	return Custom{
		WssApiKey:  wssKey,
		HttpApiKey: httpKey,
		Contract:   contract,
		DomainName: domainName,
	}
}
