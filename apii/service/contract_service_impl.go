package service

import (
	"fmt"
	"go-pzn-restful-api/repository"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ContractServiceImpl struct {
	repository.ContractRepository
}

func (s *ContractServiceImpl) BalanceOf(address string) (*big.Int, error) {
	addr := common.HexToAddress(address)
	var result []interface{}
	callOpts := &bind.CallOpts{} 

	err := s.ContractRepository.GetLinkToken().Call(callOpts, &result, "balanceOf", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve balance: %v", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no result returned")
	}

	balance, ok := result[0].(*big.Int)
	if !ok {
		return nil, fmt.Errorf("unexpected type for balance result")
	}
	return balance, nil
}


func NewContractService(contractRepository repository.ContractRepository) ContractService {
	return &ContractServiceImpl{
		ContractRepository: contractRepository,
	}
}
