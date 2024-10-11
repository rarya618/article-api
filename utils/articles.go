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

func AddArticle(current_articles map[int]dataTypes.Article, idInt int, newArticle dataTypes.Article) (bool, string) {
	// Check if Article ID is unique
	if current_articles[idInt].ID == newArticle.ID {
		return false, "Invalid Article ID: should have a unique Article ID"
	}

	// Add the new article to the slice.
	current_articles[idInt] = newArticle
	return true, ""
}
