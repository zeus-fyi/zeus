package hercules_chain_snapshots

import (
	"strings"
)

type BucketRequest struct {
	BucketName string `json:"bucketName"`

	Protocol   string `json:"protocol"`
	Network    string `json:"network"`
	ClientType string `json:"clientType"`
	ClientName string `json:"clientName"`
}

func (b *BucketRequest) CreateBucketKey() string {
	key := []string{strings.ToLower(b.Protocol), strings.ToLower(b.Network), strings.ToLower(b.ClientName), strings.ToLower(b.ClientType)}
	return strings.Join(key, ".")
}
