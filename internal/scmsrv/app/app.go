package app

import (
	"context"
	"time"

	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
	"github.com/thanhpp/scm/internal/scmsrv/infra/adapter/nftsvclient"
	"github.com/thanhpp/scm/pkg/booting"
	"github.com/thanhpp/scm/pkg/constx"
	"github.com/thanhpp/scm/pkg/fileutil"
	"github.com/thanhpp/scm/pkg/logger"
)

type App struct {
	ImportTicketHandler ImportTicketHandler
	SupplierHandler     SupplierHandler
	StorageHandler      StorageHandler
	ItemHandler         ItemHandler
	UserHandler         UserHandler
}

func New(
	fac entity.Factory,
	itemRepo repo.ItemRepo, supplierRepo repo.SupplierRepo, storageRepo repo.StorageRepo,
	importTicketRepo repo.ImportTicketRepo, serialRepo repo.SerialRepo, userRepo repo.UserRepo,
	fileUtil fileutil.FileUtil, nftSrvClient *nftsvclient.NFTServiceClient,
) App {
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
	}
}

func (a *App) AutoMintAndUpdateSerial(
	serialRepo repo.SerialRepo, nftClient *nftsvclient.NFTServiceClient,
) booting.Daemon {
	return func(ctx context.Context) (start func() error, cleanup func()) {
		start = func() error {
			t := time.NewTicker(constx.AutoMintAndUpdateSerialInterval)
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
					if err := nftClient.MintSeriNFT(ctx, unmintedSerials[i]); err != nil {
						logger.Errorw("mint err", "err", err)
					}
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
