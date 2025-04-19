package commons

import (
	"math"
	"math/rand"
	"store/src/configs"
	"strconv"
	"time"
)

func MakeOtp() string {
	rand.Seed(time.Now().UTC().UnixNano())
	min := int(math.Pow(float64(10), (float64(configs.GetConfig().Otp.Digit)-1)))
	max := int(math.Pow(float64(10), (float64(configs.GetConfig().Otp.Digit)))) -1
	otp := rand.Intn(max-min) + min
	return strconv.Itoa(otp)
}