package custom

type Custom struct {
	DbUrl        string
	InfuraApiKey string
	Contract     string
}

func New(dbUrl, infuraKey, contract string) Custom {
	return Custom{
		DbUrl:        dbUrl,
		InfuraApiKey: infuraKey,
		Contract:     contract,
	}
}
