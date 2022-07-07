package smartcontracts_test

import (
	"context"
	"crypto/ecdsa"
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
	"github.com/thanhpp/scm/smartcontracts"
)

const (
	contractAddr  = "0xEA92d9D15139c316944aeba2aBe828B87cBA1265"
	toAddr        = "0x0000000000000000000000000000000000000001"
	envFile       = ".env"
	envPrivateKey = "PRIVATE_KEY"
	envRestURL    = "API_URL"
)

var (
	restURL = "https://rinkeby.infura.io"
)

func init() {
	if err := godotenv.Load(envFile); err != nil {
		log.Println("load env file", envFile, "err", err)
	}
}

func TestLoadSmartContract(t *testing.T) {
	client, err := ethclient.Dial(restURL)
	require.NoError(t, err)

	address := common.HexToAddress(contractAddr)
	instance, err := smartcontracts.NewSmartcontracts(address, client)
	require.NoError(t, err)

	t.Log("loaded contract")
	t.Logf("%+v", instance)
}

func TestMintNFT(t *testing.T) {
	var (
		ctx      = context.Background()
		tokenURI = "sampleTokenURI"
	)

	client, err := ethclient.Dial(os.Getenv(envRestURL))
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA(os.Getenv(envPrivateKey))
	require.NoError(t, err)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	require.NoError(t, err)
	t.Logf("current pending nonce %d", nonce)

	gasPrice, err := client.SuggestGasPrice(ctx)
	require.NoError(t, err)
	t.Logf("current gas price %s", gasPrice.String())

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 300000 // 0.1
	auth.GasPrice = gasPrice

	instance, err := smartcontracts.NewSmartcontracts(common.HexToAddress(contractAddr), client)
	require.NoError(t, err)

	tx, err := instance.MintNFT(auth, fromAddr, tokenURI)
	require.NoError(t, err)
	t.Logf("new mint nft tx (%s) %+v", tx.Hash().Hex(), tx)
}

func TestSafeTransfer(t *testing.T) {
	// mini new NFT
	var (
		ctx      = context.Background()
		tokenURI = "sampleTokenURI"
	)

	client, err := ethclient.Dial(os.Getenv(envRestURL))
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA(os.Getenv(envPrivateKey))
	require.NoError(t, err)

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	require.NoError(t, err)
	t.Logf("current pending nonce %d", nonce)

	gasPrice, err := client.SuggestGasPrice(ctx)
	require.NoError(t, err)
	t.Logf("current gas price %s", gasPrice.String())

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = 300000 // 0.1
	auth.GasPrice = gasPrice

	instance, err := smartcontracts.NewSmartcontracts(common.HexToAddress(contractAddr), client)
	require.NoError(t, err)

	mintTX, err := instance.MintNFT(auth, fromAddr, tokenURI)
	require.NoError(t, err)
	t.Logf("new mint nft tx (%s) %+v", mintTX.Hash().Hex(), mintTX)

	// transfer
	newNonce, err := client.PendingNonceAt(ctx, fromAddr)
	require.NoError(t, err)

	auth.Nonce = big.NewInt(int64(newNonce))

	// for now, the nonce == tokenID
	transerTx, err := instance.SafeTransferFrom(auth, fromAddr, common.HexToAddress(toAddr), big.NewInt(4))
	require.NoError(t, err)
	t.Logf("new transfer tx (%s) %+v", transerTx.Hash().Hex(), transerTx)
}
