package storage

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
)

func marshalSerial(in entity.Serial) repo.Serial {
	return repo.Serial{
		Seri:           in.Seri,
		ItemSKU:        in.Item.SKU,
		StorageID:      in.Storage.ID,
		ImportTicketID: in.ImportTicket.ID,
	}
}

func unmarshalSeri(in repo.Serial) entity.Serial {
	return entity.Serial{
		Seri: in.Seri,
	}
}
