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
	opts := flashbotsrpc.FlashbotsCallBundleParam{
		Txs:              []string{"0xf8643a85122370eda482520894287e21b9201e98ef3e2e0e8759ee36ca8257a6d2808026a01a6bc020e18258db911d11f7d93e8efb365470bec4734ed2e95011565d8d8da4a0570e81a2b179d1de691cf81549a0c5947762ac63ddb835017bf67010c1ae81e8"},
		BlockNumber:      fmt.Sprintf("0x%x", 14747150),
		StateBlockNumber: "latest",
	}

	s, body, err := signature.Signature(rpc, privateKey, flashbotsrpc.EthCallbundle, opts)
	if err != nil {
		_ = fmt.Errorf("call signature error: %s", err)
		return
	}
	result, err := rpc.FlashbotsCallBundle(crypto.PubkeyToAddress(privateKey.PublicKey), s, body)
	if err != nil {
		if errors.Is(err, flashbotsrpc.ErrRelayErrorResponse) { // user/tx error, rather than JSON or network error
			fmt.Println(err.Error())
		} else {
			fmt.Printf("error: %+v\n", err)
		}
		return
	}

	// Print result
	fmt.Printf("%+v\n", result)
}
