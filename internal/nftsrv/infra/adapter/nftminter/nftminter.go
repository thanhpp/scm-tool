package nftminter

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thanhpp/scm/pkg/smartcontracts"
)

type NFTMinter struct {
	ethClient    ethclient.Client
	privKey      *ecdsa.PrivateKey
	pubKey       ecdsa.PublicKey
	contractAddr common.Address
	sm           *smartcontracts.Smartcontracts
}

func NewNFTMinter(ethClient *ethclient.Client, contractAddr string) (*NFTMinter, error) {
	addrContractAddr := common.HexToAddress(contractAddr)
	instance, err := smartcontracts.NewSmartcontracts(
		addrContractAddr,
		ethClient,
	)
	if err != nil {
		return nil, err
	}

	return &NFTMinter{
		contractAddr: addrContractAddr,
		sm:           instance,
	}, nil
}
