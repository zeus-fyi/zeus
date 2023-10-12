package web3_actions

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
	"github.com/zeus-fyi/gochain/web3/accounts"
)

func ValidateToAddress(ctx context.Context, toAddress string) error {
	if toAddress == "" {
		err := errors.New("the recipient address cannot be empty")
		log.Ctx(ctx).Err(err).Msg("Transfer: toAddress")
		return err
	}
	if !accounts.IsHexAddress(toAddress) {
		err := fmt.Errorf("invalid to 'address': %s", toAddress)
		log.Ctx(ctx).Err(err).Msg("Transfer: IsHexAddress")
		return err
	}
	return nil
}

func marshalJSON(ctx context.Context, data interface{}) (string, error) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("marshalJSON")
		return "", err
	}
	return string(b), err
}

func isValidUrl(toTest string) bool {
	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
func downloadFile(ctx context.Context, url string) ([]byte, error) {
	var dst bytes.Buffer
	response, err := http.Get(url)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("downloadFile: Get")
		return nil, err
	}
	defer response.Body.Close()
	_, err = io.Copy(&dst, response.Body)
	if err != nil {
		log.Ctx(ctx).Err(err).Msg("downloadFile: Copy")
		return nil, err
	}
	return dst.Bytes(), nil
}
