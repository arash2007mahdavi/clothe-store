package handlers

import (
	"fmt"
	"net/http"
	"store/src/configs"
	"store/src/profiles"
	"store/src/responses"
	"strings"

	"github.com/gin-gonic/gin"
)

var Amounts = map[string]int{
	"Hat":   10,
	"Shoes": 10,
	"Pant":  10,
	"Shirt": 10,
}

type Helper struct {
}

// @Summary welcome to store
// @Description this welcome message is for testing api
// @Tags Welcome
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router / [get]
func (h Helper) MainStore(c *gin.Context) {
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "Welcome To Store"))
}

type Thing struct {
	Amount   int
	Price    float64
	Currency string
}

type Things struct {
	Hat   Thing
	Shoes Thing
	Pant  Thing
	Shirt Thing
}

// @Summary get clothes
// @Description get clothes information (price and amount)
// @Tags Clothes
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /clothes [get]
func (h Helper) GetClothes(c *gin.Context) {
	cfg := configs.GetConfig()
	things := Things{
		Hat:   Thing{Amount: Amounts["Hat"], Price: cfg.Store.Hat.Price, Currency: cfg.Store.Hat.Currency},
		Shoes: Thing{Amount: Amounts["Shoes"], Price: cfg.Store.Shoes.Price, Currency: cfg.Store.Shoes.Currency},
		Pant:  Thing{Amount: Amounts["Pant"], Price: cfg.Store.Pant.Price, Currency: cfg.Store.Pant.Currency},
		Shirt: Thing{Amount: Amounts["Shirt"], Price: cfg.Store.Shirt.Price, Currency: cfg.Store.Shirt.Currency},
	}
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, things))
}

type GetProfile struct {
	ID       string `json:"id" binding:"required,id"`
	Password string `json:"password" binding:"required,password"`
	Fullname string `json:"fullname" binding:"required,min=10,max=25"`
}

// @Summary create new profile
// @Description create new profile
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /profile/new [post]
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
		Basket:   profiles.Clothe{Hat: 0, Shoes: 0, Pant: 0, Shirt: 0},
	}
	profiles.AddProfile(new2)
	c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "new profile created"))
}

type ProfileSeeAccess struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary watch profile
// @Description watch your profile
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /profile/see [get]
func (h Helper) ProfileSee(c *gin.Context) {
	p := ProfileSeeAccess{}
	c.ShouldBindJSON(&p)
	for _, prof := range profiles.Profiles {
		if prof.ID == p.ID {
			if prof.Password == p.Password {
				c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, prof))
				return
			}
			c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("wrong password")))
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusBadGateway, responses.MakeResponseWithError(false, http.StatusBadGateway, fmt.Errorf("this id doesnt exist")))
}

type ProfileChargeWalletAccess struct {
	ID     string  `json:"id" binding:"required"`
	Amount float64 `json:"amount" binding:"required,numeric"`
}

// @Summary charge wallet
// @Description charge wallet
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /profile/charge/wallet [post]
func (h Helper) ProfileChargeWallet(c *gin.Context) {
	p := ProfileChargeWalletAccess{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, responses.MakeResponseWithValidationError(false, http.StatusBadRequest, err))
		return
	}
	for i := range profiles.Profiles {
		if p.ID == profiles.Profiles[i].ID {
			profiles.Profiles[i].Wallet += p.Amount
			c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "charged successfully"))
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusBadGateway, responses.MakeResponseWithError(false, http.StatusBadGateway, fmt.Errorf("this id doesnt exist")))
}

// @Summary watch profiles
// @Description watch all of profiles
// @Tags Profile
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /profile/see/all [get]
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

type BuyClotheAccess struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Target   string `json:"target" binding:"required"`
	Amount   int    `json:"amount" binding:"required,numeric"`
}

// @Summary buy clothe
// @Description buy clothe
// @Tags Clothes
// @Accept json
// @Produce json
// @Success 200 {object} responses.Response "Success"
// @Router /clothes/buy [post]
func (h Helper) BuyClothe(c *gin.Context) {
	p := BuyClotheAccess{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithValidationError(false, http.StatusLocked, err))
		return
	}
	for x := range profiles.Profiles {
		if p.ID == profiles.Profiles[x].ID {
			if p.Password == profiles.Profiles[x].Password {
				targets := []string{"hat", "shoes", "pant", "shirt"}
				p.Target = strings.ToLower(p.Target)
				ok := contains(targets, p.Target)
				if !ok {
					c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("invalid target")))
					return
				}
				switch p.Target {
				case "hat":
					all_price := configs.GetConfig().Store.Hat.Price * float64(p.Amount)
					if all_price > profiles.Profiles[x].Wallet {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("money is not enoght")))
						return
					}
					if Amounts["Hat"] < p.Amount {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("hat is over")))
						return
					}
					Amounts["Hat"] = Amounts["Hat"] - p.Amount
					profiles.Profiles[x].Wallet -= all_price
					profiles.Profiles[x].Basket.Hat += p.Amount
					c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "bought successfuly"))
					return
				case "shoes":
					all_price := configs.GetConfig().Store.Shoes.Price * float64(p.Amount)
					if all_price > profiles.Profiles[x].Wallet {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("money is not enoght")))
						return
					}
					if Amounts["Shoes"] <= p.Amount {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("shoes is over")))
						return
					}
					Amounts["Shoes"] = Amounts["Shoes"] - p.Amount
					profiles.Profiles[x].Wallet -= all_price
					profiles.Profiles[x].Basket.Shoes += p.Amount
					c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "bought successfuly"))
					return
				case "pant":
					all_price := configs.GetConfig().Store.Pant.Price * float64(p.Amount)
					if all_price > profiles.Profiles[x].Wallet {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("money is not enoght")))
						return
					}
					if Amounts["Pant"] <= p.Amount {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("pant is over")))
						return
					}
					Amounts["Pant"] = Amounts["Pant"] - p.Amount
					profiles.Profiles[x].Wallet -= all_price
					profiles.Profiles[x].Basket.Pant += p.Amount
					c.JSON(http.StatusOK, responses.MakeNormalResponse(true, http.StatusOK, "bought successfuly"))
					return
				case "shirt":
					all_price := configs.GetConfig().Store.Shirt.Price * float64(p.Amount)
					if all_price > profiles.Profiles[x].Wallet {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("money is not enoght")))
						return
					}
					if Amounts["Shirt"] <= p.Amount {
						c.AbortWithStatusJSON(http.StatusLocked, responses.MakeResponseWithError(false, http.StatusLocked, fmt.Errorf("shirt is over")))
						return
					}
					Amounts["Shirt"] = Amounts["Shirt"] - p.Amount
					profiles.Profiles[x].Wallet -= all_price
					profiles.Profiles[x].Basket.Shirt += p.Amount
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
