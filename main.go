package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Represents data about an article.
type Article struct {
	ID    string   `json:"id"`
	Title string   `json:"title"`
	Date  string   `json:"date"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

type TagData struct {
	Tag         string   `json:"id"`
	Count       float64  `json:"title"`
	Articles    []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

var articles = []Article{
	{
		ID:    "1",
		Title: "latest science shows that potato chips are better for you than sugar",
		Date:  "2016-09-22",
		Body:  "some text, potentially containing simple markup about how potato chips are great",
		Tags:  []string{"health", "fitness", "science"},
	},
}

var tags = []TagData{}

// Responds with the list of all articles as JSON.
func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}

// Adds an article from JSON received in the request body.
func postArticles(c *gin.Context) {
	var newArticle Article

	// Call BindJSON to bind the received JSON to
	// newArticle.
	if err := c.BindJSON(&newArticle); err != nil {
		return
	}

	// Add the new article to the slice.
	articles = append(articles, newArticle)
	c.IndentedJSON(http.StatusCreated, newArticle)
}

// Responds with the list of all tags as JSON.
func getTags(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, tags)
}

// The main function
func main() {
	// Set up router
	router := gin.Default()

	// Gets articles from the API
	router.GET("/articles", getArticles)

	// Posts articles to the API
	router.POST("/articles", postArticles)

	// Gets tags from the API
	router.GET("/tags", getTags)

	// Runs the server
	router.Run("localhost:8080")
}
