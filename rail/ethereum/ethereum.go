package ethereum

import (
	"crypto-wallet/rail/models"
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// https://goethereumbook.org/wallet-generate/
// https://www.quicknode.com/guides/ethereum-development/wallets/how-to-generate-a-new-ethereum-address-in-go
func GenerateWallet() (*models.Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	wallet := new(models.Wallet)
	// Never share it
	hexPrivateKey := hexutil.Encode(crypto.FromECDSA(privateKey))
	wallet.PrivateKey = hexPrivateKey[2:] // removing 0x

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// parse public key
	wallet.PublicKey = hexutil.Encode(publicKeyBytes)[4:] // removing 0x04

	// parse public address
	wallet.PublicAddress = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return wallet, nil
}
