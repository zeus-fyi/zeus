package bls_serverless_signatures

import aegis_inmemdbs "github.com/zeus-fyi/zeus/pkg/aegis/inmemdbs"

type SignatureRequests struct {
	SecretName        string                                         `json:"secretName"`
	SignatureRequests aegis_inmemdbs.EthereumBLSKeySignatureRequests `json:"signatureRequests"`
}
