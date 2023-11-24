package arbitrum

import (
	"crypto-wallet/rail/ethereum"
	"crypto-wallet/rail/models"
)

// Use the same way of ethereum https://arbitrum.io/
func GenerateWallet() (*models.Wallet, error) {
	return ethereum.GenerateWallet()
}
