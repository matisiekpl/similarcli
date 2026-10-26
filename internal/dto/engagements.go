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
	SocialOrganic float64 `json:"SocialOrganic"`
	SocialPaid    float64 `json:"SocialPaid"`
	Mail          float64 `json:"Mail"`
	Referrals     float64 `json:"Referrals"`
	SearchOrganic float64 `json:"SearchOrganic"`
	SearchPaid    float64 `json:"SearchPaid"`
	Direct        float64 `json:"Direct"`
	GenAi         float64 `json:"GenAi"`
	Affiliate     float64 `json:"Affiliate"`
	DisplayAds    float64 `json:"DisplayAds"`
}
