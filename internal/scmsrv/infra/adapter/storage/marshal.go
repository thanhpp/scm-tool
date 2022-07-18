package storage

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
)

func marshalSerial(in entity.Serial) repo.Serial {
	return repo.Serial{
		Seri:           in.Seri,
		Status:         in.Status,
		TokenID:        in.TokenID,
		ItemSKU:        in.Item.SKU,
		ImportTicketID: in.ImportTicket.ID,
	}
}

func unmarshalSerial(in repo.Serial) entity.Serial {
	return entity.Serial{
		Seri:         in.Seri,
		Status:       in.Status,
		TokenID:      in.TokenID,
		Item:         unmarshalItem(in.Item),
		ImportTicket: unmarshalImportTicket(&in.ImportTicket),
	}
}
