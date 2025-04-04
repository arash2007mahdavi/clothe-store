package handlers

import (
	"net/http"
	"store/src/responses"

	"github.com/gin-gonic/gin"
)

type Helper struct {
}

func (h Helper) MainStore(c *gin.Context) {
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "Welcome To Store"))
}
