package constx

import "time"

const (
	DefaultENVFile    = ".env"
	DefaultConfigFile = "config.yml"
	SaveFilePaths     = "./files"
)

const (
	RinkebyURL = "https://rinkeby.infura.io"
)

const (
	DefaultHTTPClientTimeout  = time.Second * 10
	AutoUpdateTokenIDInterval = time.Second * 2
)
