package market

// Provider is a market data provider (real-time)
type Provider interface {
	Subscribe() chan *Ticker
}
