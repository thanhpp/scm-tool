package repo

import (
	"context"
	"time"

	"github.com/thanhpp/scm/internal/nftsrv/domain/entity"
)

type SeriNFT struct {
	Seri      string `gorm:"column:seri;type:text;primaryKey"`
	TxHash    string `gorm:"column:tx_hash;type:text"`
	IPFSHash  string `gorm:"column:ipfs_hash;type:text"`
	Metadata  string `gorm:"column:metadata;type:text"`
	TokenID   int64  `gorm:"column:token_id;type:bigint"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateSeriNFTFunc func(*entity.SerialNFT) (*entity.SerialNFT, error)

type SeriNFTRepo interface {
	CheckDuplicateSeri(ctx context.Context, seri string) (bool, error)
	GetBySeri(ctx context.Context, seri string) (*entity.SerialNFT, error)
	GetSeriNFTWithEmptyTokenID(ctx context.Context) ([]*entity.SerialNFT, error)
	GetSeriNFTByTokenID(ctx context.Context, tokenID int64) (*entity.SerialNFT, error)
	GetSeriNFTByTxHash(ctx context.Context, txHash string) (*entity.SerialNFT, error)

	Create(ctx context.Context, seriNFT *entity.SerialNFT) error
	UpdateTokenIDByTxHash(ctx context.Context, txHash string, tokenID uint64) error
	UpdateSeriNFTBySeri(ctx context.Context, seri string, fn UpdateSeriNFTFunc) error
}
