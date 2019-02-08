package main

import (
	"encoding/json"
	"fmt"
	"github.com/vmpartner/bitmex/websocket"
	"strings"
)

type orderBook struct {
	Table  string   `json:"table"`
	Action string   `json:"action"`
	Keys   []string `json:"keys"`
	Types  struct {
		Symbol string `json:"symbol"`
		ID     string `json:"id"`
		Side   string `json:"side"`
		Size   string `json:"size"`
		Price  string `json:"price"`
	} `json:"types"`
	ForeignKeys struct {
		Symbol string `json:"symbol"`
		Side   string `json:"side"`
	} `json:"foreignKeys"`
	Attributes struct {
		Symbol string `json:"symbol"`
		ID     string `json:"id"`
	} `json:"attributes"`
	Filter struct {
		Symbol string `json:"symbol"`
	} `json:"filter"`
	Data []struct {
		Symbol string  `json:"symbol"`
		ID     int64   `json:"id"`
		Side   string  `json:"side"`
		Size   int     `json:"size"`
		Price  float64 `json:"price"`
	} `json:"data"`
}

type order struct {
	Symbol string  `json:"symbol"`
	ID     int64   `json:"id"`
	Side   string  `json:"side"`
	Size   int     `json:"size"`
	Price  float64 `json:"price"`
}

func e(err error) {
	if err != nil {
		panic(err)
	}
}

func createLimitOrder(order order) {

}

func createMarketOrder(order order) {
	//Data
	//Symbol = {string} "ETHUSD"
	//ID = {int64} 0
	//Side = {string} "Sell"
	//Size = {int} 50
	//Price = {float64} 101.85

}

func updateLimitOrder(order order) {

}

func deleteLimitOrder(order order) {

}
func orderBookStorage(bookData []order) {
	// create local snapshot of orderbook
}



// RoboDEX market maker
func main() {

	// Connect to bitMex WS
	bmx := websocket.Connect("www.bitmex.com")
	defer bmx.Close()

	// Listen read WS
	chReadFromWS := make(chan []byte, 100)
	go websocket.ReadFromWSToChannel(bmx, chReadFromWS)

	// Listen write WS
	chWriteToWS := make(chan interface{}, 100)
	go websocket.WriteFromChannelToWS(bmx, chWriteToWS)

	// Subscribe
	subMsm := websocket.Message{Op: "subscribe"}
	subMsm.AddArgument("orderBookL2_25:ETHUSD")
	subMsm.AddArgument("trade:ETHUSD")
	chWriteToWS <- subMsm

	// Read first response message
	message := <-chReadFromWS
	if !strings.Contains(string(message), "Welcome to the BitMEX") {
		fmt.Println(string(message))
		panic("No welcome message")
	}

	// Read auth response success
	message = <-chReadFromWS
	//Data, err := bitmex.DecodeMessage(message)
	//CheckErr(err)
	//fmt.Printf("%+v\n", Data)
	// Listen websocket before subscribe

	go func() {

		for {
			message := <-chReadFromWS
			var encodedJson orderBook
			err := json.Unmarshal([]byte(message), &encodedJson)
			e(err)
			fmt.Println(encodedJson)

			for _ , row := range encodedJson.Data {
				switch encodedJson.Table{
				case "orderBookL2_25":
					switch encodedJson.Action{
					case "partial": //create initial orders
						go createLimitOrder(row)
					case "update":  //update orderSize
						go updateLimitOrder(row)
					case "delete":  //delete order
						go deleteLimitOrder(row)
					case "insert":	//new order
						go createLimitOrder(row)
					}
				case "trades":
					createMarketOrder(row)

				}
			}


		}
	}()

	select {}
}

