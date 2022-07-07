package smartcontracts

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
