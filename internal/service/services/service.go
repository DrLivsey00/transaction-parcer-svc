package services

import (
	"math/big"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/db"
	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/requests"
	"github.com/DrLivsey00/transaction-parcer-svc/resources"
)

type StorageService interface {
	SaveTransfer(transfer resources.Transfer) error
	GetTransferBySenderTx(senderTx string) ([]resources.Transfer, error)
	GetTransferByReceiverTx(receiverTx string) ([]resources.Transfer, error)
	GetTransfers(filters requests.TransferRequest) ([]resources.Transfer, int, error)
	GetLatestBlockNumber() (*big.Int, error)
}

type Services struct {
	StorageService
}

func NewServices(repo *db.Repository) *Services {
	return &Services{
		StorageService: newStorageService(repo),
	}
}
