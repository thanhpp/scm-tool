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
)

const (
	DefaultHTTPClientTimeout        = time.Second * 10
	AutoUpdateTokenIDInterval       = time.Second * 2
	AutoMintAndUpdateSerialInterval = time.Second * 30
)
