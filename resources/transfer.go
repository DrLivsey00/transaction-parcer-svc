package resources

type Transfer struct {
	TransactionHash string `db:"tx_hash"`
	From            string `db:"sender"`
	To              string `db:"receiver"`
	TokenAmount     string `db:"token_amount"`
	EventIndex      uint   `db:"event_index"`
	BlockNumber     uint   `db:"block_number"`
}
