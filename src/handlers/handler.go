package handlers

import (
	"fmt"
	"net/http"
	"store/src/configs"
	"store/src/profiles"
	"store/src/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Helper struct {
}

func (h Helper) MainStore(c *gin.Context) {
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "Welcome To Store"))
}

func (h Helper) GetClothes(c *gin.Context) {
	cfg := configs.GetConfig()
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, cfg.Store))
}

type GetProfile struct {
	ID       string  `json:"id" binding:"required,id"`
	Password string  `json:"password" binding:"required,password"`
	Fullname string  `json:"fullname" binding:"required,min=10,max=25"`
}

func (h Helper) ProfileNew(c *gin.Context) {
	new := GetProfile{}
	err := c.ShouldBindJSON(&new)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadGateway, responses.MakeResponseWithValidationError(false, http.StatusBadGateway, err))
		return
	}
	for _, profile := range profiles.Profiles {
		if new.ID == profile.ID {
			c.AbortWithStatusJSON(http.StatusBadGateway, responses.MakeResponseWithError(false, http.StatusBadGateway, fmt.Errorf("this id used by someone")))
			return
		}
	}
	new2 := profiles.Profile{
		ID: new.ID,
		Password: new.Password,
		Fullname: new.Fullname,
		Wallet: 0.0,
	}
	profiles.AddProfile(new2)
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "new profile created"))
}

func (h Helper) ProfileSee(c *gin.Context) {
	id := c.Query("id")
	for _, prof := range profiles.Profiles {
		if prof.ID == id {
			password := c.Query("password")
			if prof.Password == password {
				c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, prof))
				return
			}
			c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("wrong password")))
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusBadGateway, responses.MakeResponseWithError(false, http.StatusBadGateway, fmt.Errorf("this id doesnt exist")))
}

func (h Helper) ProfileChargeWallet(c *gin.Context) {
	amount := c.Query("amount")
	amount1, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.MakeResponseWithError(false, http.StatusBadRequest, fmt.Errorf("invalid amount for charging")))
		return
	}
	id := c.Query("id")
	for i := range profiles.Profiles {
        if id == profiles.Profiles[i].ID {
            profiles.Profiles[i].Wallet += amount1
            c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "charged successfully"))
            return
        }
    }
	c.AbortWithStatusJSON(http.StatusBadGateway, responses.MakeResponseWithError(false, http.StatusBadGateway, fmt.Errorf("this id doesnt exist")))
}