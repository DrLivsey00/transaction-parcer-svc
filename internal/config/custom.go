package config

import (
	"github.com/DrLivsey00/transaction-parcer-svc/internal/config/custom"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type Customer interface {
	Custom() custom.Custom
}

type customer struct {
	getter kv.Getter
	once   comfig.Once
}

func NewCustomer(getter kv.Getter) Customer {
	return &customer{
		getter: getter,
	}
}

func (c *customer) Custom() custom.Custom {
	return c.once.Do(func() interface{} {
		var config struct {
			db_url          string `fig:"url"`
			infuraKey       string `fig:"infura_api_key"`
			contractAddress string `fig: "contract_address"`
		}
		if err := figure.Out(&config).From(kv.MustGetStringMap(c.getter, "custom")).Please(); err != nil {
			panic("error getting custom config: " + err.Error())
		}
		if err := figure.Out(&config).From(kv.MustGetStringMap(c.getter, "db")).Please(); err != nil {
			panic("error gettting db config" + err.Error())
		}
		custom := custom.New(config.db_url, config.infuraKey, config.contractAddress)
		return custom
	}).(custom.Custom)

}