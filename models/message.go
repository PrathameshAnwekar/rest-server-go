package models

type Message struct {
	ID         int    `json:"id"`
	SenderID   int    `json:"senderId"`
	ReceiverID int    `json:"receiverId"`
	Content    string `json:"content"`
	Timestamp  int64  `json:"timestamp"`
}
