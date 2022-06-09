package enum

//go:generate go run github.com/dmarkham/enumer -type=ImportTicketStatus -linecomment -json=true -sql=true
type ImportTicketStatus int

const (
	ImportTicketStatusNew ImportTicketStatus = iota + 1
	ImportTicketStatusDone
)
