package services

import "github.com/DrLivsey00/transaction-parcer-svc/resources"

type StorageService interface {
	SaveTransfer(transfer resources.Transfer) error
	GetTransferBySenderTx(senderTx string) ([]resources.Transfer, error)
	GetTransferByReceiverTx(receiverTx string) ([]resources.Transfer, error)
}

type Services struct {
	StorageService
}
