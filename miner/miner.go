/*
	author: licowei
	description: ubbey公链miner接口，web3风格封装http请求。
*/
package miner

import (
	"fmt"
	"go-web3/dto"
	"go-web3/providers"
)

type Miner struct {
	provider providers.ProviderInterface
}

func NewMiner(provider providers.ProviderInterface) *Miner {
	plotter := new(Miner)
	plotter.provider = provider
	return plotter
}

func (p *Miner) Start() error {
	params := []string{}
	pointer := &dto.RequestResult{}

	err := p.provider.SendRequest(&pointer, "miner_start", params)

	if err != nil {
		fmt.Printf("start mine , http req err : %s\n", err.Error())
		return err
	}

	return nil
}

func (p *Miner) Stop() error {
	params := []string{}
	pointer := &dto.RequestResult{}

	err := p.provider.SendRequest(&pointer, "miner_stop", params)

	if err != nil {
		fmt.Printf("stop mine , http req err : %s\n", err.Error())
		return err
	}

	return nil
}

func (p *Miner) SetEtherbase(addr string) error {
	params := []string{}
	params = append(params, addr)
	pointer := &dto.RequestResult{}

	err := p.provider.SendRequest(&pointer, "miner_setEtherbase", params)

	if err != nil {
		fmt.Printf("SetEtherbase , http req err : %s\n", err.Error())
		return err
	}

	return nil
}
