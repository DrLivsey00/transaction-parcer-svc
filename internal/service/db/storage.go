package db

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/config"
	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/requests"
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
		Columns("tx_hash", "sender", "receiver", "token_amount", "block_number", "event_index").
		Values(t.TransactionHash, t.From, t.To, t.TokenAmount, t.BlockNumber, t.EventIndex))
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

//New FilterFunc

func (s *dbStorage) GetTransfers(filters requests.TransferRequest) ([]resources.Transfer, int, error) {
	var transfers []resources.Transfer
	var transfersNumber int
	offset := *filters.PageSize * (*filters.Page - 1)

	query := squirrel.Select("tx_hash", "sender", "receiver", "token_amount", "block_number", "event_index").From("transfers")
	allTransfers := squirrel.Select("COUNT(*)").From("transfers")

	if filters.FromAdresses != nil {
		query = query.Where(squirrel.Eq{"sender": filters.FromAdresses})
		allTransfers = allTransfers.Where(squirrel.Eq{"sender": filters.FromAdresses})
	}
	if filters.ToAdresses != nil {
		query = query.Where(squirrel.Eq{"receiver": filters.ToAdresses})
		allTransfers = allTransfers.Where(squirrel.Eq{"receiver": filters.ToAdresses})
	}
	if filters.Counterparty != nil {
		query = query.Where(squirrel.Or{
			squirrel.Eq{"sender": filters.Counterparty},
			squirrel.Eq{"receiver": filters.Counterparty},
		})
		allTransfers = allTransfers.Where(squirrel.Or{
			squirrel.Eq{"sender": filters.Counterparty},
			squirrel.Eq{"receiver": filters.Counterparty},
		})
	}
	err := s.DB().Select(&transfersNumber, allTransfers)
	if err != nil {
		return nil, 0, fmt.Errorf("error counting transfers: %s", err.Error())
	}

	query = query.Limit(25).Offset(uint64(offset))

	err = s.DB().Select(&transfers, query)

	if err != nil {
		return nil, 0, fmt.Errorf("error finding transfers: %s", err.Error())
	}

	if len(transfers) == 0 {
		return nil, 0, fmt.Errorf("no transfers found: %s", err.Error())
	}

	return transfers, 0, nil
}

//Get latest block bunc

func (s *dbStorage) GetLatestBlockNumber() (*big.Int, error) {
	var BlockNumber int64
	err := s.DB().Get(&BlockNumber, squirrel.Select("block_number").
		From("transfers").
		OrderBy("block_number DESC").
		Limit(1))
	if err != nil {

		return big.NewInt(0), err
	}

	return big.NewInt(BlockNumber), nil
}
