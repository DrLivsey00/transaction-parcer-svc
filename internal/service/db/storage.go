package db

import (
	"github.com/DrLivsey00/transaction-parcer-svc/internal/config"
	"github.com/DrLivsey00/transaction-parcer-svc/resources"
	"github.com/Masterminds/squirrel"
)

type DbStorage struct {
	config.Config
}

func NewStorage(cfg config.Config) *DbStorage {
	return &DbStorage{
		cfg,
	}
}

func (s *DbStorage) AddTransfer(t resources.Transfer) error {
	err := s.DB().Exec(squirrel.Insert("transfers").
		Columns("tx_hash", "sender", "receiver", "token_amount").
		Values(t.TransactionHash, t.From, t.To, t.Token_amount))
	if err != nil {
		s.Log().Errorf("failed to save transfer: %s", err.Error())
		return err
	}
	s.Log().Info("succesfully saved transfer.")
	return nil
}

func (s *DbStorage) GetBySender(senderTx string) ([]resources.Transfer, error) {
	var transfers []resources.Transfer
	err := s.DB().Select(&transfers, squirrel.Select("tx_hash", "sender", "receiver", "token_amount").
		From("transfers").
		Where(squirrel.Eq{"sender": senderTx}))
	if err != nil {
		s.Log().Errorf("failed to get transfers with senderTx: %s", err.Error())
	}
	s.Log().Info("succesfully found transfers.")
	return transfers, nil
}

func (s *DbStorage) GetByReceiver(receiverTx string) ([]resources.Transfer, error) {
	var transfers []resources.Transfer
	err := s.DB().Select(&transfers, squirrel.Select("tx_hash", "sender", "receiver", "token_amount").
		From("transfers").
		Where(squirrel.Eq{"receiver": receiverTx}))
	if err != nil {
		s.Log().Errorf("failed to get transfers with senderTx: %s", err.Error())
	}
	s.Log().Info("succesfully found transfers.")
	return transfers, nil
}
