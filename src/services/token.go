package services

import (
	"fmt"
	"store/src/configs"
	"store/src/loggers"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	logger loggers.Logger
	cfg    *configs.Config
}

type TokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	MobileNumber string
	Email        string
	Roles        []string
}

func NewTokenService(cfg *configs.Config) *TokenService {
	logger := loggers.NewLogger(cfg)
	return &TokenService{
		cfg:    cfg,
		logger: logger,
	}
}

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int    `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int    `json:"refreshTokenExpireTime"`
}

func (s *TokenService) GenerateToken(token *TokenDto) (*TokenDetail, error) {
	accessToken := &TokenDetail{}
	accessToken.AccessTokenExpireTime = int(time.Now().Add(s.cfg.Jwt.AccessTokenExpireDuration * time.Minute).Unix())
	accessToken.RefreshTokenExpireTime = int(time.Now().Add(s.cfg.Jwt.RefreshTokenExpireDuration * time.Minute).Unix())

	atc := jwt.MapClaims{}
	atc["user_id"] = token.UserId
	atc["first_name"] = token.FirstName
	atc["last_name"] = token.LastName
	atc["username"] = token.Username
	atc["email"] = token.Email
	atc["mobileNumber"] = token.MobileNumber
	atc["roles"] = token.Roles
	atc["exp"] = accessToken.AccessTokenExpireTime

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atc)
	var err error
	accessToken.AccessToken, err = at.SignedString([]byte(s.cfg.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	rtc := jwt.MapClaims{}
	rtc["user_id"] = token.UserId
	rtc["exp"] = accessToken.RefreshTokenExpireTime

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtc)
	accessToken.RefreshToken, err = rt.SignedString([]byte(s.cfg.Jwt.Secret))
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("error in verify token")
		}
		return []byte(s.cfg.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return at, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = map[string]interface{}{}

	verifyToken , err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}
		return claimMap, nil
	}
	return nil, fmt.Errorf("claims not found")
}