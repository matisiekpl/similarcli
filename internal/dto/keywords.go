package dto

type TopKeyword struct {
	Name           string   `json:"Name"`
	EstimatedValue int      `json:"EstimatedValue"`
	Volume         int      `json:"Volume"`
	Cpc            *float64 `json:"Cpc"`
}
