package models

type Message struct {
	Content   string    `json:"content"`
	Username  string    `json:"username"`
	Timestamp time.Time `json:"timestamp"`
}
