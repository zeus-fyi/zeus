package ethereum_web3signer_actions

type PostRemoteKeys struct {
	RemoteKeys []struct {
		Pubkey string `json:"pubkey"`
		Url    string `json:"url"`
	} `json:"remote_keys"`
}

type DeleteRemoteKeysResp struct {
	Pubkeys []string `json:"pubkeys"`
}

type GetRemoteKeys struct {
	Data []struct {
		Pubkey   string `json:"pubkey"`
		Url      string `json:"url"`
		Readonly bool   `json:"readonly"`
	} `json:"data"`
}
