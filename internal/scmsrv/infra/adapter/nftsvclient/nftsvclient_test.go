package nftsvclient_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/infra/adapter/nftsvclient"
	"github.com/thanhpp/scm/pkg/logger"
)

func newNFTServiceClient() *nftsvclient.NFTServiceClient {
	const baseURL = "http://127.0.0.1:11000"
	c := nftsvclient.New(baseURL)

	return c
}

func newTestSeri() *entity.Serial {
	return &entity.Serial{
		Seri: "test-seri",
		Item: &entity.Item{
			Name: "test-item-name",
			Desc: "test-item-desc",
		},
		ImportTicket: &entity.ImportTicket{
			ID: 1,
			FromSupplier: entity.Supplier{
				Name: "test-supplier-name",
			},
		},
	}
}

func TestMintNFT(t *testing.T) {
	var (
		testSeri = newTestSeri()
		ctx      = context.Background()
	)
	logger.SetDefaultLog()

	c := newNFTServiceClient()

	err := c.MintSeriNFT(ctx, testSeri)
	require.NoError(t, err)
}

func TestGetTokenID(t *testing.T) {
	var (
		testSeri = newTestSeri()
		ctx      = context.Background()
	)
	logger.SetDefaultLog()

	c := newNFTServiceClient()

	info, err := c.GetNFTInfoBySeri(ctx, testSeri.Seri)
	require.NoError(t, err)

	t.Logf("token id info %+v\n", info)
}
