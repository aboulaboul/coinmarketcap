package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type OutPutter interface {
	Unmarshal(b []byte) error
	Print()
}


type Currency struct {
	Price                 float64   `json:"price"`
	Volume24H             float64   `json:"volume_24h"`
	PercentChange1H       float64   `json:"percent_change_1h"`
	PercentChange24H      float64   `json:"percent_change_24h"`
	PercentChange7D       float64   `json:"percent_change_7d"`
	PercentChange30D      float64   `json:"percent_change_30d"`
	PercentChange60D      float64   `json:"percent_change_60d"`
	PercentChange90D      float64   `json:"percent_change_90d"`
	MarketCap             float64   `json:"market_cap"`
	MarketCapDominance    float64   `json:"market_cap_dominance"`
	FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
	LastUpdated           time.Time `json:"last_updated"`
}


type CoinMktCapEUR struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
		TotalCount   int         `json:"total_count"`
	} `json:"status"`
	Data []struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		Slug              string      `json:"slug"`
		NumMarketPairs    int         `json:"num_market_pairs"`
		DateAdded         time.Time   `json:"date_added"`
		Tags              []string    `json:"tags"`
		MaxSupply         float64         `json:"max_supply"`
		CirculatingSupply float64         `json:"circulating_supply"`
		TotalSupply       float64         `json:"total_supply"`
		Platform          interface{} `json:"platform"`
		CmcRank           int         `json:"cmc_rank"`
		LastUpdated       time.Time   `json:"last_updated"`
		Quote             struct {
			Currency `json:"EUR"`
		} `json:"quote"`
	} `json:"data"`
}

func (cmc *CoinMktCapEUR) Print(){
	fmt.Println("CmcRank, Symbol, Name, LastUpdated, PriceEUR, Volume24H, %Chg1H, %Chg24H, %Chg7d, %Chg30d, %Chg60d, %Chg90d")
	for _, c := range cmc.Data{
		fmt.Printf("%v, %v, %v, %v, %.2f, %.0f, %.1f, %.1f, %.1f, %.1f, %.1f, %.1f\n",
			c.CmcRank,
			c.Symbol,
			c.Name,
			c.LastUpdated.Format("2006/01/02 03:04:05"),
			c.Quote.Currency.Price,
			c.Quote.Volume24H,
			c.Quote.PercentChange1H,
			c.Quote.PercentChange24H,
			c.Quote.PercentChange7D,
			c.Quote.PercentChange30D,
			c.Quote.PercentChange60D,
			c.Quote.PercentChange90D)
		}
}


func (cmc *CoinMktCapEUR) Unmarshal(b []byte) error {
	return json.Unmarshal(b, cmc)
}



type CoinMktCapUSD struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
		TotalCount   int         `json:"total_count"`
	} `json:"status"`
	Data []struct {
		ID                int         `json:"id"`
		Name              string      `json:"name"`
		Symbol            string      `json:"symbol"`
		Slug              string      `json:"slug"`
		NumMarketPairs    int         `json:"num_market_pairs"`
		DateAdded         time.Time   `json:"date_added"`
		Tags              []string    `json:"tags"`
		MaxSupply         float64         `json:"max_supply"`
		CirculatingSupply float64         `json:"circulating_supply"`
		TotalSupply       float64         `json:"total_supply"`
		Platform          interface{} `json:"platform"`
		CmcRank           int         `json:"cmc_rank"`
		LastUpdated       time.Time   `json:"last_updated"`
		Quote             struct {
			Currency `json:"EUR"`
		} `json:"quote"`
	} `json:"data"`
}

func (cmc *CoinMktCapUSD) Print(){
	fmt.Println("CmcRank, Symbol, Name, LastUpdated, PriceUSD, Volume24H, %Chg1H, %Chg24H, %Chg7d, %Chg30d, %Chg60d, %Chg90d")
	for _, c := range cmc.Data{
		fmt.Printf("%v, %v, %v, %v, %.2f, %.0f, %.1f, %.1f, %.1f, %.1f, %.1f, %.1f\n",
			c.CmcRank,
			c.Symbol,
			c.Name,
			c.LastUpdated.Format("2006/01/02 03:04:05"),
			c.Quote.Currency.Price,
			c.Quote.Volume24H,
			c.Quote.PercentChange1H,
			c.Quote.PercentChange24H,
			c.Quote.PercentChange7D,
			c.Quote.PercentChange30D,
			c.Quote.PercentChange60D,
			c.Quote.PercentChange90D)
	}
}

func (cmc *CoinMktCapUSD) Unmarshal(b []byte) error {
	return json.Unmarshal(b, cmc)
}
