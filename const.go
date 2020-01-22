package xrp

import "github.com/mr-tron/base58"

type Account string
type Amount string
type TxType uint16

const HardenedKeyZeroIndex = uint32(0x80000000)

var (
	RippleAlphabet = base58.NewAlphabet("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")
)

const (
	TypePayment        TxType = 0
	TypeEscrowCreate   TxType = 1
	TypeEscrowFinish   TxType = 2
	TypeAccountSet     TxType = 3
	TypeEscrowCancel   TxType = 4
	TypeRegularKeySet  TxType = 5
	TypeNicknameSet    TxType = 6
	TypeOfferCreate    TxType = 7
	TypeOfferCancel    TxType = 8
	TypeTicketCreate   TxType = 10
	TypeTicketCancel   TxType = 11
	TypeSignerListSet  TxType = 12
	TypePaychanCreate  TxType = 13
	TypePaychanFund    TxType = 14
	TypePaychanClaim   TxType = 15
	TypeCheckCreate    TxType = 16
	TypeCheckCash      TxType = 17
	TypeCheckCancel    TxType = 18
	TypeDepositPreauth TxType = 19
	TypeTrustSet       TxType = 20
	TypeAccountDelete  TxType = 21
	TypeAmendment      TxType = 100
	TypeFee            TxType = 101
)
