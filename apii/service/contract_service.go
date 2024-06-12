package service

import (
	"math/big"
)

type ContractService interface {
	BalanceOf(address string) (*big.Int, error)
}
