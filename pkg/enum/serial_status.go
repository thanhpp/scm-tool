package enum

//go:generate go run github.com/dmarkham/enumer -type=SerialStatus -linecomment -json=true -sql=true
type SerialStatus int

const (
	SerialStatusNew SerialStatus = iota + 1
	SerialStatusSold
)
