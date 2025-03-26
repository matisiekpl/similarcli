package dto

type Engagments struct {
	BounceRate   string `json:"BounceRate"`
	Month        string `json:"Month"`
	Year         string `json:"Year"`
	PagePerVisit string `json:"PagePerVisit"`
	Visits       string `json:"Visits"`
	TimeOnSite   string `json:"TimeOnSite"`
}

type TrafficSources struct {
	Social        float64 `json:"Social"`
	PaidReferrals float64 `json:"Paid Referrals"`
	Mail          float64 `json:"Mail"`
	Referrals     float64 `json:"Referrals"`
	Search        float64 `json:"Search"`
	Direct        float64 `json:"Direct"`
}
