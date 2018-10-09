package models

type CurrentFlap struct {
	Previous *CurrentMessage
	Next     *CurrentMessage
}
type CurrentMessage struct {
	TYPE           string
	MARKET         string
	FROMSYMBOL     string
	TOSYMBOL       string
	FLAGS          string
	PRICE          string
	BID            string
	OFFER          string
	LASTUPDATE     string
	AVG            string
	LASTVOLUME     string
	LASTVOLUMETO   string
	LASTTRADEID    string
	VOLUMEHOUR     string
	VOLUMEHOURTO   string
	VOLUME24HOUR   string
	VOLUME24HOURTO string
	OPENHOUR       string
	HIGHHOUR       string
	LOWHOUR        string
	OPEN24HOUR     string
	HIGH24HOUR     string
	LOW24HOUR      string
	LASTMARKET     string
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

type Subs struct {
	USD struct {
		TRADES     []string `json:"TRADES"`
		CURRENT    []string `json:"CURRENT"`
		CURRENTAGG string   `json:"CURRENTAGG"`
	} `json:"USD"`
}
