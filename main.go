package main

import (
	"fmt"
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

// Returns characters from strings using index
// Reference: https://www.tutorialspoint.com/golang-program-to-get-characters-from-a-string-using-the-index
func GetChar(str string, index int) rune {
	return []rune(str)[index]
}

// Converts YYYYMMDD to YYYY-MM-DD, returns blank value and message if failed
func FormatDate(oldDateFormat string) (string, string) {
	year := oldDateFormat[:4]
	month := oldDateFormat[4:6]
	day := oldDateFormat[6:]

	// If year, month, day is not a valid integer, greater than 0, has 8 digits
	if yearInt, yearErr := strconv.Atoi(year); yearErr != nil || yearInt <= 0 {
		return "", "Invalid date: year invalid"
	} else if monthInt, monthErr := strconv.Atoi(month); monthErr != nil || monthInt <= 0 {
		return "", "Invalid date: month invalid"
	} else if dayInt, dayErr := strconv.Atoi(day); dayErr != nil || dayInt <= 0 {
		return "", "Invalid date: day invalid"
	}

	// If no issues, return normal date
	return fmt.Sprintf("%s-%s-%s", year, month, day), ""
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

	// Add the new article
	isSuccessful, message := utils.AddArticle(current_articles, idInt, newArticle)

	// If adding successful
	if isSuccessful {
		c.IndentedJSON(http.StatusCreated, newArticle)
	} else {
		// Throw error if adding fails
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": message})
	}

}

// Responds with the tag as JSON.
func getTagHandler(c *gin.Context) {
	// Extract params from the URL path
	tagName := c.Param("tagName")
	date := c.Param("date")

	// If tagName is not provided
	if len(tagName) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Tag name not provided"})
		return
	}

	// If date string is not 8 characters long
	if len(date) != 8 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid date: should have exactly 8 characters"})
		return
	}

	dateInt, dateErr := strconv.Atoi(date)

	// If date is not a valid integer, greater than 0, has 8 digits
	if dateErr != nil || dateInt <= 0 || dateInt > 10000000 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid date: should be a valid number"})
		return
	}

	// Format the date
	formattedDate, message := FormatDate(date)

	// If message exists, throw error 400
	if message != "" {
		c.IndentedJSON(http.StatusBadRequest, message)
	}

	// Pass in the tag name and formatted date to get tag data
	tag := utils.GetTagData(current_articles, tagName, formattedDate)

	// Send tag as response to client
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
