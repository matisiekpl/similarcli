package dto

type Rank struct {
	Rank *int `json:"Rank"`
}

type CountryRank struct {
	Country     *int    `json:"Country"`
	CountryCode *string `json:"CountryCode"`
	Rank        *int    `json:"Rank"`
}

type CategoryRank struct {
	Rank     *int    `json:"Rank"`
	Category *string `json:"Category"`
}
