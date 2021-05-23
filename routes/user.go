package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tuanden0/user_api/helpers"
)

func UserRoute(r *gin.Engine) *gin.Engine {

	// Load html file
	// This load all the template files in templates folder
	// once loaded and don't need to load it again anymore
	r.LoadHTMLGlob("templates/*")

	// Set User Routes
	r.GET("/", helpers.Index)
	r.GET("/article/view/:article_id", helpers.GetArticle)

	return r
}
