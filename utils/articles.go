package utils

import (
	"strconv"

	"github.com/rarya618/article-api/dataTypes"
)

// Responds with an Article with the given id (sends a message in the case the request is not valid)
func GetArticleByID(current_articles map[int]dataTypes.Article, id string) (dataTypes.Article, bool, string) {
	idInt, err := strconv.Atoi(id)

	// If int conversion fails
	if err != nil {
		// Send message to main
		return dataTypes.Article{}, false, "Invalid Article ID: Article ID needs to be a number"
	}

	article, exists := current_articles[idInt]
	return article, exists, ""

}
