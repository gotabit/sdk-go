package types

type Config struct {
	IntefaceId      uint64 `json:"interface_id"`
	Owner           string `json:"owner"`
	RegistryAddress string `json:"registry_address"`
}

type TextData struct {
	Key  string `json:"key"`
	Node []uint `json:"node"`
}
