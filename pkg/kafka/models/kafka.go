package models

type Producer struct {
	Topic string `json:"topic"`
	Data  any    `json:"data"`
}
