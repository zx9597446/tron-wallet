package grpcClient

import (
	"bytes"
	"fmt"
	"github.com/zx9597446/tron-wallet/grpcClient/proto/api"
	"github.com/zx9597446/tron-wallet/grpcClient/proto/core"
	"github.com/zx9597446/tron-wallet/util"
)

func (g *GrpcClient) GetAccount(addr string) (*core.Account, error) {
	account := new(core.Account)
	var err error

	account.Address, err = util.DecodeCheck(addr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := g.getContext()
	defer cancel()

	acc, err := g.Client.GetAccount(ctx, account)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(acc.Address, account.Address) {
		return nil, fmt.Errorf("account not found")
	}
	return acc, nil
}

func (g *GrpcClient) GetAccountResource(addr string) (*api.AccountResourceMessage, error) {
	account := new(core.Account)
	var err error

	account.Address, err = util.DecodeCheck(addr)
	if err != nil {
		return nil, err
	}

	ctx, cancel := g.getContext()
	defer cancel()

	acc, err := g.Client.GetAccountResource(ctx, account)
	if err != nil {
		return nil, err
	}

	return acc, nil
}
