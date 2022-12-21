package types

type SetBounty struct {
	Round uint64 `json:"round"`
}

type MsgSetBounty struct {
	Msg SetBounty `json:"set_bounty"`
}

type Add struct {
	PreviousSignature []uint `json:"previous_signature"`
	Round             uint64 `json:"round"`
	Signature         []uint `json:"signature"`
}

type MsgAdd struct {
	Msg Add `json:"add"`
}
