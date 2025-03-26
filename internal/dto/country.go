package dto

type Country struct {
	Code    string `json:"Code"`
	UrlCode string `json:"UrlCode"`
	Name    string `json:"Name"`
}

type TopCountryShare struct {
	Country     int     `json:"Country"`
	CountryCode string  `json:"CountryCode"`
	Value       float64 `json:"Value"`
}
