package nftminter

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thanhpp/scm/pkg/smartcontracts"
)

const (
	gasLimit = 300000
)

var (
	bigIntZero = big.NewInt(0)
)

type NFTMinter struct {
	ethClient    *ethclient.Client
	privKey      *ecdsa.PrivateKey
	pubKey       *ecdsa.PublicKey
	sm           *smartcontracts.Smartcontracts
	contractAddr common.Address
	fromAddr     common.Address
}

func NewNFTMinter(ethClient *ethclient.Client, privateKey, contractAddr string) (*NFTMinter, error) {
	addrContractAddr := common.HexToAddress(contractAddr)
	instance, err := smartcontracts.NewSmartcontracts(
		addrContractAddr,
		ethClient,
	)
	if err != nil {
		return nil, err
	}

	ecdsaPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	publicKey := ecdsaPrivateKey.Public()
	ecsdaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("parse ecdsa public key error")
	}

	return &NFTMinter{
		ethClient:    ethClient,
		privKey:      ecdsaPrivateKey,
		pubKey:       ecsdaPublicKey,
		sm:           instance,
		contractAddr: addrContractAddr,
		fromAddr:     crypto.PubkeyToAddress(*ecsdaPublicKey),
	}, nil
}

func (m *NFTMinter) newAuth(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := m.ethClient.PendingBalanceAt(ctx, m.fromAddr)
	if err != nil {
		return nil, err
	}

	gasPrice, err := m.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(m.privKey)
	auth.Nonce = nonce
	auth.Value = bigIntZero
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

func (m *NFTMinter) MintNFT(ctx context.Context, tokenURI string) (*struct{}, error) {
	auth, err := m.newAuth(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := m.sm.MintNFT(auth, m.fromAddr, tokenURI)
	if err != nil {
		return nil, err
	}

	fmt.Printf("tx: %v\n", tx)

	return nil, nil
}
