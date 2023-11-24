package models

type Wallet struct {
	/*
		This is the private key which is used for signing transactions
		and is to be treated like a password and never be shared, since
		who ever is in possesion of it will have access to all your funds.
	*/
	PrivateKey    string
	PublicKey     string
	PublicAddress string
}
