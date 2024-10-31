package custom

type Custom struct {
	InfuraApiKey string
	Contract     string
}

func New(infuraKey, contract string) Custom {
	return Custom{
		InfuraApiKey: infuraKey,
		Contract:     contract,
	}
}
