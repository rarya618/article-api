package utils

import (
	"strconv"

	"github.com/rarya618/article-api/dataTypes"
)

// Responds with an Article with the given id.
func GetArticleByID(current_articles map[int]dataTypes.Article, id string) (dataTypes.Article, bool) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return dataTypes.Article{}, false
	}

	article, exists := current_articles[idInt]
	return article, exists

}
