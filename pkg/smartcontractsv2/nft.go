package smartcontractsv2

const (
	ContractAddress = "0xEA92d9D15139c316944aeba2aBe828B87cBA1265"
)

type NFT struct {
	Name       string      `json:"name"`
	Desc       string      `json:"desc"`
	Image      string      `json:"image"`
	Attributes []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}
