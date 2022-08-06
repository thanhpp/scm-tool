package constx

import "time"

const (
	DefaultENVFile    = ".env"
	DefaultConfigFile = "config.yml"
	SaveFilePaths     = "./files"
)

const (
	RinkebyURL           = "https://rinkeby.infura.io"
	DefaultNFTServiceURL = "http://nftsrv:11000"
	RinkebyBurnAddress   = "0x000000000000000000000000000000000000dEaD"
)

const (
	DefaultHTTPClientTimeout         = time.Second * 10
	AutoUpdateTokenIDInterval        = time.Second * 2
	AutoTransferTokenInterval        = time.Second * 10
	AutoFallbackUpdateSerialInterval = time.Minute * 5
	AutoMintNFTInterval              = time.Minute
)

const (
	RabbitMQServerURL = "amqp://guest:guest@rabbitmq:5672/"
)
