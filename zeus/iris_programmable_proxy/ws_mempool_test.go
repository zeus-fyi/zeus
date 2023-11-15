package iris_programmable_proxy

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gorilla/websocket"
)

func (t *IrisConfigTestSuite) TestLiveMempoolWebsocket() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	addr := flag.String("addr", "iris.zeus.fyi", "ws service address")
	u := url.URL{Scheme: "wss", Host: *addr, Path: "/v1/mempool"}

	requestHeader := http.Header{}
	requestHeader.Add("Authorization", "Bearer "+t.Tc.Bearer)
	ws, _, werr := websocket.DefaultDialer.Dial(u.String(), requestHeader)
	t.Require().Nil(werr)
	defer ws.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := ws.ReadMessage()
			t.Require().Nil(err)
			fmt.Println(message)
			tx := &types.Transaction{}

			err = tx.UnmarshalBinary(message)
			if err != nil {
				continue
			}

			fmt.Println(tx.Size(), "bytes")
		}
	}()

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case tick := <-ticker.C:
			err := ws.WriteMessage(websocket.TextMessage, []byte(tick.String()))
			t.NoError(err)
		case <-interrupt:
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			t.NoError(err)
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
