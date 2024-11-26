package model

type NewChatRequest struct {
	UserId  string `json:user_id`
	Balance float64	`json:balance`
}

type EndChatRequest struct {
	UserId string
	ChatId string
}

type ChatRequest struct {
	UserId   string
	ChatId   string
	Question string
}
