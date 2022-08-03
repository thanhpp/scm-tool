package nftminter

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thanhpp/scm/pkg/logger"
	"github.com/thanhpp/scm/pkg/smartcontracts"
)

const (
	gasLimit = 300000
)

var (
	bigIntZero = big.NewInt(0)
)

var (
	ErrPendingTx = errors.New("pending tx")
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

func (m *NFTMinter) FromAddr() common.Address {
	return m.FromAddr()
}

func (m *NFTMinter) newAuth(ctx context.Context) (*bind.TransactOpts, error) {
	nonce, err := m.ethClient.NonceAt(ctx, m.fromAddr, nil)
	if err != nil {
		return nil, err
	}

	gasPrice, err := m.ethClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(m.privKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = bigIntZero
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	return auth, nil
}

type MintTxInfo struct {
	TxHash string
}

func (m *NFTMinter) MintNFT(ctx context.Context, tokenURI string) (*MintTxInfo, error) {
	auth, err := m.newAuth(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := m.sm.MintNFT(auth, m.fromAddr, tokenURI)
	if err != nil {
		return nil, err
	}

	txJson, err := tx.MarshalJSON()
	if err != nil {
		return nil, err
	}
	logger.Debugw("mint nft tx", "tx", string(txJson))

	return &MintTxInfo{
		TxHash: tx.Hash().String(),
	}, nil
}

func (m *NFTMinter) GetTokenIDByTxHash(ctx context.Context, txHash string) (uint64, error) {
	// get tx receipt
	hashTxHash := common.HexToHash(txHash)
	_, isPending, err := m.ethClient.TransactionByHash(ctx, hashTxHash)
	if err != nil {
		return 0, fmt.Errorf("get tx by hash error: %w", err)
	}
	if isPending {
		return 0, ErrPendingTx
	}

	receipt, err := m.ethClient.TransactionReceipt(ctx, common.HexToHash(txHash))
	if err != nil {
		return 0, fmt.Errorf("get tx receipt error: %w", err)
	}

	// check logs
	for i := range receipt.Logs {
		if len(receipt.Logs[i].Topics) != 4 {
			logger.Warnw("invalid topics", "txHash", txHash, "topics", receipt.Logs[i].Topics)
			continue
		}

		// parse to get tokenID
		hexTokenID := receipt.Logs[i].Topics[3].Hex()

		if len(hexTokenID) >= 3 {
			for i := 2; i < len(hexTokenID); i++ {
				if hexTokenID[i] != '0' {
					hexTokenID = "0x" + hexTokenID[i:]

					uint64TokenID, err := hexutil.DecodeUint64(hexTokenID)
					if err != nil {
						return 0, fmt.Errorf("parse tokenID error: %v", err)
					}

					return uint64TokenID, nil
				}
			}
		}
	}

	return 0, errors.New("tokenID not found in tx hash " + txHash)
}

func (m *NFTMinter) Transfer(ctx context.Context, tokenID int, toAddr common.Address) (string, error) {
	auth, err := m.newAuth(ctx)
	if err != nil {
		return "", err
	}

	tx, err := m.sm.SafeTransferFrom(auth, m.fromAddr, toAddr, big.NewInt(int64(tokenID)))
	if err != nil {
		return "", err
	}
	logger.Infow("Safe Tranfer", "to", toAddr.String(), "tokenID", tokenID, "tx hash", tx.Hash())

	return tx.Hash().String(), nil
}
