package hercules_chain_snapshots

import (
	"strings"
)

type BucketRequest struct {
	Network    string `json:"network"`
	ClientType string `json:"clientType"`
	ClientName string `json:"clientName"`
}

func (b *BucketRequest) CreateBucketKey() []string {
	key := []string{strings.ToLower(b.Network), strings.ToLower(b.ClientType), strings.ToLower(b.ClientName)}
	return key
}
