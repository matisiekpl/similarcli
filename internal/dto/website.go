package dto

type Website struct {
	Version                int            `json:"Version"`
	SiteName               string         `json:"SiteName"`
	Description            string         `json:"Description"`
	Title                  string         `json:"Title"`
	EstimatedMonthlyVisits map[string]int `json:"EstimatedMonthlyVisits"`
	TrafficSources         TrafficSources `json:"TrafficSources"`
	Category               string         `json:"Category"`
	SnapshotDate           string         `json:"SnapshotDate"`
}
