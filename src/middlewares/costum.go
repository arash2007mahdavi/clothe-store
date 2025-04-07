package middlewares

import (
	"fmt"
	"net/http"
	"store/src/responses"

	"github.com/gin-gonic/gin"
)

func CheckApiKey(c *gin.Context) {
	apikey := c.GetHeader("api-key")
	if apikey == "password" {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("wrong api-key")))
}

func CheckAdmin(c *gin.Context) {
	admin_id := c.GetHeader("admin-id")
	admin_password := c.GetHeader("admin-password")
	if admin_id == "arash2007mahdavi" && admin_password == "@rash2007" {
		c.Next()
		return
	}
	c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("just admin can see the informations")))
}
