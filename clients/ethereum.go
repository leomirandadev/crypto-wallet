package clients

import (
	"context"
	"crypto-wallet/models"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewEthereum(url string) ethereumClient {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	return ethereumClient{client}
}

type ethereumClient struct {
	client *ethclient.Client
}

func (c ethereumClient) Sign(ctx context.Context, tr models.Transaction, privateKeyStr string, nonce uint64, contract []byte) (*types.Transaction, error) {
	chainID, err := c.client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, err
	}

	tx := TransactionToEthereum(tr, nonce, contract)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nil, err
	}
	return signedTx, nil
}

func TransactionToEthereum(tr models.Transaction, nonce uint64, contract []byte) *types.Transaction {
	return types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: tr.GasPrice,
		Gas:      tr.GasLimit,
		To:       &tr.ToAddress,
		Value:    tr.Value,
		Data:     contract,
	})
}
