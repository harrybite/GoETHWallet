package cmd

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var generate = &cobra.Command{
	Use:   "generate",
	Short: "generate ethereum base wallet address",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		generateEthWallet()
	},
}

func generateEthWallet() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Convert the private key to a hexadecimal string
	privateKeyHex := hex.EncodeToString(privateKey.D.Bytes())

	// Get the corresponding public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// Get the wallet address from the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println("Private key:", "0x"+privateKeyHex)
	fmt.Println("Address:", address)
}

func init() {
	ethereum.AddCommand(generate)
}
