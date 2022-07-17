package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/thanhpp/scm/internal/nftsrv/domain/entity"
	"github.com/thanhpp/scm/internal/nftsrv/domain/repo"
	"github.com/thanhpp/scm/internal/nftsrv/infra/adapter/ipfsclient"
	"github.com/thanhpp/scm/internal/nftsrv/infra/adapter/nftminter"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/logger"
	"github.com/thanhpp/scm/pkg/smartcontracts"
)

type App struct {
	ipfs        *ipfsclient.IPFSClient
	minter      *nftminter.NFTMinter
	seriNFTRepo repo.SeriNFTRepo
}

func NewApp(ctx context.Context, ipfs *ipfsclient.IPFSClient, minter *nftminter.NFTMinter, seriNFTRepo repo.SeriNFTRepo) *App {
	a := &App{
		ipfs:        ipfs,
		minter:      minter,
		seriNFTRepo: seriNFTRepo,
	}

	go a.autoUpdateTokenID(ctx)

	return a
}

func (a *App) MintSeriNFT(ctx context.Context, seri string, metadata map[string]string) (*entity.SerialNFT, error) {
	if metadata == nil {
		return nil, errors.New("mint seri nft: empty metadata")
	}

	if len(seri) == 0 {
		return nil, errors.New("mint seri nft: empty seri")
	}

	ok, err := a.seriNFTRepo.CheckDuplicateSeri(ctx, seri)
	if err != nil {
		return nil, err
	}
	if ok {
		return nil, fmt.Errorf("mint seri nft: duplicate seri %s", seri)
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

	txInfo, err := a.minter.MintNFT(ctx, fmt.Sprintf("ipfs://%s", ipfsCid))
	if err != nil {
		return nil, err
	}

	newSeriNFT := &entity.SerialNFT{
		Seri:     seri,
		TxHash:   txInfo.TxHash,
		IPFSHash: ipfsCid,
		Metadata: string(tmpFileData),
	}

	if err := a.seriNFTRepo.Create(ctx, newSeriNFT); err != nil {
		return nil, err
	}

	return newSeriNFT, nil
}

func (a *App) GetSeriNFTBySeri(ctx context.Context, seri string) (*entity.SerialNFT, error) {
	return a.seriNFTRepo.GetBySeri(ctx, seri)
}

func (a *App) GetSeriNFTByTokenID(ctx context.Context, tokenID int64) (*entity.SerialNFT, error) {
	return a.seriNFTRepo.GetSeriNFTByTokenID(ctx, tokenID)
}

func (a *App) GetSeriNFTByTxHash(ctx context.Context, txHash string) (*entity.SerialNFT, error) {
	return a.seriNFTRepo.GetSeriNFTByTxHash(ctx, txHash)
}

func (a *App) autoUpdateTokenID(ctx context.Context) {
	updateTicker := time.NewTicker(constx.AutoUpdateTokenIDInterval)
	defer updateTicker.Stop()

	for ; true; <-updateTicker.C {
		if err := ctx.Err(); err != nil {
			logger.Errorw("auto update token id: context error", "error", err)
			return
		}

		seriNFTs, err := a.seriNFTRepo.GetSeriNFTWithEmptyTokenID(ctx)
		if err != nil {
			logger.Errorw("auto update token id: get seri nft with empty token id error", "error", err)
			continue
		}

		for _, seriNFT := range seriNFTs {
			tokenID, err := a.minter.GetTokenIDByTxHash(ctx, seriNFT.TxHash)
			if err != nil {
				if errors.Is(err, nftminter.ErrPendingTx) {
					logger.Warnw("auto update token id: pending tx",
						"seri", seriNFT.Seri, "txHash", seriNFT.TxHash, "error", err)
					continue
				}
				logger.Errorw("auto update token id: get token id by tx hash error",
					"seri", seriNFT.Seri, "txHash", seriNFT.TxHash, "error", err)
				continue
			}

			if err := a.seriNFTRepo.UpdateTokenIDByTxHash(ctx, seriNFT.TxHash, tokenID); err != nil {
				logger.Errorw("auto update token id: update token id by tx hash error",
					"seri", seriNFT.Seri, "txHash", seriNFT.TxHash, "error", err)
				continue
			}
		}
	}
}
