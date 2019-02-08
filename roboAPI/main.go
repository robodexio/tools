package main

import (
	"./db"
	"./order"
	"./orderBook"
	"./trades"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)


func Routes() (*chi.Mux, *db.Db) {

	dataBase, err := db.New(
		db.ConnString("51.143.167.229", 5432, "postgres", "stg9bZuxUc9vKJ9CqeYGc", "postgres"),
	)

	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()
	//throttle := middleware.Throttle(100)
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		middleware.Logger,                             // Log API request calls
		middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		middleware.Recoverer,                          // Recover from panics without crashing server
		middleware.Timeout(2500 * time.Millisecond),   // Stop processing after 2.5 seconds
	)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/order", order.Routes(dataBase))
		r.Mount("/api/trades", trades.Routes())
		r.Mount("/api/orderBook", orderBook.Routes(dataBase))
	})

	return router, dataBase
}


func main() {
	router, DataBase := Routes()
	defer DataBase.Close()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Fatal(http.ListenAndServe(":8080", router)) // Note, the port is usually gotten from the environment.

}