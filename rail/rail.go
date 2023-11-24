package rail

import (
	"crypto-wallet/rail/avalanche"
	"crypto-wallet/rail/binancechain"
	"crypto-wallet/rail/ethereum"
	"crypto-wallet/rail/models"
	"crypto-wallet/rail/polygon"
	"crypto-wallet/rail/telos"
	"errors"
)

const (
	Ethereum     = "ETHEREUM"
	Avalanche    = "AVALANCHE"
	Telos        = "TELOS"
	BinanceChain = "BINANCE_CHAIN"
	Polygon      = "POLYGON"
)

func GenerateWallet(rail string) (*models.Wallet, error) {
	switch rail {
	case Ethereum:
		return ethereum.GenerateWallet()
	case Avalanche:
		return avalanche.GenerateWallet()
	case Telos:
		return telos.GenerateWallet()
	case BinanceChain:
		return binancechain.GenerateWallet()
	case Polygon:
		return polygon.GenerateWallet()
	default:
		return nil, errors.New("rail not found")
	}
}
