package resty_base

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/pretty"
)

type Resty struct {
	*resty.Client
	PrintReq  bool
	PrintResp bool
}

func GetBaseRestyTestClient(baseURL, bearer string) Resty {
	r := Resty{}
	r.Client = resty.New()
	r.Client.SetBaseURL(baseURL)
	r.Client.SetAuthToken(bearer)
	r.PrintResp = true
	return r
}

func (r *Resty) PrintReqJson(payload interface{}) {
	if !r.PrintReq {
		return
	}
	topologyActionRequestPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	fmt.Println("action request json")
	requestJSON := pretty.Pretty(topologyActionRequestPayload)
	requestJSON = pretty.Color(requestJSON, pretty.TerminalStyle)
	fmt.Println(string(requestJSON))
}

func (r *Resty) PrintRespJson(body []byte) {
	if !r.PrintResp {
		return
	}
	fmt.Println("response json")
	respJSON := pretty.Pretty(body)
	respJSON = pretty.Color(respJSON, pretty.TerminalStyle)
	fmt.Println(string(respJSON))
}
