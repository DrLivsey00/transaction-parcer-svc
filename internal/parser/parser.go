package parser

import (
	"context"

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
	p.cfg.Log().Infof("Infura API Key: %s", p.cfg.Custom().InfuraApiKey)
	p.cfg.Log().Infof("Contract Address: %s", p.cfg.Custom().Contract)
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
	transfer := resources.Transfer{}
	for event := range sinc {
		//p.cfg.Log().Infof("Received transfer event: From %s, To %s, Hash %s, Tokens %d",
		//event.From.Hex(), event.To.Hex(), event.Raw.TxHash.Hex(), event.Tokens)
		transfer.From = event.From.Hex()
		transfer.To = event.To.Hex()
		transfer.TransactionHash = event.Raw.TxHash.Hex()
		transfer.Token_amount = event.Tokens.String()
		//p.cfg.Log().Infof("Saving transfer: %+v", transfer)

		err := p.srv.SaveTransfer(transfer)
		if err != nil {
			p.cfg.Log().Error(err)
		}

	}
}

func NewParser(cfg config.Config, srv *services.Services) Parser {
	return &parser{
		srv: srv,
		cfg: cfg,
	}
}
