package api_calls

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/pretty"
	"github.com/zeus-fyi/zeus/demos/api_calls/endpoints"
)

func GetBaseRestyClient() *resty.Client {
	client := resty.New()
	client.BaseURL = endpoints.ZeusApiEndpoint
	return client
}

func PrintReqJson(payload interface{}) {
	topologyActionRequestPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	fmt.Println("action request json")
	requestJSON := pretty.Pretty(topologyActionRequestPayload)
	requestJSON = pretty.Color(requestJSON, pretty.TerminalStyle)
	fmt.Println(string(requestJSON))
}

func PrintRespJson(body []byte) {
	fmt.Println("response json")
	respJSON := pretty.Pretty(body)
	respJSON = pretty.Color(respJSON, pretty.TerminalStyle)
	fmt.Println(string(respJSON))
}
