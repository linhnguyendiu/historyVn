package repository

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type ContractRepository interface {
	GetLinkToken() *bind.BoundContract
}
