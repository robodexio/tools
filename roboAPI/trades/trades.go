package trades


import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"../order"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{Symbol}", GetTrades)
	return router
}

func GetTrades(w http.ResponseWriter, r *http.Request) {

	symbol := chi.URLParam(r, "Symbol")
	trades := []order.Order{
		{
			Symbol:  symbol,
			ID: "ksdhgkjsdhfg",
			Side: "Sell",
			Size:  10 ,
			Price: 120.32,
		},
	}
	render.JSON(w, r, trades) // A chi router helper for serializing and returning json
}

