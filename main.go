package main

import (
	"strconv"

	"github.com/rarya618/article-api/dataTypes"
	"github.com/rarya618/article-api/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Current Articles variable
var current_articles = map[int]dataTypes.Article{
	1: {
		ID:    "1",
		Title: "latest science shows that potato chips are better for you than sugar",
		Date:  "2016-09-22",
		Body:  "some text, potentially containing simple markup about how potato chips are great",
		Tags:  []string{"health", "fitness", "science"},
	},
}

func getArticleHandler(c *gin.Context) {
	// Extract the ID from the URL path
	id := c.Param("id")

	// Fetch the article by ID
	article, exists := utils.GetArticleByID(current_articles, id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Return the article as JSON
	c.JSON(http.StatusOK, article)
}

// Post an article to the API from JSON in the request body.
func postArticleHandler(c *gin.Context) {
	var newArticle dataTypes.Article

	// Call BindJSON to bind the received JSON to
	// newArticle.
	if err := c.BindJSON(&newArticle); err != nil {
		return
	}

	// Get int from string
	idInt, err := strconv.Atoi(newArticle.ID)

	// If error occurs, return
	if err != nil {
		return
	}

	// Add the new article to the slice.
	current_articles[idInt] = newArticle
	c.IndentedJSON(http.StatusCreated, newArticle)
}

// Responds with the list of all tags as JSON.
func getTagHandler(c *gin.Context) {
	// Extract params from the URL path
	tagName := c.Param("tagName")
	date := c.Param("date")

	tags := utils.GetTagData(current_articles, tagName, date)

	c.IndentedJSON(http.StatusOK, tags)
}

// The main function
func main() {
	// Set up router
	router := gin.Default()

	// Gets article with a specific id from the API
	router.GET("/articles/:id", getArticleHandler)

	// Posts articles to the API
	router.POST("/articles", postArticleHandler)

	// Gets tag with a specific tag name from the API
	router.GET("/tags/:tagName/:date", getTagHandler)

	// Runs the server
	router.Run("localhost:8080")
}
