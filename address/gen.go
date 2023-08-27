package main

import (
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/manifoldco/promptui"
	"gitlab.com/sepior/go-tsm-sdk/sdk/tsm"
)

func main() {
	b, err := os.ReadFile("../creds.json")
	if err != nil {
		log.Fatal(err)
	}
	credentials := string(b)

	// Create ECDSA client from credentials

	tsmClient, err := tsm.NewPasswordClientFromEncoding(3, 1, credentials)
	if err != nil {
		log.Fatal(err)
	}

	ecdsaClient := tsm.NewECDSAClient(tsmClient) // ECDSA with secp256k1 curve

	// Prompt for key id
	fmt.Println("Enter key id")
	keyIDPrompt := promptui.Prompt{
		Label: "Key ID",
	}

	keyID, err := keyIDPrompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	chainPath := []uint32{1, 2, 3, 4}
	derPublicKey, err := ecdsaClient.PublicKey(keyID, chainPath)
	if err != nil {
		panic(err)
	}

	publicKey, err := ecdsaClient.ParsePublicKey(derPublicKey)
	if err != nil {
		panic(err)
	}

	compressedPublicKey := make([]byte, 1+32)
	ySignFlag := byte(publicKey.Y.Bit(0))
	compressedPublicKey[0] = 2 | ySignFlag
	publicKey.X.FillBytes(compressedPublicKey[1:])

	address, err := btcutil.NewAddressPubKey(compressedPublicKey, &chaincfg.TestNet3Params)
	if err != nil {
		panic(err)
	}
	// Note: Encoding a *AddressPubKey (pay-to-pubkey) results in a P2PKH address
	//
	//	(pay-to-pubkey-hash). Convert address to a *AddressPubKeyHash before using it.
	btcAddress := address.EncodeAddress()
	fmt.Println("Bitcoin address: ", btcAddress)
}
