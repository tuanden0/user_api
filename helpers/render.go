package helpers

import "github.com/gin-gonic/gin"

func Render(c *gin.Context, status int, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(status, data["payload"])
	case "application/xml":
		c.XML(status, data["payload"])
	default:
		c.HTML(status, templateName, data)
	}
}
