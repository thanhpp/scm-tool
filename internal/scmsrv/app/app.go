package app

import (
	"context"
	"encoding/json"
	"time"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/internal/scmsrv/infra/adapter/nftsvclient"
	"github.com/thanhpp/scm/pkg/booting"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/fileutil"
	"github.com/thanhpp/scm/pkg/logger"
	"github.com/thanhpp/scm/pkg/rbmq"
	"github.com/thanhpp/scm/pkg/sharedto"
)

type App struct {
	ImportTicketHandler ImportTicketHandler
	SupplierHandler     SupplierHandler
	StorageHandler      StorageHandler
	ItemHandler         ItemHandler
	UserHandler         UserHandler
	rbmqClient          *rbmq.Client
}

func New(
	fac entity.Factory,
	itemRepo repo.ItemRepo, supplierRepo repo.SupplierRepo, storageRepo repo.StorageRepo,
	importTicketRepo repo.ImportTicketRepo, serialRepo repo.SerialRepo, userRepo repo.UserRepo,
	fileUtil fileutil.FileUtil, nftSrvClient *nftsvclient.NFTServiceClient,
) App {
	rbmqClient, err := rbmq.NewClient(constx.RabbitMQServerURL)
	if err != nil {
		panic(err)
	}
	if err := rbmqClient.CreateQueue(sharedto.MintNFTRequestQueue); err != nil {
		panic(err)
	}
	if err := rbmqClient.CreateQueue(sharedto.SeriNFTInfoQueue); err != nil {
		panic(err)
	}

	return App{
		ImportTicketHandler: ImportTicketHandler{
			itemRepo:         itemRepo,
			supplierRepo:     supplierRepo,
			storageRepo:      storageRepo,
			importTicketRepo: importTicketRepo,
			serialRepo:       serialRepo,
			fac:              fac,
			fileUtil:         fileUtil,
			nftServiceClient: nftSrvClient,
		},
		SupplierHandler: SupplierHandler{
			fac:          fac,
			supplierRepo: supplierRepo,
		},
		StorageHandler: StorageHandler{
			fac:         fac,
			storageRepo: storageRepo,
			itemRepo:    itemRepo,
		},
		ItemHandler: ItemHandler{
			fac:      fac,
			itemRepo: itemRepo,
			fileUtil: fileUtil,
		},
		UserHandler: UserHandler{
			f:        fac,
			userRepo: userRepo,
		},
		rbmqClient: rbmqClient,
	}
}

func (a *App) AutoFallbackUpdateSerial(
	serialRepo repo.SerialRepo, nftClient *nftsvclient.NFTServiceClient,
) booting.Daemon {
	return func(ctx context.Context) (start func() error, cleanup func()) {
		start = func() error {
			t := time.NewTicker(constx.AutoFallbackUpdateSerialInterval)
			defer t.Stop()

			for ; true; <-t.C {
				if err := ctx.Err(); err != nil {
					return nil
				}

				unmintedSerials, err := serialRepo.GetSeriWithEmptyTokenID(ctx)
				if err != nil {
					logger.Errorw("get unminted serials err", "err", err)
					continue
				}

				for i := range unmintedSerials {
					info, err := nftClient.GetNFTInfoBySeri(ctx, unmintedSerials[i].Seri)
					if err != nil {
						logger.Errorw("get tokenID err", "err", err)
					} else {
						if err := serialRepo.UpdateSerial(ctx, unmintedSerials[i].Seri,
							func(ctx context.Context, s *entity.Serial) (*entity.Serial, error) {
								s.TokenID = int(info.TokenID)
								return s, nil
							}); err != nil {
							logger.Errorw("update serial err", "err", err)
							continue
						}
					}
				}
			}
			return nil
		}
		return
	}
}

func (a *App) AutoMintNFT(serialRepo repo.SerialRepo) booting.Daemon {
	return func(ctx context.Context) (start func() error, cleanup func()) {
		start = func() error {
			t := time.NewTicker(constx.AutoMintNFTInterval)
			defer t.Stop()

			for ; true; <-t.C {
				if err := ctx.Err(); err != nil {
					return nil
				}

				unmintedSerials, err := serialRepo.GetSeriWithEmptyTokenID(ctx)
				if err != nil {
					logger.Errorw("get unminted serials err", "err", err)
					continue
				}

				req := new(sharedto.ReqCreateSeriNFT)
				for i := range unmintedSerials {
					if err := req.SetData(unmintedSerials[i]); err != nil {
						logger.Errorw("set mint nft data error", "err", err)
						continue
					}

					data, err := json.Marshal(req)
					if err != nil {
						logger.Errorw("marshal mint nft req err", "err", err)
						continue
					}

					if err := a.rbmqClient.PublishJSONMessage(sharedto.MintNFTRequestQueue, data); err != nil {
						logger.Errorw("publish mint nft error", "err", err)
						continue
					}

					logger.Infow("publish mint nft", "seri", unmintedSerials[i].Seri)
				}
			}

			cleanup = func() {
				logger.Infow("auto mint nft stopped")
			}

			return nil
		}

		return
	}
}

func (a *App) AutoUpdateSeriNFT(serialRepo repo.SerialRepo) booting.Daemon {
	return func(ctx context.Context) (start func() error, cleanup func()) {
		start = func() error {
			msgC, err := a.rbmqClient.GetConsumerChannel(sharedto.SeriNFTInfoQueue)
			if err != nil {
				return err
			}

			for {
				select {
				case <-ctx.Done():
					return nil
				case byteMsg := <-msgC:
					msg := new(sharedto.SeriNFTInfo)
					if err := json.Unmarshal(byteMsg.Body, msg); err != nil {
						logger.Errorw("unmarshal seri nft info", "err", err, "data", string(byteMsg.Body))
						continue
					}

					logger.Infow("received seri nft info update", "msg", msg)

					if err := serialRepo.UpdateSerial(ctx, msg.Seri,
						func(ctx context.Context, s *entity.Serial) (*entity.Serial, error) {
							s.TokenID = int(msg.TokenID)
							return s, nil
						}); err != nil {
						logger.Errorw("auto update serial error", "err", err, "data", msg)
						continue
					}

					logger.Infow("serial info updated", "seri", msg.Seri)
				}
			}
		}

		cleanup = func() {
			logger.Info("auto update seri nft stopped")
		}
		return
	}
}
