package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nii236/margin/pkg/positions"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/nii236/margin/pkg/market"
	"github.com/nii236/margin/pkg/server"
)

func main() {
	fmt.Println("Starting market service...")
	m := market.New()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go m.Run(ctx)
	rw := market.NewRandomWalk()

	rwChan, _ := rw.Subscribe(context.Background())
	go func(rwChan chan *market.Ticker) {
		for {
			select {
			case t := <-rwChan:
				fmt.Println(t.Price)
			}
		}
	}(rwChan)

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	positionsRepo, err := positions.NewSQLite(":memory:", true, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	r.Mount("/v1/api/positions", server.NewPositions(positionsRepo))
	fmt.Println("Starting position service...")
	http.ListenAndServe(":8080", r)
}
