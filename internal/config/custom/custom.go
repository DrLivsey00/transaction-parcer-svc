package custom

type Custom struct {
	WssApiKey  string
	HttpApiKey string
	Contract   string
}

func New(wssKey, httpKey, contract string) Custom {
	return Custom{
		WssApiKey:  wssKey,
		HttpApiKey: httpKey,
		Contract:   contract,
	}
}
