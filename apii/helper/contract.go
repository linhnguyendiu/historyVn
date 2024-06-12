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
	client *ethclient.Client
	token  *contracts.LINKToken
	cert   *contracts.CertNFT
	manage *contracts.EduManage
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
	client = DialClient()
	var err error
	linkTokenAddress := common.HexToAddress("0x5FbDB2315678afecb367f032d93F642f64180aa3")
	token, err = contracts.NewLINKToken(linkTokenAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	log.Println("Connect to LINKToken successfully!!!")

	time.Sleep(250 * time.Millisecond)

	name, err := token.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Name", name)
}

func ConnectToCertNFT() {
	client = DialClient()
	var err error
	certNFTAddress := common.HexToAddress("0x5FC8d32690cc91D4c39d9d3abcBD16989F875707")
	token, err = contracts.NewLINKToken(certNFTAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a CertNFT contract: %v", err)
	}
	log.Println("Connect to CertNFT successfully!!!")

	time.Sleep(250 * time.Millisecond)
}

func ConnectToEduManage() {
	client = DialClient()
	var err error
	eduManageAddress := common.HexToAddress("0x0165878A594ca255338adfa4d48449f69242Eb8F")
	token, err = contracts.NewLINKToken(eduManageAddress, client)
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

func GetTokenInstance() *contracts.LINKToken {
	return token
}

func GetCertNFTInstance() *contracts.CertNFT {
	return cert
}

func GetEduManageInstance() *contracts.EduManage {
	return manage
}
