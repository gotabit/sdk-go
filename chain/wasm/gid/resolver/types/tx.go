package types

type SetAddress struct {
	Address  string `json:"address"`
	CoinType uint64 `json:"coin_type"`
	Node     []uint `json:"node"`
}

type MsgSetAddress struct {
	Msg SetAddress `json:"set_address"`
}

type SetTextData struct {
	Key   string `json:"key"`
	Node  []uint `json:"node"`
	Value string `json:"value"`
}

type MsgSetTextData struct {
	Msg SetTextData `json:"set_text_data"`
}

type SetContentHash struct {
	Hash []uint `json:"hash"`
	Node []uint `json:"node"`
}

type MsgSetContentHash struct {
	Msg SetContentHash `json:"set_content_hash"`
}

type MsgSetConfig struct {
	Msg Config `json:"set_config"`
}

type MsgDeleteTextData struct {
	Msg TextData `json:"delete_text_data"`
}
