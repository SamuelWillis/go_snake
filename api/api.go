package api

import(
	"encoding/json"
	"log"
	"net/http"
	"github.com/SamuelWillis/go_snake/helpers"
	"github.com/SamuelWillis/go_snake/game"
)

// Index endpont handler.
func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Yew!"))
}

// Start endpoint handler.
func Start(res http.ResponseWriter, req *http.Request) {
	state := game.State{}
	err := DecodeSnakeRequest(req, &state)

	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	helpers.Dump("decoded board state", state)

	respond(res, StartResponse{
		HeadType: "bendr",
		TailType: "round-bum",
	})
}

// Move endpoint handler.
func Move(res http.ResponseWriter, req *http.Request) {
	state := game.State{}
	err := DecodeSnakeRequest(req, &state)

	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	helpers.Dump("decoded board state", state)
	move := game.GetNextMove(state)
	helpers.Dump("next move", move)

	respond(res, MoveResponse{
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

// DecodeSnakeRequest function.
func DecodeSnakeRequest(req *http.Request, state *game.State) error {
	err := json.NewDecoder(req.Body).Decode(&state)
	return err
}
