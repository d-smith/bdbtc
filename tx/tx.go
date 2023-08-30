package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
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

	keyId := "uZQ4rHjR0HeM8wUbHbDNqByZ1sT4"

	// Previous txn id - from Explora
	prevTxID := "f517a44469936d300c6c14ec92f7eab65a689d0ab0c01dbf0403e806abfed3f1"

	// Destination address
	destination := "moPFejHJc4LZ1hgNLLKxYAWCgCe64qQRP3"
	destinationAddr, err := btcutil.DecodeAddress(destination, &chaincfg.TestNet3Params)
	if err != nil {
		log.Fatal(err)
	}

	destinationAddrByte, err := txscript.PayToAddrScript(destinationAddr)
	if err != nil {
		log.Fatal(err)
	}

	utxoHash, err := chainhash.NewHashFromStr(prevTxID)
	if err != nil {
		log.Fatal(err)
	}

	// Spending utxo
	outPoint := wire.NewOutPoint(utxoHash, 0)

	redeemTx := wire.NewMsgTx(wire.TxVersion)

	txIn := wire.NewTxIn(outPoint, nil, nil)
	redeemTx.AddTxIn(txIn)

	redeemTxOut := wire.NewTxOut(100000000-2000, destinationAddrByte)
	redeemTx.AddTxOut(redeemTxOut)

	prevScript := "76a914cf8c27c50a962ce9d5750121cbef451d0b06112a88ac"
	prevScriptByte, err := hex.DecodeString(prevScript)

	// RawTxInSignature
	hash, err := txscript.CalcSignatureHash(prevScriptByte, txscript.SigHashAll, redeemTx, 0)
	if err != nil {
		log.Fatal(err)
	}
	signature, _, err := ecdsaClient.Sign(keyId, []uint32{0, 0}, hash)
	if err != nil {
		log.Fatal(err)
	}

	sigIn := append(signature, byte(txscript.SigHashAll))

	derPublicKey, err := ecdsaClient.PublicKey(keyId, []uint32{0, 0})
	if err != nil {
		log.Fatal(err)
	}

	publicKey, err := ecdsaClient.ParsePublicKey(derPublicKey)
	if err != nil {
		log.Fatal(err)
	}

	compressedPublicKey := make([]byte, 1+32)
	ySignFlag := byte(publicKey.Y.Bit(0))
	compressedPublicKey[0] = 2 | ySignFlag
	publicKey.X.FillBytes(compressedPublicKey[1:])

	sigScript, err := txscript.NewScriptBuilder().AddData(sigIn).AddData(compressedPublicKey).Script()
	if err != nil {
		log.Fatal(err)
	}

	redeemTx.TxIn[0].SignatureScript = sigScript
	var signedTx bytes.Buffer
	redeemTx.Serialize(&signedTx)

	hexSignedTx := hex.EncodeToString(signedTx.Bytes())

	fmt.Printf("Signed Tx: %s\n", hexSignedTx)

}
