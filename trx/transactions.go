package trx

import (
	"go-xrp"
)

type Transaction struct {
	Account            xrp.Account
	TransactionType    xrp.TxType
	Fee                xrp.Amount
	Sequence           uint32
	AccountTxnID       string
	LastLedgerSequence uint32
	Flags              uint32
	Memos              []Memo
	Signers            []Signer
	SourceTag          uint32
	SigningPubKey      string
	TxnSignature       string
}

type PaymentTransaction struct {
	Transaction
	Amount         Amount
	Destination    Account
	DestinationTag uint32
	InvoiceID      string
	Paths          []Path
	SendMax        Amount
	DeliverMin     Amount
}

type Path struct {
	Account  Account `json:"account,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Issuer   Account `json:"issuer,omitempty"`
}

type Signer struct {
	Account       Account
	TxnSignature  string
	SigningPubKey string
}

type Memo struct {
	MemoData   string
	MemoFormat string
	MemoType   string
}
