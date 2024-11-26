package model

type NewChatRequest struct {
	UserId  string
	Balance float64
}

type EndChatRequest struct {
	UserId string
	ChatId string
}

type ChatRequest struct {
	UserId   string
	Question string
}
