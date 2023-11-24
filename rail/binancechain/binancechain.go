package binancechain

import (
	"crypto-wallet/rail/models"
	"crypto/rand"
	"math/big"

	"github.com/binance-chain/go-sdk/keys"
)

func GenerateWallet() (*models.Wallet, error) {
	privKey, err := GenerateRandomString(64)
	if err != nil {
		return nil, err
	}

	keyManager, err := keys.NewPrivateKeyManager(privKey)
	if err != nil {
		return nil, err
	}

	return &models.Wallet{
		PrivateKey:    privKey,
		PublicKey:     string(keyManager.GetPrivKey().PubKey().Bytes()),
		PublicAddress: keyManager.GetAddr().String(),
	}, nil

}

func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789abcdef"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
