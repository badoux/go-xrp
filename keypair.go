package xrp

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/mr-tron/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
)

type KeyPair struct {
	MasterKey       *bip32.Key
	ExtendedPrivKey *bip32.Key
}

func NewKeyPairFromMnemonic(mnemonic, passphrase string) (*KeyPair, error) {
	seed := bip39.NewSeed(mnemonic, passphrase)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, err
	}
	keys := &KeyPair{
		MasterKey:       masterKey,
		ExtendedPrivKey: nil,
	}
	_, err = keys.Extend()
	if err != nil {
		return nil, err
	}
	return keys, nil
}

func (k *KeyPair) Extend() (*bip32.Key, error) {
	if k.ExtendedPrivKey != nil {
		return k.ExtendedPrivKey, nil
	}
	key, err := k.MasterKey.NewChildKey(HardenedKeyZeroIndex + 44) // purpose
	if err != nil {
		return nil, err
	}
	key, err = key.NewChildKey(HardenedKeyZeroIndex + 144) // coin
	if err != nil {
		return nil, err
	}
	key, err = key.NewChildKey(HardenedKeyZeroIndex)
	if err != nil {
		return nil, err
	}
	extKey, err := key.NewChildKey(0)
	if err != nil {
		return nil, err
	}

	k.ExtendedPrivKey = extKey
	return k.ExtendedPrivKey, nil
}

func (k *KeyPair) ExtendedPrivateKey() *bip32.Key {
	if k.ExtendedPrivKey == nil && k.MasterKey != nil {
		k.ExtendedPrivKey, _ = k.Extend()
	}
	return k.ExtendedPrivKey
}

func (k *KeyPair) Wallet(addrIdx uint32) (*Wallet, error) {
	if k.ExtendedPrivKey == nil {
		_, err := k.Extend()
		if err != nil {
			return nil, err
		}
	}
	adrKeyIndex, _ := k.ExtendedPrivKey.NewChildKey(addrIdx)

	secp256k1PrivKey, secp256k1PubKey := secp256k1.PrivKeyFromBytes(adrKeyIndex.Key)

	h256 := sha256.New()
	h256.Write(secp256k1PubKey.Serialize())

	md160 := ripemd160.New()
	md160.Write(h256.Sum(nil))
	// Account ID
	accountID := md160.Sum(nil)

	// Payload
	payload := []byte{0x00}
	payload = append(payload, accountID...)

	h256 = sha256.New()
	h256.Write(payload)
	tmp256 := h256.Sum(nil)

	h256 = sha256.New()
	h256.Write(tmp256)
	tmp256 = h256.Sum(nil)

	// Checksum
	checksum := tmp256[0:4]

	data := append(payload, checksum...)
	xrpAddress := base58.EncodeAlphabet(data, RippleAlphabet)

	return &Wallet{
		Address:    Account(xrpAddress),
		PublicKey:  hex.EncodeToString(secp256k1PubKey.Serialize()),
		PrivateKey: hex.EncodeToString(secp256k1PrivKey.Serialize()),
	}, nil
}
