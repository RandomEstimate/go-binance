package futures

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestWsTradeServe(t *testing.T) {
	_, _, err := WsTradeServe("bnbusdt", func(event *WsTradeEvent) {

		marshal, err := json.Marshal(event)
		if err != nil {
			return
		}
		fmt.Println(string(marshal))
	}, func(err error) {
		fmt.Println(err)
	})
	if err != nil {
		return
	}

	select {}

}

func TestWsCombinedTradeServe(t *testing.T) {
	_, _, err := WsCombinedTradeServe([]string{"bnbusdt", "btcusdt"}, func(event *WsTradeEvent) {

		marshal, err := json.Marshal(event)
		if err != nil {
			return
		}
		fmt.Println(string(marshal))
	}, func(err error) {
		fmt.Println(err)
	})
	if err != nil {
		return
	}

	select {}

}
