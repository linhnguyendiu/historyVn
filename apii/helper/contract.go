package helper

import (
	"go-pzn-restful-api/contracts"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	Client *ethclient.Client
	Token  *contracts.LINKToken
	Cert   *contracts.CertNFT
	Manage *contracts.EduManage
)

func DialClient() *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("we have a connection")

	return client
}

func ConnectToLINKToken() {
	Client = DialClient()
	var err error
	linkTokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	Token, err = contracts.NewLINKToken(linkTokenAddress, Client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	log.Println("Connect to LINKToken successfully!!!")

	time.Sleep(250 * time.Millisecond)

	name, err := Token.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Name", name)
}

func ConnectToCertNFT() {
	Client = DialClient()
	var err error
	certNFTAddress := common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
	Cert, err = contracts.NewCertNFT(certNFTAddress, Client)
	if err != nil {
		log.Fatalf("Failed to instantiate a CertNFT contract: %v", err)
	}
	log.Println("Connect to CertNFT successfully!!!")

	time.Sleep(250 * time.Millisecond)
}

func ConnectToEduManage() {
	Client = DialClient()
	var err error
	eduManageAddress := common.HexToAddress("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")
	Manage, err = contracts.NewEduManage(eduManageAddress, Client)
	if err != nil {
		log.Fatalf("Failed to instantiate a EduManage contract: %v", err)
	}
	log.Println("Connect to EduManageAddress successfully!!!")

	time.Sleep(250 * time.Millisecond)
}

func AuthGenerator(client *ethclient.Client) *bind.TransactOpts {
	privateKeyHex := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(31337)) // Hardhat uses chain ID 31337
	if err != nil {
		log.Fatal(err)
	}

	return auth
}

// func GetTokenInstance() *contracts.LINKToken {
// 	return token
// }

// func GetCertNFTInstance() *contracts.CertNFT {
// 	return cert
// }

// func GetEduManageInstance() *contracts.EduManage {
// 	return manage
// }
