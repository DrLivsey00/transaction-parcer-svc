package parser

import (
	"context"
	"errors"

	"github.com/DrLivsey00/transaction-parcer-svc/internal/config"
	"github.com/DrLivsey00/transaction-parcer-svc/internal/service/services"
	"github.com/DrLivsey00/transaction-parcer-svc/resources"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Parser interface {
	Start()
}

type parser struct {
	srv *services.Services
	cfg config.Config
}

func (p *parser) Start() {
	err := p.recoverMissedTransfers()
	if err != nil {
		panic(err)
	}
	go p.parse()
}

func (p *parser) parse() {
	// p.cfg.Log().Infof("Infura Wss API Url: %s", p.cfg.Custom().WssApiKey)
	// p.cfg.Log().Infof("Contract Address: %s", p.cfg.Custom().Contract)
	client, err := ethclient.Dial(p.cfg.Custom().WssApiKey)
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
		transfer.TokenAmount = event.Tokens.String()
		transfer.EventIndex = event.Raw.Index
		transfer.BlockNumber = uint(event.Raw.BlockNumber)

		//p.cfg.Log().Infof("Saving transfer: %+v", transfer)

		err := p.srv.SaveTransfer(transfer)
		if err != nil {
			p.cfg.Log().Error(err)
			panic(err)
		}

	}
}

func NewParser(cfg config.Config, srv *services.Services) Parser {
	return &parser{
		srv: srv,
		cfg: cfg,
	}
}

//reload

func (p *parser) recoverMissedTransfers() error {
	httpApiurl := p.cfg.Custom().HttpApiKey

	httpClient, err := ethclient.Dial(httpApiurl) //http client initialization
	if err != nil {
		return errors.New("error fetching eth client")
	}

	contractAddress := common.HexToAddress(p.cfg.Custom().Contract)

	latestBlock, err := p.srv.GetLatestBlockNumber() //get the last block recorded in db
	if err != nil {
		return err
	}

	endBlock, err := httpClient.BlockNumber(context.Background()) // get latest block from blockachain
	if err != nil {
		return err
	}

	tokenFilter, err := NewTokenFilter(contractAddress, httpClient) //initialize token filter
	if err != nil {
		return err
	}
	opts := &bind.FilterOpts{
		Start:   latestBlock.Uint64(),
		End:     &endBlock,
		Context: context.Background(),
	}

	iterator, err := tokenFilter.FilterTransfer(opts, nil, nil) //fetching transfers
	if err != nil {
		return err
	}

	var transfer resources.Transfer
	for iterator.Next() {
		event := iterator.Event

		transfer.From = event.From.Hex()
		transfer.To = event.To.Hex()
		transfer.TransactionHash = event.Raw.TxHash.Hex()
		transfer.TokenAmount = event.Tokens.String()
		transfer.EventIndex = event.Raw.Index
		transfer.BlockNumber = uint(event.Raw.BlockNumber)

		err := p.srv.SaveTransfer(transfer)
		if err != nil {
			p.cfg.Log().Errorf("Failed to save transfer: %v", err)
		}
	}

	if err = iterator.Error(); err != nil {
		return err
	}

	return nil
}
