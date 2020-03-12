package api

// StartResponse type
type StartResponse struct {
	Color string `json:"color,omitempty"`
	Shout string `json:"shout:omitempty"`
	HeadType string `json:"headType:omitempty"`
	TailType string `json:"tailType:omitempty"`
}

// MoveResponse type
type MoveResponse struct {
	Move string `json:"move"`
}
