package parser

import (
	"context"
	"fmt"
	"math/big"

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
	p.cfg.Log().Info("Started recovering...")
	err := p.recoverMissedTransfers()
	if err != nil {
		p.cfg.Log().Info("skipping getting transfers")
		p.cfg.Log().Error(err)
	}
	p.cfg.Log().Info("Recovery ended succesfully, now you are on websockets.")
	p.parse()
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
	httpApiUrl := p.cfg.Custom().HttpApiKey

	//init http client
	httpClient, err := ethclient.Dial(httpApiUrl)
	if err != nil {
		return fmt.Errorf("failed to initialize Ethereum client: %w", err)
	}

	contractAddress := common.HexToAddress(p.cfg.Custom().Contract)

	//get the latest block recorded in the database
	latestBlock, err := p.srv.GetLatestBlockNumber()
	if err != nil {
		return fmt.Errorf("failed to get the latest block from DB: %w", err)
	}

	//get the latest block on the blockchain
	endBlock, err := httpClient.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get the latest block from blockchain: %w", err)
	}

	//set block range
	blockRange := big.NewInt(1000)
	tokenFilter, err := NewTokenFilter(contractAddress, httpClient)
	if err != nil {
		return fmt.Errorf("failed to initialize token filter: %w", err)
	}

	for latestBlock.Cmp(big.NewInt(int64(endBlock))) < 0 {

		localEndBlock := new(big.Int).Add(latestBlock, blockRange)
		if localEndBlock.Cmp(big.NewInt(int64(endBlock))) > 0 {
			localEndBlock = big.NewInt(int64(endBlock))
		}
		endbl := localEndBlock.Uint64()

		//update filter options
		opts := &bind.FilterOpts{
			Start:   latestBlock.Uint64(),
			End:     &endbl,
			Context: context.Background(),
		}
		iterator, err := tokenFilter.FilterTransfer(opts, nil, nil)
		if err != nil {
			return fmt.Errorf("failed to filter transfers: %w", err)
		}

		//save transfers
		for iterator.Next() {
			event := iterator.Event
			transfer := resources.Transfer{
				From:            event.From.Hex(),
				To:              event.To.Hex(),
				TransactionHash: event.Raw.TxHash.Hex(),
				TokenAmount:     event.Tokens.String(),
				EventIndex:      event.Raw.Index,
				BlockNumber:     uint(event.Raw.BlockNumber),
			}

			if err := p.srv.SaveTransfer(transfer); err != nil {
				p.cfg.Log().Errorf("Failed to save transfer: %v", err)
			}
		}

		if err := iterator.Error(); err != nil {
			return fmt.Errorf("iterator error: %w", err)
		}

		latestBlock.Add(latestBlock, blockRange)

		if latestBlock.Cmp(big.NewInt(int64(endBlock))) >= 0 {
			break
		}
	}

	return nil
}
