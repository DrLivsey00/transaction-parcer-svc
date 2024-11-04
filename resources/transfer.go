package resources

type Transfer struct {
	Id              int    `db:"id" json:""`
	TransactionHash string `db:"tx_hash" json:"tx_hash"`
	From            string `db:"sender" json:"sender"`
	To              string `db:"receiver" json:"receiver"`
	TokenAmount     string `db:"token_amount" json:""`
	EventIndex      uint   `db:"event_index"`
	BlockNumber     uint   `db:"block_number"`
}
