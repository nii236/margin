package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/graarh/golang-socketio"
	"github.com/graarh/golang-socketio/transport"
	"github.com/nii236/margin/helpers"
	"github.com/nii236/margin/models"
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
	// ethbtcSubscription, err := subtrades("ETH", "USD")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	err = c.Emit("SubAdd", map[string][]string{"subs": btcusdSubscription})
	if err != nil {
		fmt.Println(err)
		return
	}

	ctrl := &Controller{Mutex: &sync.Mutex{}}
	err = c.On(gosocketio.OnConnection, ctrl.HandleOnConnection)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On(gosocketio.OnDisconnection, ctrl.HandleOnDisconnection)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On(gosocketio.OnError, ctrl.HandleOnError)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = c.On("m", ctrl.HandleMessage)
	if err != nil {
		fmt.Println(err)
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()

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
	subs := &models.Subs{}
	err = json.NewDecoder(resp.Body).Decode(subs)
	if err != nil {
		return nil, err
	}
	subSlice := []string{}
	// subSlice = append(subSlice, subs.USD.TRADES...)
	subSlice = append(subSlice, subs.USD.CURRENT...)
	// subSlice = append(subSlice, subs.USD.CURRENTAGG)

	return subSlice, nil

}

type Controller struct {
	*sync.Mutex
}

func (ctrl *Controller) HandleOnConnection(c *gosocketio.Channel, msg interface{}) string {
	fmt.Println(msg)
	return "OK"
}

func (ctrl *Controller) HandleOnDisconnection(c *gosocketio.Channel, msg interface{}) string {
	fmt.Println(msg)
	return "OK"
}

func (ctrl *Controller) HandleOnError(c *gosocketio.Channel, msg interface{}) string {
	fmt.Println(msg)
	return "OK"
}

func (ctrl *Controller) HandleMessage(c *gosocketio.Channel, msg interface{}) string {
	ctrl.Mutex.Lock()
	defer ctrl.Mutex.Unlock()
	data := msg.(string)
	dataSlice := strings.Split(data, "~")
	if dataSlice[0] == SUBSCRIPTION_TRADE {
		// fmt.Println("Received TRADE payload")
		// trade, err := unmarshalTradeMessage(dataSlice)
		// if err != nil {
		// 	fmt.Println("could not unmarshal trade")
		// 	return "Not OK"
		// }
		// fmt.Println(trade)
	}
	if dataSlice[0] == SUBSCRIPTION_CURRENTAGG {
		// fmt.Println("Received CURRENTAGG payload")
		// currentAgg, err := unmarshalCurrentAgg(dataSlice)
		// if err != nil {
		// 	fmt.Println("could not unmarshal currentAgg")
		// 	return "Not OK"
		// }
		// fmt.Println(currentAgg)

	}
	if dataSlice[0] == SUBSCRIPTION_LOADCOMPLETE {
		return "OK"
	}
	if dataSlice[0] == SUBSCRIPTION_CURRENT {
		curr, err := helpers.UnpackCurrent(msg.(string))
		if err != nil {
			fmt.Println(err)
			return "Error"
		}

		msg := ""
		// if curr.TYPE != "" {
		// 	msg += "TYPE: " + curr.TYPE + " "
		// }
		if curr.MARKET != "" {
			msg += "MARKET: " + curr.MARKET + "\t"
		}
		if curr.FROMSYMBOL != "" {
			msg += "FROMSYMBOL: " + curr.FROMSYMBOL + "\t"
		}
		if curr.TOSYMBOL != "" {
			msg += "TOSYMBOL: " + curr.TOSYMBOL + "\t"
		}
		if curr.FLAGS != "" {
			msg += "FLAGS: " + curr.FLAGS + "\t"
		}
		if curr.PRICE != "" {
			msg += "PRICE: " + curr.PRICE + "\t"
		}
		if current[curr.MARKET] == nil {
			current[curr.MARKET] = &models.CurrentFlap{}
		}

		if current[curr.MARKET].Previous == nil {
			current[curr.MARKET].Previous = curr
			current[curr.MARKET].Next = curr
		} else {
			current[curr.MARKET].Previous = current[curr.MARKET].Next
			current[curr.MARKET].Next = helpers.MergeCurrent(current[curr.MARKET].Next, curr)
		}

		f, err := os.OpenFile("ticker.csv", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println(err)
			return "Error"
		}

		defer f.Close()
		if curr.FLAGS != "4" {
			w := csv.NewWriter(f)
			err = w.Write([]string{strconv.FormatInt(time.Now().UnixNano(), 10), curr.MARKET, curr.FROMSYMBOL, curr.TOSYMBOL, curr.PRICE})
			if err != nil {
				fmt.Println(err)
				return "Error"
			}
			fmt.Println("Appended CSV line")
			defer w.Flush()
		}
		return "OK"
	}
	return "OK"
}

const SUBSCRIPTION_TRADE = "0"
const TRADE_BUY = "1"
const TRADE_SELL = "2"
const TRADE_UNKNOWN = "4"
const SUBSCRIPTION_CURRENT = "2"
const SUBSCRIPTION_CURRENTAGG = "5"
const SUBSCRIPTION_LOADCOMPLETE = "3"

var ticker = map[string]float64{}
var current = map[string]*models.CurrentFlap{}
