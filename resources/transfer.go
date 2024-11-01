package resources

type Transfer struct {
	TransactionHash string  `db:"tx_hash"`
	From            string  `db:"sender"`
	To              string  `db:"receiver"`
	Token_amount    float64 `db:"token_amount"`
}
