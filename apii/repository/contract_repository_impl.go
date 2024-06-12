package repository

import (
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
)

type EduManage struct {
	// Define the methods and properties for the EduManage contract
}

type ContractRepositoryImpl struct {
	db        *gorm.DB
	client    *ethclient.Client
	authh     *bind.TransactOpts
	linkToken *bind.BoundContract
	certNFT   *bind.BoundContract
	eduManage *bind.BoundContract
}

// loadABI reads and parses the ABI file.
func loadABI(file string) (abi.ABI, error) {
	abiFile, err := os.ReadFile(file)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to read ABI file: %v", err)
	}

	contractABI, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("failed to parse ABI: %v", err)
	}
	return contractABI, nil
}

func NewContractRepositoryImpl(db *gorm.DB) (*ContractRepositoryImpl, error) {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	privateKeyHex := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	authh, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337)) // Hardhat uses chain ID 31337
	if err != nil {
		return nil, fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	linkTokenAddress := common.HexToAddress("0x5fbdb2315678afecb367f032d93f642f64180aa3")
	certNFTAddress := common.HexToAddress("0xe7f1725e7734ce288f8367e1bb143e90bb3f0512")
	eduManageAddress := common.HexToAddress("0x9fe46736679d2d9a65f0992f2272de9f3c7fa6e0")

	linkTokenABI, err := loadABI("/home/linh/DATN/thesis/thesis-/apii/smartContract/artifacts/contracts/LINKToken.sol/LINKToken.json")
	if err != nil {
		return nil, err
	}
	certNFTABI, err := loadABI("/home/linh/DATN/thesis/thesis-/apii/smartContract/artifacts/contracts/CertNFT.sol/CertificateNFT.json")
	if err != nil {
		return nil, err
	}
	eduManageABI, err := loadABI("/home/linh/DATN/thesis/thesis-/apii/smartContract/artifacts/contracts/EduManage.sol/EduManage.json")
	if err != nil {
		return nil, err
	}

	linkTokenInstance := bind.NewBoundContract(linkTokenAddress, linkTokenABI, client, client, client)
	certNFTInstance := bind.NewBoundContract(certNFTAddress, certNFTABI, client, client, client)
	eduManageInstance := bind.NewBoundContract(eduManageAddress, eduManageABI, client, client, client)

	// // Example of calling a function from the LINKToken contract
	// totalSupply := new(big.Int)
	// err = linkTokenInstance.Call(nil, totalSupply, "totalSupply")
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to retrieve total supply: %v", err)
	// }
	// fmt.Printf("Total Supply: %s\n", totalSupply.String())

	return &ContractRepositoryImpl{
		db:        db,
		client:    client,
		authh:     authh,
		linkToken: linkTokenInstance,
		certNFT:   certNFTInstance,
		eduManage: eduManageInstance,
	}, nil
}

func (r *ContractRepositoryImpl) GetLinkToken() *bind.BoundContract {
	return r.linkToken
}
