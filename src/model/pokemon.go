package model

type Pokemon struct {
	Order  int    `json:"order"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
}
