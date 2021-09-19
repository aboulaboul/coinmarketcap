package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {

  //define parameters for command line
  parameter := make(map[string]*string)
  parameter["apiKey"] = flag.String("api-key", "", "your personnal key to use coimarketcap api")
  parameter["convert"] = flag.String("convert", "EUR", "in which base currency converting (USD, EUR...)")
  parameter["limit"] = flag.String("limit", "200", "how many cryptos to fetch")
  //get parameters from command line
  flag.Parse()

  //build web request from parameters (source from coinmarketcap.com Crypto API documentation)
  client := &http.Client{}
  req, err := http.NewRequest("GET","https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest", nil)
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }
  q := url.Values{}
  q.Add("start", "1")
  q.Add("limit", *parameter["limit"])
  q.Add("convert", *parameter["convert"])
  req.Header.Set("Accepts", "application/json")
  req.Header.Add("X-CMC_PRO_API_KEY", *parameter["apiKey"])
  req.URL.RawQuery = q.Encode()

  //send request
  resp, err := client.Do(req);
  if err != nil {
    log.Println("Error sending request to server")
    os.Exit(1)
  }

  //handle request
  respBody, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Println(err)
    os.Exit(1)
  }

  //choosing EUR or USD
  var CoinMktCapOutPut OutPutter
  switch *parameter["convert"] {
  case "EUR":
    CoinMktCapOutPut = &CoinMktCapEUR{}
  case "USD":
    CoinMktCapOutPut = &CoinMktCapUSD{}
  default:
    log.Println("no base currency choosen")
  }

  //extracting from json
  err = CoinMktCapOutPut.Unmarshal(respBody)
  if err != nil {
	  log.Println(err)
  }
  CoinMktCapOutPut.Print()
}
