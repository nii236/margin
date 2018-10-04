package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func handleMessage(c *gosocketio.Channel, msg interface{}) string {
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
		if curr.TYPE != "" {
			msg += "TYPE: " + curr.TYPE + " "
		}
		if curr.MARKET != "" {
			msg += "MARKET: " + curr.MARKET + " "
		}
		if curr.FROMSYMBOL != "" {
			msg += "FROMSYMBOL: " + curr.FROMSYMBOL + " "
		}
		if curr.TOSYMBOL != "" {
			msg += "TOSYMBOL: " + curr.TOSYMBOL + " "
		}
		if curr.FLAGS != "" {
			msg += "FLAGS: " + curr.FLAGS + " "
		}
		if curr.PRICE != "" {
			msg += "PRICE: " + curr.PRICE + " "
		}
		if curr.BID != "" {
			msg += "BID: " + curr.BID + " "
		}
		if curr.OFFER != "" {
			msg += "OFFER: " + curr.OFFER + " "
		}
		if curr.LASTUPDATE != "" {
			msg += "LASTUPDATE: " + curr.LASTUPDATE + " "
		}
		if curr.AVG != "" {
			msg += "AVG: " + curr.AVG + " "
		}
		if curr.LASTVOLUME != "" {
			msg += "LASTVOLUME: " + curr.LASTVOLUME + " "
		}
		if curr.LASTVOLUMETO != "" {
			msg += "LASTVOLUMETO: " + curr.LASTVOLUMETO + " "
		}
		if curr.LASTTRADEID != "" {
			msg += "LASTTRADEID: " + curr.LASTTRADEID + " "
		}
		if curr.VOLUMEHOUR != "" {
			msg += "VOLUMEHOUR: " + curr.VOLUMEHOUR + " "
		}
		if curr.VOLUMEHOURTO != "" {
			msg += "VOLUMEHOURTO: " + curr.VOLUMEHOURTO + " "
		}
		if curr.VOLUME24HOUR != "" {
			msg += "VOLUME24HOUR: " + curr.VOLUME24HOUR + " "
		}
		if curr.VOLUME24HOURTO != "" {
			msg += "VOLUME24HOURTO: " + curr.VOLUME24HOURTO + " "
		}
		if curr.OPENHOUR != "" {
			msg += "OPENHOUR: " + curr.OPENHOUR + " "
		}
		if curr.HIGHHOUR != "" {
			msg += "HIGHHOUR: " + curr.HIGHHOUR + " "
		}
		if curr.LOWHOUR != "" {
			msg += "LOWHOUR: " + curr.LOWHOUR + " "
		}
		if curr.OPEN24HOUR != "" {
			msg += "OPEN24HOUR: " + curr.OPEN24HOUR + " "
		}
		if curr.HIGH24HOUR != "" {
			msg += "HIGH24HOUR: " + curr.HIGH24HOUR + " "
		}
		if curr.LOW24HOUR != "" {
			msg += "LOW24HOUR: " + curr.LOW24HOUR + " "
		}
		if curr.LASTMARKET != "" {
			msg += "LASTMARKET: " + curr.LASTMARKET + " "
		}
		next := &models.CurrentMessage{}
		if current[curr.MARKET] == nil {
			current[curr.MARKET] = &models.CurrentFlap{}
		}
		if curr.TYPE == "" && current[curr.MARKET].Previous != nil {
			next.TYPE = current[curr.MARKET].Previous.TYPE
		} else {
			next.TYPE = curr.TYPE
		}
		if curr.MARKET == "" && current[curr.MARKET].Previous != nil {
			next.MARKET = current[curr.MARKET].Previous.MARKET
		} else {
			next.MARKET = curr.MARKET
		}
		if curr.FROMSYMBOL == "" && current[curr.MARKET].Previous != nil {
			next.FROMSYMBOL = current[curr.MARKET].Previous.FROMSYMBOL
		} else {
			next.FROMSYMBOL = curr.FROMSYMBOL
		}
		if curr.TOSYMBOL == "" && current[curr.MARKET].Previous != nil {
			next.TOSYMBOL = current[curr.MARKET].Previous.TOSYMBOL
		} else {
			next.TOSYMBOL = curr.TOSYMBOL
		}
		if curr.FLAGS == "" && current[curr.MARKET].Previous != nil {
			next.FLAGS = current[curr.MARKET].Previous.FLAGS
		} else {
			next.FLAGS = curr.FLAGS
		}
		if curr.PRICE == "" && current[curr.MARKET].Previous != nil {
			next.PRICE = current[curr.MARKET].Previous.PRICE
		} else {
			next.PRICE = curr.PRICE
		}
		if curr.BID == "" && current[curr.MARKET].Previous != nil {
			next.BID = current[curr.MARKET].Previous.BID
		} else {
			next.BID = curr.BID
		}
		if curr.OFFER == "" && current[curr.MARKET].Previous != nil {
			next.OFFER = current[curr.MARKET].Previous.OFFER
		} else {
			next.OFFER = curr.OFFER
		}
		if curr.LASTUPDATE == "" && current[curr.MARKET].Previous != nil {
			next.LASTUPDATE = current[curr.MARKET].Previous.LASTUPDATE
		} else {
			next.LASTUPDATE = curr.LASTUPDATE
		}
		if curr.AVG == "" && current[curr.MARKET].Previous != nil {
			next.AVG = current[curr.MARKET].Previous.AVG
		} else {
			next.AVG = curr.AVG
		}
		if curr.LASTVOLUME == "" && current[curr.MARKET].Previous != nil {
			next.LASTVOLUME = current[curr.MARKET].Previous.LASTVOLUME
		} else {
			next.LASTVOLUME = curr.LASTVOLUME
		}
		if curr.LASTVOLUMETO == "" && current[curr.MARKET].Previous != nil {
			next.LASTVOLUMETO = current[curr.MARKET].Previous.LASTVOLUMETO
		} else {
			next.LASTVOLUMETO = curr.LASTVOLUMETO
		}
		if curr.LASTTRADEID == "" && current[curr.MARKET].Previous != nil {
			next.LASTTRADEID = current[curr.MARKET].Previous.LASTTRADEID
		} else {
			next.LASTTRADEID = curr.LASTTRADEID
		}
		if curr.VOLUMEHOUR == "" && current[curr.MARKET].Previous != nil {
			next.VOLUMEHOUR = current[curr.MARKET].Previous.VOLUMEHOUR
		} else {
			next.VOLUMEHOUR = curr.VOLUMEHOUR
		}
		if curr.VOLUMEHOURTO == "" && current[curr.MARKET].Previous != nil {
			next.VOLUMEHOURTO = current[curr.MARKET].Previous.VOLUMEHOURTO
		} else {
			next.VOLUMEHOURTO = curr.VOLUMEHOURTO
		}
		if curr.VOLUME24HOUR == "" && current[curr.MARKET].Previous != nil {
			next.VOLUME24HOUR = current[curr.MARKET].Previous.VOLUME24HOUR
		} else {
			next.VOLUME24HOUR = curr.VOLUME24HOUR
		}
		if curr.VOLUME24HOURTO == "" && current[curr.MARKET].Previous != nil {
			next.VOLUME24HOURTO = current[curr.MARKET].Previous.VOLUME24HOURTO
		} else {
			next.VOLUME24HOURTO = curr.VOLUME24HOURTO
		}
		if curr.OPENHOUR == "" && current[curr.MARKET].Previous != nil {
			next.OPENHOUR = current[curr.MARKET].Previous.OPENHOUR
		} else {
			next.OPENHOUR = curr.OPENHOUR
		}
		if curr.HIGHHOUR == "" && current[curr.MARKET].Previous != nil {
			next.HIGHHOUR = current[curr.MARKET].Previous.HIGHHOUR
		} else {
			next.HIGHHOUR = curr.HIGHHOUR
		}
		if curr.LOWHOUR == "" && current[curr.MARKET].Previous != nil {
			next.LOWHOUR = current[curr.MARKET].Previous.LOWHOUR
		} else {
			next.LOWHOUR = curr.LOWHOUR
		}
		if curr.OPEN24HOUR == "" && current[curr.MARKET].Previous != nil {
			next.OPEN24HOUR = current[curr.MARKET].Previous.OPEN24HOUR
		} else {
			next.OPEN24HOUR = curr.OPEN24HOUR
		}
		if curr.HIGH24HOUR == "" && current[curr.MARKET].Previous != nil {
			next.HIGH24HOUR = current[curr.MARKET].Previous.HIGH24HOUR
		} else {
			next.HIGH24HOUR = curr.HIGH24HOUR
		}
		if curr.LOW24HOUR == "" && current[curr.MARKET].Previous != nil {
			next.LOW24HOUR = current[curr.MARKET].Previous.LOW24HOUR
		} else {
			next.LOW24HOUR = curr.LOW24HOUR
		}
		if curr.LASTMARKET == "" && current[curr.MARKET].Previous != nil {
			next.LASTMARKET = current[curr.MARKET].Previous.LASTMARKET
		} else {
			next.LASTMARKET = curr.LASTMARKET
		}

		if current[curr.MARKET].Previous == nil {
			current[curr.MARKET].Previous = next
			current[curr.MARKET].Next = next
		} else {
			current[curr.MARKET].Previous = current[curr.MARKET].Next
			current[curr.MARKET].Next = next
		}

		fmt.Println(next)
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
