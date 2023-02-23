package models

type Transfer struct {
	Id          string `json:"id"`
	Amount      string `json:"amount"`
	IsProcessed bool   `json:"is_processed"`
}
