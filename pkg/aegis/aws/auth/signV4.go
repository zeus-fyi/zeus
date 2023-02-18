package aws_aegis_auth

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"io"
	"net/http"
	"strings"
	"time"
)

func (a *AuthAWS) CreateV4AuthPOSTReq(ctx context.Context, service, url string, payload any) (*http.Request, error) {
	creds := a.GetCredentials(ctx)
	signer := v4.NewSigner()

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	reader := strings.NewReader(string(b))
	h := sha256.New()
	_, err = io.Copy(h, reader)
	if err != nil {
		return nil, err
	}
	payloadHash := hex.EncodeToString(h.Sum(nil))
	now := time.Now()
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}
	err = signer.SignHTTP(ctx, creds, req, payloadHash, service, a.Region, now)
	return req, nil
}
