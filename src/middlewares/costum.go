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
	c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("Wrong api-key")))
}