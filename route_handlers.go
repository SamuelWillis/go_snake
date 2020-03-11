package main

import (
	"log"
	"net/http"
	"github.com/SamuelWillis/go_snake/api"
	"github.com/SamuelWillis/go_snake/game"
)

// Index endpont handler.
func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Yew!"))
}

// Start endpoint handler.
func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)

	if err != nil {
		log.Printf("Bad start request: %v", err)
	}

	dump(decoded)

	respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

// Move endpoint handler.
func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)

	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	dump(decoded)

	move := game.GetNextMove(decoded)

	dump(move)

	respond(res, api.MoveResponse{
		Move: move,
	})
}

// End endpoint handler.
func End(res http.ResponseWriter, req *http.Request) {
	return
}

// Ping endpoint handler.
func Ping(res http.ResponseWriter, req *http.Request) {
	return
}
