package erc20

import (
	"context"
	"math/big"

	_ "embed"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

//go:embed abi/IERC20.abi
var erc20ABI string

type ERC20Interface interface {
	BalanceOf(ctx context.Context, address common.Address) (*big.Int, error)
	Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error)
	Transfer(ctx context.Context, to common.Address, amount *big.Int) (*types.Transaction, error)
	Approve(ctx context.Context, spender common.Address, amount *big.Int) (*types.Transaction, error)
}

type ERC20Client struct {
	client   *ethclient.Client
	contract *IERC20
}

func NewERC20Client(url string, contractAddress common.Address) (ERC20Interface, error) {
	client, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}
	contract, err := NewIERC20(contractAddress, client)
	if err != nil {
		return nil, err
	}
	return &ERC20Client{client: client, contract: contract}, nil
}

func (c *ERC20Client) BalanceOf(ctx context.Context, address common.Address) (*big.Int, error) {
	return c.contract.BalanceOf(nil, address)
}

func (c *ERC20Client) Allowance(ctx context.Context, owner, spender common.Address) (*big.Int, error) {
	return c.contract.Allowance(nil, owner, spender)
}

func (c *ERC20Client) Transfer(ctx context.Context, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return c.contract.Transfer(nil, to, amount)
}

func (c *ERC20Client) Approve(ctx context.Context, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return c.contract.Approve(nil, spender, amount)
}
