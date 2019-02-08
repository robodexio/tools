package orderBook

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"../order"
	"../db"
)

type orderBook struct {
	Table  string   `json:"table"`
	Symbol string `json:"symbol"`
	Data []order.Order `json:"data"`
}
var source *db.Db

func Routes(db *db.Db) *chi.Mux {
	source = db
	router := chi.NewRouter()
	router.Get("/{Symbol}", GetOrderBook)

	return router
}

func GetOrderBook(w http.ResponseWriter, r *http.Request)  {
	symbol := chi.URLParam(r, "Symbol")
	source.Exec("SElect Order from FB where orderID = OrderID")

	data := []order.Order{
		{
			Symbol:  symbol,
			ID: "ksdhgkjsdhfg",
			Side: "Sell",
			Size:  10 ,
			Price: 120.32,
		},
	}
	var orderBook = orderBook{
		Table: "orderBook",
		Symbol: symbol,
		Data: data,
	}
	render.JSON(w, r, orderBook) // A chi router helper for serializing and returning json
}
