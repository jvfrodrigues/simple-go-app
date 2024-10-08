package model

type Route struct {
	From string `json:"from"`
	To   string `json:"to"`
	Cost int    `json:"cost"`
}
