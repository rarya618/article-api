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

// Represents data about a tag.
type TagData struct {
	Tag         string   `json:"id"`
	Count       float64  `json:"title"`
	Articles    []string `json:"articles"`
	RelatedTags []string `json:"related_tags"`
}

// Articles variable
var articles = map[string]Article{
	"1": {
		ID:    "1",
		Title: "latest science shows that potato chips are better for you than sugar",
		Date:  "2016-09-22",
		Body:  "some text, potentially containing simple markup about how potato chips are great",
		Tags:  []string{"health", "fitness", "science"},
	},
}

// Tags variable
var tags = []TagData{
	{
		Tag:         "health",
		Count:       1,
		Articles:    []string{"1"},
		RelatedTags: []string{"science", "fitness"},
	},
}

// Responds with an Article with the given id.
func getArticleByID(id string) (Article, bool) {
	article, exists := articles[id]
	return article, exists
}

func getArticleHandler(c *gin.Context) {
	// Extract the ID from the URL path
	id := c.Param("id")

	// Fetch the article by ID
	article, exists := getArticleByID(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return
	}

	// Return the article as JSON
	c.JSON(http.StatusOK, article)
}

// Post an article to the API from JSON in the request body.
func postArticleHandler(c *gin.Context) {
	var newArticle Article

	// Call BindJSON to bind the received JSON to
	// newArticle.
	if err := c.BindJSON(&newArticle); err != nil {
		return
	}

	// Add the new article to the slice.
	articles[newArticle.ID] = newArticle
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

	// Gets article with a specific id from the API
	router.GET("/articles/:id", getArticleHandler)

	// Posts articles to the API
	router.POST("/articles", postArticleHandler)

	// Gets tag with a specific tag name from the API
	router.GET("/tags/:tagName/:date", getTags)

	// Runs the server
	router.Run("localhost:8080")
}
