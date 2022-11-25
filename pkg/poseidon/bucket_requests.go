package poseidon

import "strings"

type BucketRequest struct {
	BucketName string `json:"bucketName"`
	BucketKey  string `json:"bucketKey,omitempty"`

	Protocol        string `json:"protocol"`
	Network         string `json:"network"`
	ClientType      string `json:"clientType"`
	ClientName      string `json:"clientName"`
	CompressionType string `json:"compressionType,omitempty"`
}

func (b *BucketRequest) CreateBucketKey() string {
	key := []string{strings.ToLower(b.Protocol), strings.ToLower(b.Network), strings.ToLower(b.ClientName), strings.ToLower(b.ClientType)}

	if len(b.CompressionType) == 0 {
		key = append(key, b.CompressionType)
	}
	return strings.Join(key, ".")
}
