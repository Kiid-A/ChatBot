package model

type NewChatRequest struct {
	UserName string 
	Balance  float64
}

type EndChatRequest struct {
	UserId
}

type ChatRequest struct {

}
