package smartcontracts_test

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/thanhpp/scm/pkg/smartcontracts"
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
	auth.GasLimit = 300000
	auth.GasPrice = gasPrice

	instance, err := smartcontracts.NewSmartcontracts(common.HexToAddress(contractAddr), client)
	require.NoError(t, err)

	tx, err := instance.MintNFT(auth, fromAddr, tokenURI)
	require.NoError(t, err)

	txJson, err := tx.MarshalJSON()
	require.NoError(t, err)
	t.Logf("new mint nft tx (%s) %+v\n %s \n", tx.Hash().Hex(), tx, string(txJson))

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	for ; true; <-ticker.C {
		t.Log("querying")
		_, isPending, err := client.TransactionByHash(ctx, tx.Hash())
		if err == nil && !isPending {
			break
		}
	}

	receipt, err := client.TransactionReceipt(ctx, tx.Hash())
	require.NoError(t, err, tx.Hash().String())

	t.Logf("tx receipt %+v\n", receipt)

	for i := range receipt.Logs {
		t.Logf("receipt logs %d %+v", i, receipt.Logs[i])

		require.Len(t, receipt.Logs[i].Topics, 4)

		hexTokenID := receipt.Logs[i].Topics[3].Hex()

		if len(hexTokenID) >= 3 {
			for i := 2; i < len(hexTokenID); i++ {
				if hexTokenID[i] != '0' {
					hexTokenID = "0x" + hexTokenID[i:]
					break
				}
			}
		}

		tokenID, err := hexutil.DecodeUint64(hexTokenID)
		require.NoError(t, err, hexTokenID)

		t.Logf("tokenID %d\n", tokenID)

	}
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
	auth.GasLimit = 300000
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

	transerTx, err := instance.SafeTransferFrom(auth, fromAddr, common.HexToAddress(toAddr), big.NewInt(4))
	require.NoError(t, err)
	t.Logf("new transfer tx (%s) %+v", transerTx.Hash().Hex(), transerTx)
}
