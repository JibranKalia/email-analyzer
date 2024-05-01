package models

type Email struct {
	ID         int    `json:"id"`
	Sender     string `json:"sender"`
	Recipients string `json:"recipients"`
	Subject    string `json:"subject"`
	Timestamp  string `json:"timestamp"`
}
