package db

import (
	"math/big"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/config"
	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/requests"
	"github.com/DrLivsey00/transaction-parcer-svc/resources"
)

type Storage interface {
	AddTransfer(t resources.Transfer) error
	GetBySender(senderTx string) ([]resources.Transfer, error)
	GetByReceiver(receiverTx string) ([]resources.Transfer, error)
	GetLatestBlockNumber() (*big.Int, error)
	GetTransfers(filters requests.TransferRequest) ([]resources.Transfer, int, error)
}

type Repository struct {
	Storage
}

func NewRepo(cfg config.Config) *Repository {
	return &Repository{
		Storage: newStorage(cfg),
	}
}
