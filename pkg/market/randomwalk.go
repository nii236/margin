package market

import (
	"context"
	"math/rand"
	"time"
)

type RandomWalk struct {
}

func NewRandomWalk() *RandomWalk {
	return &RandomWalk{}
}

type Ticker struct {
	Timestamp int64
	Exchange  string
	From      string
	To        string
	Price     float64
}

func (rw *RandomWalk) Subscribe(ctx context.Context) (chan *Ticker, error) {
	tickerTimer := time.NewTicker(5 * time.Second)
	tickerChan := make(chan *Ticker)
	go func(ctx context.Context, tickerChan chan *Ticker) {
		for {
			select {
			case <-tickerTimer.C:
				tickerChan <- &Ticker{
					Timestamp: time.Now().UnixNano(),
					Exchange:  "RandomWalk",
					From:      "BTC",
					To:        "USD",
					Price:     rand.Float64(),
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx, tickerChan)

	return tickerChan, nil
}
