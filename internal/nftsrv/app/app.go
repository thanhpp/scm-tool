package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/thanhpp/scm/internal/nftsrv/domain/entity"
	"github.com/thanhpp/scm/internal/nftsrv/infra/adapter/ipfsclient"
	"github.com/thanhpp/scm/pkg/logger"
	"github.com/thanhpp/scm/pkg/smartcontracts"
)

type App struct {
	ipfs *ipfsclient.IPFSClient
}

func (a *App) MintSeriNFT(ctx context.Context, seri string, metadata map[string]string) (*entity.SerialNFT, error) {
	if metadata == nil {
		return nil, errors.New("mint seri nft: empty metadata")
	}

	nftData := smartcontracts.NFT{
		Name: seri,
	}

	for k, v := range metadata {
		nftData.Attributes = append(nftData.Attributes,
			smartcontracts.Attribute{
				TraitType: k,
				Value:     v,
			})
	}

	tmpFileData, err := json.Marshal(nftData)
	if err != nil {
		return nil, fmt.Errorf("marshal nft data error: %w", err)
	}

	f, err := os.Create(fmt.Sprintf("%s.json", seri))
	if err != nil {
		return nil, fmt.Errorf("create nft data file error: %w", err)
	}

	if _, err := f.Write(tmpFileData); err != nil {
		return nil, fmt.Errorf("write nft data error: %w", err)
	}

	ipfsCid, err := a.ipfs.UploadFile(ctx, f.Name())
	if err != nil {
		return nil, err
	}
	logger.Infow("uploaded file to ipfs", "file", f.Name(), "cid", ipfsCid)

	return nil, nil
}
