package parser

import (
	"context"
	"math/big"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/config"
	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/services"
	"github.com/DrLivsey00/transaction-parcer-svc/resources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Parser interface {
	Parse()
}

type parser struct {
	srv *services.Services
	cfg config.Config
}

func (p *parser) Parse() {
	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/" + p.cfg.Custom().InfuraApiKey)
	if err != nil {
		panic(err)
	}
	contractAdress := common.HexToAddress(p.cfg.Custom().Contract)

	tokenFilter, err := NewTokenFilter(contractAdress, client)
	if err != nil {
		panic(err)
	}

	sinc := make(chan *TokenFilterTransfer)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	watchOpts := &bind.WatchOpts{
		Context: ctx,
	}

	sub, err := tokenFilter.WatchTransfer(watchOpts, sinc, nil, nil)
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	go func() {
		transfer := resources.Transfer{}
		for event := range sinc {
			transfer.From = event.From.Hex()
			transfer.To = event.To.Hex()
			transfer.TransactionHash = event.Raw.TxHash.Hex()
			tokenAmountFloat := new(big.Float).SetInt(event.Tokens)
			tokenAmountFloat.Quo(tokenAmountFloat, big.NewFloat(1e18))
			tokenAmount, _ := tokenAmountFloat.Float64()
			transfer.Token_amount = tokenAmount
			err := p.srv.SaveTransfer(transfer)
			if err != nil {
				p.cfg.Log().Error(err)
			}

		}
	}()
}

func NewParser(cfg config.Config, srv *services.Services) Parser {
	return &parser{
		srv: srv,
		cfg: cfg,
	}
}