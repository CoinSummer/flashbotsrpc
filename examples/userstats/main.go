package main

import (
	"errors"
	"fmt"
	"github.com/metachris/flashbotsrpc/examples/signature"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

//var privateKey, _ = crypto.GenerateKey() // creating a new private key for testing. you probably want to use an existing key.
var privateKey, _ = crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

func main() {
	rpc := flashbotsrpc.New("https://relay.flashbots.net")

	blockNum := fmt.Sprintf("0x%x", 14747145)

	s, body, err := signature.Signature(rpc, privateKey, flashbotsrpc.FbGetuserstats, blockNum)
	if err != nil {
		_ = fmt.Errorf("call signature error: %s", err)
		return
	}

	// Query relay for user stats
	result, err := rpc.FlashbotsGetUserStats(crypto.PubkeyToAddress(privateKey.PublicKey), s, body)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) {
			fmt.Println(err.Error())
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		return
	}

	// Print result
	fmt.Printf("%+v\n", result)
}
