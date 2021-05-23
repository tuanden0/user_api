package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuanden0/user_api/models"
)

// Function call by "/"
func Index(c *gin.Context) {

	articles, _ := models.Article{}.All()

	// Render HTML template
	Render(
		c,
		http.StatusOK,
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
		"index.html",
	)
}

// Function call by /article/view/:article_id
func GetArticle(c *gin.Context) {

	article, _ := models.Article{}.Get(c.Param("article_id"))

	// Render HTML template
	Render(
		c,
		http.StatusOK,
		gin.H{
			"title":   "Home Page",
			"payload": article,
		},
		"article.html",
	)
}
