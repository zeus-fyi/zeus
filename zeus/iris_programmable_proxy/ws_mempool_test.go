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
	requestHeader.Add("Authorization", "Bearer "+t.BearerToken)
	ws, _, werr := websocket.DefaultDialer.Dial(u.String(), requestHeader)
	t.Require().Nil(werr)
	defer ws.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := ws.ReadMessage()
			t.Require().Nil(err)
			tx := &types.Transaction{}

			err = tx.UnmarshalBinary(message)
			if err != nil {
				continue
			}
			fmt.Printf("tx-hash: %s\n", tx.Hash().String())
			fmt.Printf("value: %s\n", tx.Value().String())
			fmt.Printf("to: %s\n", tx.To().String())
			fmt.Println("tx data size \n", tx.Size())
		}
	}()

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	// Adding a 5-second timer
	shutdownTimer := time.After(5 * time.Second)

	for {
		select {
		case <-done:
			return
		case s := <-ticker.C:
			err := ws.WriteMessage(websocket.TextMessage, []byte(s.String()))
			t.NoError(err)
		case <-shutdownTimer:
			// Close the connection and terminate the function after 5 seconds
			err := ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Shutting down after 5 seconds"))
			t.NoError(err)
			return
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
