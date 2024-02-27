package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/cobra"
)

var ethwallet = &cobra.Command{
	Use:   "ethwallet",
	Short: "get ethereum wallet form private key",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		privateKey, err := crypto.HexToECDSA(key)
		if err != nil {
			log.Fatalf("failed to parse private key: %v", err)
		}
		importPrivateKeyExample(privateKey)
	},
}

func importPrivateKeyExample(privateKey *ecdsa.PrivateKey) {

	ks := keystore.NewKeyStore(".", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.ImportECDSA(privateKey, "password")
	if err != nil {
		log.Fatal(err)
	}
	address := account.Address.Hex()

	fmt.Println("Imported Address:", address)
}

func init() {
	ethwallet.Flags().StringP("key", "k", "", "private key")
	ethereum.AddCommand(ethwallet)
}
