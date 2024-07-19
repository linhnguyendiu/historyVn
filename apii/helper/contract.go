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
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/LLnypZlAP6s6LsS8ugbapx3S9soRziXM")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("we have a connection")

	return client
}

func ConnectToLINKToken() {
	Client = DialClient()
	var err error
	linkTokenAddress := common.HexToAddress("0x33Fe0D80e17bF3aB170Ba57D53Af63fefD316917")
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
	certNFTAddress := common.HexToAddress("0x0A79C71AECcdf0564Ce7b6cFcd8A413C1119575b")
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
	eduManageAddress := common.HexToAddress("0x25eaa81E1a3da566e30f51c3e9b6cbC1c0667df2")
	Manage, err = contracts.NewEduManage(eduManageAddress, Client)
	if err != nil {
		log.Fatalf("Failed to instantiate a EduManage contract: %v", err)
	}
	log.Println("Connect to EduManageAddress successfully!!!")

	time.Sleep(250 * time.Millisecond)
}

func AuthGenerator(client *ethclient.Client) *bind.TransactOpts {
	privateKeyHex := "0xd056eee1f82a146479b40e56d416caed0a1b926103295c081e3ba641bb387e18"
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:])
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // Hardhat uses chain ID 31337
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
