package polygon

import (
	"crypto-wallet/rail/ethereum"
	"crypto-wallet/rail/models"
)

// Use the same algorithm of ethereum
func GenerateWallet() (*models.Wallet, error) {
	return ethereum.GenerateWallet()
}
