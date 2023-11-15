package iris_programmable_proxy

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
)

func EstablishLongRunningWs(bearer string) {
	addr := flag.String("addr", "iris.zeus.fyi", "ws service address")
	u := url.URL{Scheme: "wss", Host: *addr, Path: "/v1/mempool"}

	requestHeader := http.Header{}
	requestHeader.Add("Authorization", "Bearer "+bearer)

	// Send the subscription request to the WebSocket server.
	for {
		ws, _, werr := websocket.DefaultDialer.Dial(u.String(), requestHeader)
		// Create subscription request.
		if werr != nil {
			panic(werr)
		}
		err := WsLoop(ws)
		if err != nil {
			ws.Close()
		}
	}

}

func WsLoop(ws *websocket.Conn) error {
	i := 0

	m := make(map[string]int)
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			return err
		}
		tx := &types.Transaction{}
		err = tx.UnmarshalBinary(message)
		if err != nil {
			fmt.Println("serialization error", err)
			err = nil
			continue
		}
		m[tx.Hash().Hex()]++
		// occasional duplication is expected as a tradeoff for speed
		fmt.Println(i, m[tx.Hash().Hex()], "count", tx.Hash().Hex(), "txHash")
		i++
	}
}
