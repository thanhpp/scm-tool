package storage

import (
	"context"

	"github.com/thanhpp/scm/internal/nftsrv/domain/entity"
	"github.com/thanhpp/scm/internal/nftsrv/domain/repo"
	"gorm.io/gorm"
)

type SeriNFTDB struct {
	db *gorm.DB
}

func (d *SeriNFTDB) CheckDuplicateSeri(ctx context.Context, seri string) (bool, error) {
	var seriNFTDB = new(repo.SeriNFT)

	if err := d.db.WithContext(ctx).Where("seri = ?", seri).Find(seriNFTDB).Error; err != nil {
		return false, err
	}

	if seriNFTDB.Seri == "" {
		return false, nil
	}

	return true, nil
}

func (d *SeriNFTDB) GetBySeri(ctx context.Context, seri string) (*entity.SerialNFT, error) {
	seriNFTDB := new(repo.SeriNFT)

	if err := d.db.WithContext(ctx).Where("seri = ?", seri).First(seriNFTDB).Error; err != nil {
		return nil, err
	}

	return unmarshalSeriNFT(seriNFTDB), nil
}

func (d *SeriNFTDB) GetSeriNFTWithEmptyTokenID(ctx context.Context) ([]*entity.SerialNFT, error) {
	var seriNFTsDB []*repo.SeriNFT

	if err := d.db.WithContext(ctx).Where("token_id = ?", 0).Find(&seriNFTsDB).Error; err != nil {
		return nil, err
	}

	seriNFTs := make([]*entity.SerialNFT, len(seriNFTsDB))
	for i, s := range seriNFTsDB {
		seriNFTs[i] = unmarshalSeriNFT(s)
	}

	return seriNFTs, nil
}

func (d *SeriNFTDB) GetWaitingTransferSerialNFT(
	ctx context.Context, excludeOwner string,
) ([]*entity.SerialNFT, error) {
	var seriNFTsDB []*repo.SeriNFT

	if err := d.db.WithContext(ctx).Model(&repo.SeriNFT{}).
		Where("(owner <> ? AND owner IS NOT NULL) AND (transfer_tx_hash = '')",
			excludeOwner).Find(&seriNFTsDB).Error; err != nil {
		return nil, err
	}

	seriNFTs := make([]*entity.SerialNFT, len(seriNFTsDB))
	for i := range seriNFTsDB {
		seriNFTs[i] = unmarshalSeriNFT(seriNFTsDB[i])
	}

	return seriNFTs, nil
}

func (d *SeriNFTDB) GetSeriNFTByTokenID(ctx context.Context, tokenID int64) (*entity.SerialNFT, error) {
	seriNFTDB := new(repo.SeriNFT)

	if err := d.db.WithContext(ctx).Where("token_id = ?", tokenID).First(seriNFTDB).Error; err != nil {
		return nil, err
	}

	return unmarshalSeriNFT(seriNFTDB), nil
}

func (d *SeriNFTDB) GetSeriNFTByTxHash(ctx context.Context, txHash string) (*entity.SerialNFT, error) {
	seriNFTDB := new(repo.SeriNFT)

	if err := d.db.WithContext(ctx).Where("tx_hash LIKE ?", txHash).First(seriNFTDB).Error; err != nil {
		return nil, err
	}

	return unmarshalSeriNFT(seriNFTDB), nil
}

func (d *SeriNFTDB) Create(ctx context.Context, seriNFT *entity.SerialNFT) error {
	newSeriNFTDB := marshalSeriNFT(seriNFT)

	if err := d.db.WithContext(ctx).Create(newSeriNFTDB).Error; err != nil {
		return err
	}

	return nil
}

func (d *SeriNFTDB) UpdateTokenIDByTxHash(ctx context.Context, txHash string, tokenID uint64) error {
	if err := d.db.WithContext(ctx).Model(&repo.SeriNFT{}).
		Where("tx_hash = ?", txHash).Update("token_id", tokenID).
		Error; err != nil {
		return err
	}

	return nil
}

func (d *SeriNFTDB) UpdateSeriNFTBySeri(
	ctx context.Context, seri string, fn repo.UpdateSeriNFTFunc,
) error {
	return d.db.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			seriNFTDB, err := txGetSeriNFTBySeri(ctx, tx, seri)
			if err != nil {
				return err
			}

			seriNFT := unmarshalSeriNFT(seriNFTDB)

			newSeri, err := fn(seriNFT)
			if err != nil {
				return err
			}

			newSeriDB := marshalSeriNFT(newSeri)

			if err := tx.
				WithContext(ctx).Model(&repo.SeriNFT{}).
				Where("seri LIKE ?", seri).Updates(newSeriDB).Error; err != nil {
				return err
			}

			return nil
		},
	)

}

func txGetSeriNFTBySeri(ctx context.Context, tx *gorm.DB, seri string) (*repo.SeriNFT, error) {
	seriNFTDB := new(repo.SeriNFT)

	err := tx.WithContext(ctx).Model(&repo.SeriNFT{}).Where("seri LIKE ?", seri).Take(&seriNFTDB).Error
	if err != nil {
		return nil, err
	}

	return seriNFTDB, nil
}

func marshalSeriNFT(seriNFT *entity.SerialNFT) *repo.SeriNFT {
	return &repo.SeriNFT{
		Seri:           seriNFT.Seri,
		TxHash:         seriNFT.TxHash,
		IPFSHash:       seriNFT.IPFSHash,
		Metadata:       seriNFT.Metadata,
		TokenID:        seriNFT.TokenID,
		Owner:          seriNFT.Owner,
		TransferTxHash: seriNFT.TransferTxHash,
	}
}

func unmarshalSeriNFT(seriNFT *repo.SeriNFT) *entity.SerialNFT {
	return &entity.SerialNFT{
		Seri:           seriNFT.Seri,
		TxHash:         seriNFT.TxHash,
		IPFSHash:       seriNFT.IPFSHash,
		Metadata:       seriNFT.Metadata,
		TokenID:        seriNFT.TokenID,
		Owner:          seriNFT.Owner,
		TransferTxHash: seriNFT.TransferTxHash,
	}
}
