package main

import (
	"context"
	"crypto-wallet/clients"
	"crypto-wallet/models"
	"crypto-wallet/rail"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	receiverWallet = "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
	ethereumURL    = "https://rinkeby.infura.io"
)

func main() {
	wallet, err := rail.GenerateWallet(rail.Ethereum)
	if err != nil {
		log.Fatal("generate wallet fails", err)
	}
	printJSON("Wallet successfuly generated:", wallet)

	var (
		eclient         = clients.NewEthereum(ethereumURL)
		ctx             = context.Background()
		nonce    uint64 = 0
		contract        = []byte("")
	)

	tr := models.Transaction{
		GasPrice:  big.NewInt(20000000000),
		GasLimit:  uint64(21000),
		ToAddress: common.HexToAddress(receiverWallet),
		Value:     big.NewInt(1000000000000000000),
	}

	signTransaction, err := eclient.Sign(ctx, tr, wallet.PrivateKey, nonce, contract)
	if err != nil {
		log.Fatal("sign transaction fails", err)
	}
	printJSON("Sign Transaction:", signTransaction)
}

func printJSON(msg any, data any) {
	fmt.Println(msg)

	empJSON, err := json.MarshalIndent(data, ``, `  `)
	if err != nil {
		return
	}
	fmt.Println(string(empJSON))
}

/*
WARNING MACOS ON BNB

# github.com/zondax/hid
In file included from ../../../../go/pkg/mod/github.com/zondax/hid@v0.9.0/hid_enabled.go:38:
../../../../go/pkg/mod/github.com/zondax/hid@v0.9.0/hidapi/mac/hid.c:693:34: warning: 'kIOMasterPortDefault' is deprecated: first deprecated in macOS 12.0 [-Wdeprecated-declarations]
/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX.sdk/System/Library/Frameworks/IOKit.framework/Headers/IOKitLib.h:133:19: note: 'kIOMasterPortDefault' has been explicitly marked deprecated here
*/
