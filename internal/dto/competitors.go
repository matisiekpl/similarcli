package dto

type Competitors struct {
	TopSimilarityCompetitors []interface{} `json:"TopSimilarityCompetitors"`
}

type Notification struct {
	Content *string `json:"Content"`
}
