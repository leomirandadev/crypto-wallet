package main

import (
	"crypto-wallet/rail"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	wallet, err := rail.GenerateWallet(rail.Ethereum)
	if err != nil {
		log.Fatal(err)
	}

	printJSON("Wallet successfuly generated:", wallet)
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
