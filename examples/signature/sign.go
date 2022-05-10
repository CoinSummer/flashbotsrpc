package signature

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/metachris/flashbotsrpc"
)

func Signature(rpc *flashbotsrpc.FlashbotsRPC, privKey *ecdsa.PrivateKey, method flashbotsrpc.FlashBotMethod,
	params interface{}) (string, []byte, error) {

	body, err := rpc.FlashbotsMessage(method, params)
	if err != nil {
		return "", nil, err
	}

	hashedBody := crypto.Keccak256Hash(body).Hex()
	sig, err := crypto.Sign(accounts.TextHash([]byte(hashedBody)), privKey)
	if err != nil {
		return "", nil, err
	}
	return hexutil.Encode(sig), body, nil
}
