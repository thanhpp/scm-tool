package smartcontractsv2_test

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/scm/pkg/smartcontractsv2"
)

const (
	contractAddr  = "0x2b01Ad90AeE0b10d20238dED96d297151b5C123d"
	toAddr        = "0x0000000000000000000000000000000000000dead"
	envFile       = ".env"
	envPrivateKey = "PRIVATE_KEY"
	envRestURL    = "API_URL"
)

func init() {
	if err := godotenv.Load(envFile); err != nil {
		log.Println("load env file", envFile, "err", err)
	}
}

func newInstance() (*ethclient.Client, *smartcontractsv2.Smartcontractsv2, error) {
	client, err := ethclient.Dial(os.Getenv(envRestURL))
	if err != nil {
		return nil, nil, err
	}

	address := common.HexToAddress(contractAddr)
	instance, err := smartcontractsv2.NewSmartcontractsv2(address, client)
	if err != nil {
		return nil, nil, err
	}

	return client, instance, nil
}

func newAuth(ctx context.Context, client *ethclient.Client) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(os.Getenv(envPrivateKey))
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("parse public key error")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		return nil, fmt.Errorf("get nonce at: %w", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("suggest gas price: %w", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 300000
	auth.GasPrice = gasPrice

	return auth, nil
}

func TestLoadSmartContract(t *testing.T) {
	client, err := ethclient.Dial(os.Getenv(envRestURL))
	require.NoError(t, err)

	address := common.HexToAddress(contractAddr)
	instance, err := smartcontractsv2.NewSmartcontractsv2(address, client)
	require.NoError(t, err)

	t.Log("loaded contract")
	t.Logf("%+v", instance)
}

func TestSelfMintNFT(t *testing.T) {
	client, ins, err := newInstance()
	require.NoError(t, err)

	auth, err := newAuth(context.Background(), client)
	require.NoError(t, err)

	tx, err := ins.SelfMintNFT(auth, "testTokenURI")
	require.NoError(t, err)

	t.Logf("tx hash: %s", tx.Hash())
}

func TestMultiSafeTransfer(t *testing.T) {
	client, ins, err := newInstance()
	require.NoError(t, err)

	auth, err := newAuth(context.Background(), client)
	require.NoError(t, err)

	tx, err := ins.MultiSafeTransferFrom(auth, auth.From, common.HexToAddress(toAddr), []*big.Int{
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(3),
	})
	require.NoError(t, err)

	t.Logf("tx hash: %s", tx.Hash().String())
	t.Logf("tx gas limit: %d", tx.Gas())
}
