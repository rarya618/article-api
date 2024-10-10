package utils

import (
	"github.com/rarya618/article-api/dataTypes"
)

// Removes duplicates and the original tag from a slice of tags
func processTags(tags []string, tagName string) []string {
	// Create a map to store unique elements
	seen := make(map[string]bool)
	result := []string{}

	// Loop through the slice, adding elements to the map if they haven't been seen before
	for _, val := range tags {
		if _, ok := seen[val]; !ok {
			seen[val] = true
			if val != tagName {
				result = append(result, val)
			}
		}
	}
	return result
}

// Gets tag data from a list of current articles
func GetTagData(current_articles map[int]dataTypes.Article, tagName string, date string) dataTypes.TagData {
	// Initialise tag data
	tagData := dataTypes.TagData{}
	tagData.Tag = tagName
	tagData.Count = 0
	tagData.Articles = []string{}

	// Loop through current articles (start from 1 as first element has id 1)
	for i := 1; i <= len(current_articles); i++ {
		// Get current article
		currentArticle := current_articles[i]

		// Loop through tags in current article
		for j := 0; j < len(currentArticle.Tags); j++ {
			// Initialise current tag
			currentTag := currentArticle.Tags[j]

			// If tag name matches
			if tagName == currentTag {
				// If date matches
				if date == currentArticle.Date {
					// Append Article ID to tag Articles list
					tagData.Articles = append(tagData.Articles, currentArticle.ID)

					// Append related tags (currently appends all tags without checking for duplicates)
					tagData.RelatedTags = append(tagData.RelatedTags, currentArticle.Tags...)

					// Increment count
					tagData.Count += 1
				}
			}
		}
	}

	// Remove duplicates and current tag from the related-tags
	tagData.RelatedTags = processTags(tagData.RelatedTags, tagName)

	return tagData
}
