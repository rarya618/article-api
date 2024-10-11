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

// Converts YYYYMMDD to YYYY-MM-DD
func FormatDate(oldDateFormat string) string {
	newDate := ""
	// Sample for testing purposes
	newDate = "2016-09-22"

	return newDate
}

// Gets an article with the given ID
func getArticleHandler(c *gin.Context) {
	// Extract the ID from the URL path
	id := c.Param("id")

	// Fetch the article by ID
	article, exists, message := utils.GetArticleByID(current_articles, id)
	if !exists {
		if message == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		} else {
			// Message received indicates that the request is not valid
			c.JSON(http.StatusBadRequest, gin.H{"error": message})
		}
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

	// If error occurs, send error response and return
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": error.Error})
		return
	}

	// Add the new article to the slice.
	current_articles[idInt] = newArticle
	c.IndentedJSON(http.StatusCreated, newArticle)
}

// Responds with the tag as JSON.
func getTagHandler(c *gin.Context) {
	// Extract params from the URL path
	tagName := c.Param("tagName")
	date := c.Param("date")

	// If tagname is not provided
	if len(tagName) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Tag name not provided"})
		return
	}

	// If date string is not 8 characters long
	if len(date) != 8 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid date: should have exactly 8 characters"})
		return
	}

	// Format the date
	formattedDate := FormatDate(date)

	// Pass in the tag name and formatted date to get tag data
	tag := utils.GetTagData(current_articles, tagName, formattedDate)

	c.IndentedJSON(http.StatusOK, tag)
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
