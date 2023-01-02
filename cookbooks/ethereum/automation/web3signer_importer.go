package ethereum_automation_cookbook

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type Web3SignerKeystores struct {
	Keystores          []string `json:"keystores"`
	Passwords          []string `json:"passwords"`
	SlashingProtection string   `json:"slashing_protection,omitempty"`
}

func ReadKeystores(ctx context.Context, b []byte) Web3SignerKeystores {
	// TODO
	keystores := Web3SignerKeystores{}
	m := make(map[string]interface{})
	err := json.Unmarshal(b, &m)
	if err != nil {
		log.Ctx(ctx).Panic().Err(err)
		panic(err)
	}
	keystores.Keystores = make([]string, len(m))
	//count := 0
	//for i, k := range m {
	//
	//	keystores.Keystores[count] = k.(string)
	//}
	return Web3SignerKeystores{}
}

/*
	ks := keystorev4.New()
	//enc, err := ks.Encrypt(sk.Marshal(), vd.Pw)
	//if err != nil {
	//	log.Ctx(ctx).Err(err)
	//	return  Web3SignerKeystores{}, err
	//}
*/
