package src

type User struct {
	ID       int
	Username string
	IsPlayed bool
	Score    int
}

type Player struct {
	ID    int `json:"id"`
	Score int `json:"score"`
}

type Players []Player

type PlayerEndGameRequest struct {
	Players Players `json:"players"`
}

type SignupResponse struct {
	Status    string               `json:"status"`
	Timestamp string               `json:"timestamp,omitempty"`
	Message   string               `json:"message,omitempty"`
	Result    SignupResponseResult `json:"result,omitempty"`
}

type SignupResponseResult struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
