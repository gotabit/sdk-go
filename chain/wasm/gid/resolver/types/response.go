package types

type AllTextDataResponse struct {
	Data  []string `json:"data"`
	Total uint64   `json:"total"`
}

type ConfigResponse Config

type ContentHashResponse struct {
	Hash []uint `json:"hash"`
}

type TextDataResponse struct {
	Data string `json:"data"`
}
