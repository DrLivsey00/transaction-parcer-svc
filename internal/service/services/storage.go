package services

import (
	"math/big"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/db"
	"github.com/DrLivsey00/transaction-parcer-svc/resources"
)

type storageSrv struct {
	repo *db.Repository
}

func newStorageService(repo *db.Repository) *storageSrv {
	return &storageSrv{
		repo: repo,
	}
}

func (s *storageSrv) SaveTransfer(transfer resources.Transfer) error {
	return s.repo.AddTransfer(transfer)
}
func (s *storageSrv) GetTransferBySenderTx(senderTx string) ([]resources.Transfer, error) {
	return s.repo.GetBySender(senderTx)
}
func (s *storageSrv) GetTransferByReceiverTx(receiverTx string) ([]resources.Transfer, error) {
	return s.repo.GetByReceiver(receiverTx)
}
func (s *storageSrv) GetLatestBlockNumber() (*big.Int, error) {
	return s.repo.GetLatestBlockNumber()
}
