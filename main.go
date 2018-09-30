package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
)

func main() {
	//connect to server, you can use your own transport settings
	fmt.Println("Dialling...")
	c, err := gosocketio.Dial(
		gosocketio.GetUrl("streamer.cryptocompare.com", 443, true),
		transport.GetDefaultWebsocketTransport(),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	btcusdSubscription, err := subtrades("BTC", "USD")
	if err != nil {
		fmt.Println(err)
		return
	}
	ethbtcSubscription, err := subtrades("ETH", "USD")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.Emit("SubAdd", map[string][]string{"subs": append(btcusdSubscription, ethbtcSubscription...)})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.On(gosocketio.OnConnection, handleOnConnection)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On(gosocketio.OnDisconnection, handleOnDisconnection)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On(gosocketio.OnError, handleOnError)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On("m", handleMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

}

type Subs struct {
	USD struct {
		TRADES     []string `json:"TRADES"`
		CURRENT    []string `json:"CURRENT"`
		CURRENTAGG string   `json:"CURRENTAGG"`
	} `json:"USD"`
}

func subtrades(from, to string) ([]string, error) {
	dataURL := fmt.Sprintf("https://min-api.cryptocompare.com/data/subs?fsym=%s&tsyms=%s", from, to)
	c := http.Client{Timeout: 5 * time.Second}
	resp, err := c.Get(dataURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("non 200")
	}
	subs := &Subs{}
	err = json.NewDecoder(resp.Body).Decode(subs)
	if err != nil {
		return nil, err
	}
	subSlice := []string{}
	subSlice = append(subSlice, subs.USD.TRADES...)
	subSlice = append(subSlice, subs.USD.CURRENT...)
	subSlice = append(subSlice, subs.USD.CURRENTAGG)

	return subSlice, nil

}

func current() {

}

func handleOnConnection(c *gosocketio.Channel, msg interface{}) string {
	fmt.Println(msg)
	return "OK"
}

func handleOnDisconnection(c *gosocketio.Channel, msg interface{}) string {
	fmt.Println(msg)
	return "OK"
}

func handleOnError(c *gosocketio.Channel, msg interface{}) string {
	fmt.Println(msg)
	return "OK"
}

type TradeMessage struct {
	SubscriptionID     string
	ExchangeName       string
	FromCurrencySymbol string
	ToCurrencySymbol   string
	Flag               string
	TradeID            string
	TimeStamp          string
	Quantity           string
	Price              string
	Total              string
}

func handleMessage(c *gosocketio.Channel, msg interface{}) string {
	data := msg.(string)
	dataSlice := strings.Split(data, "~")
	if dataSlice[0] == TRADE_MESSAGE {
		// fmt.Println("Received TRADE payload")
		// trade, err := unmarshalTradeMessage(dataSlice)
		// if err != nil {
		// 	fmt.Println("could not unmarshal trade")
		// 	return "Not OK"
		// }
		// fmt.Println(trade)
	}
	if dataSlice[0] == CURRENTAGG {
		// fmt.Println("Received CURRENTAGG payload")
		// currentAgg, err := unmarshalCurrentAgg(dataSlice)
		// if err != nil {
		// 	fmt.Println("could not unmarshal currentAgg")
		// 	return "Not OK"
		// }
		// fmt.Println(currentAgg)

	}
	if dataSlice[0] == TICKER {
		fmt.Println(data)
		// fmt.Println("Received TICKER payload. Updating ticker...")
		fl, err := strconv.ParseFloat(dataSlice[2], 64)
		if err != nil {
			fmt.Println("could not parse float:", err)
			return "Not OK"
		}
		ticker[dataSlice[1]] = fl
		for k, v := range ticker {
			fmt.Println(k, v)
		}

	}
	fmt.Println("m:", msg)
	return "OK"
}

const TRADE_MESSAGE = "0"
const TRADE_BUY = "1"
const TRADE_SELL = "2"
const TRADE_UNKNOWN = "4"
const CURRENTAGG = "5"
const TICKER = "11"

var ticker = map[string]float64{}

func unmarshalTradeMessage(dataSlice []string) (*TradeMessage, error) {
	return &TradeMessage{
		SubscriptionID:     dataSlice[0],
		ExchangeName:       dataSlice[1],
		FromCurrencySymbol: dataSlice[2],
		ToCurrencySymbol:   dataSlice[3],
		Flag:               dataSlice[4],
		TradeID:            dataSlice[5],
		TimeStamp:          dataSlice[6],
		Quantity:           dataSlice[7],
		Price:              dataSlice[8],
		Total:              dataSlice[9],
	}, nil
}

type CurrentAgg struct {
	Type         string
	ExchangeName string
	FromCurrency string
	ToCurrency   string
	Flag         string
	Price        string
	LastUpdate   string
	LastVolume   string
	LastVolumeTo string
	LastTradeID  string
	Volume24h    string
	Volume24hTo  string
	MaskInt      string
}

func unmarshalCurrentAgg(dataSlice []string) (*CurrentAgg, error) {
	return &CurrentAgg{
		Type:         dataSlice[0],
		ExchangeName: dataSlice[1],
		FromCurrency: dataSlice[2],
		ToCurrency:   dataSlice[3],
		Flag:         dataSlice[4],
		Price:        dataSlice[5],
		LastUpdate:   dataSlice[6],
		LastVolume:   dataSlice[7],
		LastVolumeTo: dataSlice[8],
		LastTradeID:  dataSlice[9],
		Volume24h:    dataSlice[10],
		Volume24hTo:  dataSlice[11],
		MaskInt:      dataSlice[12],
	}, nil
}
