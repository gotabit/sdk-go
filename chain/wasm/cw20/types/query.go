package types

type Cw20Balance struct {
	Address string `json:"address"`
}

type Cw20BalanceQuery struct {
	Query Cw20Balance `json:"balance"`
}
