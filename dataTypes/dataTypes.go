package dataTypes

// Represents data about an article.
type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

// Represents data about a tag.
type TagData struct {
	Tag         string   `json:"id"`
	Count       float64  `json:"title"`
	Articles    []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}
