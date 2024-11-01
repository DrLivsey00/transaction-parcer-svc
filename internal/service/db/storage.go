package db

import (
	"errors"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/config"
	"github.com/DrLivsey00/transaction-parcer-svc/resources"
	"github.com/Masterminds/squirrel"
)

type dbStorage struct {
	config.Config
}

func newStorage(cfg config.Config) *dbStorage {
	return &dbStorage{
		cfg,
	}
}

func (s *dbStorage) AddTransfer(t resources.Transfer) error {
	err := s.DB().Exec(squirrel.Insert("transfers").
		Columns("tx_hash", "sender", "receiver", "token_amount").
		Values(t.TransactionHash, t.From, t.To, t.Token_amount))
	if err != nil {
		//s.Log().Errorf("failed to save transfer: %s", err.Error())
		return err
	}
	//s.Log().Info("succesfully saved transfer.")
	return nil
}

func (s *dbStorage) GetBySender(senderTx string) ([]resources.Transfer, error) {
	var transfers []resources.Transfer
	s.Log().Infof("Incoming txHash: %s", senderTx)
	err := s.DB().Select(&transfers, squirrel.Select("tx_hash", "sender", "receiver", "token_amount").
		From("transfers").
		Where(squirrel.Eq{"sender": senderTx}))
	if err != nil {
		s.Log().Errorf("failed to get transfers with senderTx: %s", err.Error())
		return nil, err
	}
	if len(transfers) == 0 {
		s.Log().Error("No trnsfers found")
		return nil, errors.New("no transfers found")
	}
	s.Log().Info(transfers)
	s.Log().Info("succesfully found transfers.")
	return transfers, nil
}

func (s *dbStorage) GetByReceiver(receiverTx string) ([]resources.Transfer, error) {
	var transfers []resources.Transfer
	err := s.DB().Select(&transfers, squirrel.Select("tx_hash", "sender", "receiver", "token_amount").
		From("transfers").
		Where(squirrel.Eq{"receiver": receiverTx}))
	if err != nil {
		s.Log().Errorf("failed to get transfers with senderTx: %s", err.Error())
		return nil, err
	}
	if len(transfers) == 0 {
		s.Log().Error("No trnsfers found")
		return nil, errors.New("no transfers found")
	}
	s.Log().Info("succesfully found transfers.")
	return transfers, nil
}
