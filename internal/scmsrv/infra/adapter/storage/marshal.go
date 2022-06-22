package storage

import (
	"github.com/thanhpp/scm/internal/scmsrv/domain/entity"
	"github.com/thanhpp/scm/internal/scmsrv/domain/repo"
)

func marshalSerial(in entity.Serial) repo.Serial {
	return repo.Serial{
		Seri:           in.Seri,
		ItemSKU:        in.Item.SKU,
		ImportTicketID: in.ImportTicket.ID,
	}
}
