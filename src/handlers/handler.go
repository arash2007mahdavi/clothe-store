package handlers

import (
	"fmt"
	"net/http"
	"store/src/configs"
	"store/src/profiles"
	"store/src/responses"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var Amounts = map[string]int{
	"Hat": 10,
	"Shoes": 10,
	"Pant": 10,
	"Shirt": 10,
}

type Helper struct {
}

func (h Helper) MainStore(c *gin.Context) {
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "Welcome To Store"))
}

type Thing struct {
	Amount int
	Price float64
	Currency string
}

type Things struct {
	Hat Thing
	Shoes Thing
	Pant Thing
	Shirt Thing
}

func (h Helper) GetClothes(c *gin.Context) {
	cfg := configs.GetConfig()
	things := Things{
		Hat: Thing{Amount: Amounts["Hat"], Price: cfg.Store.Hat.Price, Currency: cfg.Store.Hat.Currency},
		Shoes: Thing{Amount: Amounts["Shoes"], Price: cfg.Store.Shoes.Price, Currency: cfg.Store.Shoes.Currency},
		Pant: Thing{Amount: Amounts["Pant"], Price: cfg.Store.Pant.Price, Currency: cfg.Store.Pant.Currency},
		Shirt: Thing{Amount: Amounts["Shirt"], Price: cfg.Store.Shirt.Price, Currency: cfg.Store.Shirt.Currency},
	}
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, things))
}

type GetProfile struct {
	ID       string `json:"id" binding:"required,id"`
	Password string `json:"password" binding:"required,password"`
	Fullname string `json:"fullname" binding:"required,min=10,max=25"`
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
		ID:       new.ID,
		Password: new.Password,
		Fullname: new.Fullname,
		Wallet:   0.0,
		Basket: profiles.Clothe{Hat: 0, Shoes: 0, Pant: 0, Shirt: 0},
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

func (h Helper) ProfileSeeAll(c *gin.Context) {
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, profiles.Profiles))
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

func (h Helper) BuyClothe(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")
	target := c.Query("target")
	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("invalid amount")))
		return
	}
	for x := range profiles.Profiles {
		if id == profiles.Profiles[x].ID {
			if password == profiles.Profiles[x].Password {
				targets := []string{"hat", "shoes", "pant", "shirt"}
				target = strings.ToLower(target)
				ok := contains(targets, target)
				if !ok {
					c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("invalid target")))
					return
				}
				switch target {
				case "hat":
					all_price := configs.GetConfig().Store.Hat.Price * float64(amount)
					fmt.Println(profiles.Profiles[x].Wallet)
					fmt.Println(all_price)
					if all_price > profiles.Profiles[x].Wallet {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("money is not enoght")))
						return
					}
					if Amounts["Hat"] <= 0 {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("hat is over")))
						return
					}
					Amounts["Hat"] = Amounts["Hat"]-amount
					profiles.Profiles[x].Wallet -= all_price
					profiles.Profiles[x].Basket.Hat += amount
					c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "bought successfuly"))
					return
				case "shoes":
					all_price := configs.GetConfig().Store.Shoes.Price * float64(amount)
					if all_price > profiles.Profiles[x].Wallet {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("money is not enoght")))
						return
					}
					if Amounts["Shoes"] <= 0 {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("shoes is over")))
						return
					}
					Amounts["Shoes"] = Amounts["Shoes"]-amount
					profiles.Profiles[x].Wallet -= all_price
					profiles.Profiles[x].Basket.Shoes += amount
					c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "bought successfuly"))
					return
				case "pant":
					all_price := configs.GetConfig().Store.Pant.Price * float64(amount)
					if all_price > profiles.Profiles[x].Wallet {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("money is not enoght")))
						return
					}
					if Amounts["Pant"] <= 0 {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("pant is over")))
						return
					}
					Amounts["Pant"] = Amounts["Pant"]-amount
					profiles.Profiles[x].Wallet -= all_price
					profiles.Profiles[x].Basket.Pant += amount
					c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "bought successfuly"))
					return
				case "shirt":
					all_price := configs.GetConfig().Store.Shirt.Price * float64(amount)
					if all_price > profiles.Profiles[x].Wallet {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("money is not enoght")))
						return
					}
					if Amounts["Shirt"] <= 0 {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("shirt is over")))
						return
					}
					Amounts["Shirt"] = Amounts["Shirt"]-amount
					profiles.Profiles[x].Wallet -= all_price
					profiles.Profiles[x].Basket.Shirt += amount
					c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "bought successfuly"))
					return
				}
			} else {
				c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("wrong password")))
				return
			}
		}
	}
	c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("this id doesnt exist")))
}
