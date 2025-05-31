package middlewares

import (
	"fmt"
	"net/http"
	"store/src/configs"
	"store/src/responses"
	"store/src/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authentication(cfg *configs.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)

	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader("Authorization")
		token := strings.Split(auth, " ")
		if auth == "" {
			err = fmt.Errorf("invalid Authorization")
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.MakeResponseWithError(false, http.StatusUnauthorized, err))
			return
		}
		c.Set("user_id", claimMap["user_id"])
		c.Set("first_name", claimMap["first_name"])
		c.Set("last_name", claimMap["last_name"])
		c.Set("username", claimMap["username"])
		c.Set("email", claimMap["email"])
		c.Set("mobileNumber", claimMap["mobileNumber"])
		c.Set("roles", claimMap["roles"])
		c.Set("exp", claimMap["exp"])	
		c.Next()
	}
}

func Authorization(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, responses.MakeNormalResponse(false, http.StatusForbidden, nil))
			return
		}
		rolesVal := c.Keys["roles"]
		fmt.Println(rolesVal)
		if rolesVal == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, responses.MakeNormalResponse(false, http.StatusForbidden, nil))
			return
		}
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}
		
		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, responses.MakeNormalResponse(false, http.StatusForbidden, nil))
	}
}