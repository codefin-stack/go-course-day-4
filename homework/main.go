package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

type TickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

var symbols = []string{"BTCUSDT", "ETHUSDT", "BNBUSDT", "XRPUSDT", "SOLUSDT"}

func fetchPrice(symbol string, wg *sync.WaitGroup, priceMap *sync.Map) {
	defer wg.Done()
	url := fmt.Sprintf("https://api.binance.com/api/v3/ticker/price?symbol=%s", symbol)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching price for %s: %v\n", symbol, err)
		return
	}
	defer resp.Body.Close()
	var ticker TickerResponse
	if err := json.NewDecoder(resp.Body).Decode(&ticker); err != nil {
		fmt.Printf("Error decoding response for %s: %v\n", symbol, err)
		return
	}
	priceMap.Store(symbol, ticker.Price)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func displayPrices(symbol string, currentPrice string) {
	previousPrice, ok := lastPrices.Load(symbol)
	if ok {
		prev, _ := strconv.ParseFloat(previousPrice.(string), 64)
		curr, _ := strconv.ParseFloat(currentPrice, 64)
		if curr > prev {
			fmt.Printf("%s: %s%s %s%s\n", symbol, green, currentPrice, upArrow, reset)
		} else if curr < prev {
			fmt.Printf("%s: %s%s %s%s\n", symbol, red, currentPrice, downArrow, reset)
		} else {
			fmt.Printf("%s: %s\n", symbol, currentPrice)
		}
	} else {
		fmt.Printf("%s: %s\n", symbol, currentPrice)
	}
	lastPrices.Store(symbol, currentPrice)
}

var lastPrices = sync.Map{}

const (
	green     = "\033[32m"
	red       = "\033[31m"
	reset     = "\033[0m"
	upArrow   = "↑"
	downArrow = "↓"
)

func main() {
	for {
		var wg sync.WaitGroup
		priceMap := &sync.Map{}
		wg.Add(len(symbols))
		for _, symbol := range symbols {
			go fetchPrice(symbol, &wg, priceMap)
		}
		wg.Wait()
		clearScreen()
		fmt.Println("Latest Crypto Prices")
		for _, symbol := range symbols {
			if price, ok := priceMap.Load(symbol); ok {
				displayPrices(symbol, price.(string))
			}
		}
		time.Sleep(300 * time.Millisecond)
	}
}
