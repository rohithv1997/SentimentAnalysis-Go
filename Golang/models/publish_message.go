package models

type OutgoingMessage struct {
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}
