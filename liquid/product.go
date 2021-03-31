package liquid

import "fmt"

type Product struct {
	ID                      string      `json:"id"`
	ProductType             string      `json:"product_type"`
	Code                    string      `json:"code"`
	Name                    interface{} `json:"name"`
	MarketAsk               float64     `json:"market_ask"`
	MarketBid               float64     `json:"market_bid"`
	Indicator               int         `json:"indicator"`
	Currency                string      `json:"currency"`
	CurrencyPairCode        string      `json:"currency_pair_code"`
	Symbol                  string      `json:"symbol"`
	BtcMinimumWithdraw      interface{} `json:"btc_minimum_withdraw"`
	FiatMinimumWithdraw     interface{} `json:"fiat_minimum_withdraw"`
	PusherChannel           string      `json:"pusher_channel"`
	TakerFee                string      `json:"taker_fee"`
	MakerFee                string      `json:"maker_fee"`
	LowMarketBid            string      `json:"low_market_bid"`
	HighMarketAsk           string      `json:"high_market_ask"`
	Volume24H               string      `json:"volume_24h"`
	LastPrice24H            string      `json:"last_price_24h"`
	LastTradedPrice         string      `json:"last_traded_price"`
	LastTradedQuantity      string      `json:"last_traded_quantity"`
	AveragePrice            string      `json:"average_price"`
	QuotedCurrency          string      `json:"quoted_currency"`
	BaseCurrency            string      `json:"base_currency"`
	TickSize                string      `json:"tick_size"`
	Disabled                bool        `json:"disabled"`
	MarginEnabled           bool        `json:"margin_enabled"`
	CfdEnabled              bool        `json:"cfd_enabled"`
	PerpetualEnabled        bool        `json:"perpetual_enabled"`
	LastEventTimestamp      string      `json:"last_event_timestamp"`
	Timestamp               string      `json:"timestamp"`
	MultiplierUp            string      `json:"multiplier_up"`
	MultiplierDown          string      `json:"multiplier_down"`
	AverageTimeInterval     int         `json:"average_time_interval"`
	ProgressiveTierEligible bool        `json:"progressive_tier_eligible"`
	ExchangeRate            int         `json:"exchange_rate"`
}

func (api API) GetProduct(id string) (Product, error) {
	var product Product
	if err := api.Get(fmt.Sprintf("/products/%s", id), &product); err != nil {
		return Product{}, err
	}
	return product, nil
}
