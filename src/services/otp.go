package services

import (
	"fmt"
	"store/src/configs"
	"store/src/database"
	"store/src/loggers"
	"time"
)

type OtpService struct {
	logger loggers.Logger
	cfg *configs.Config
}

type OtpDto struct {
	Value string
	Used bool
}

func NewOtpService(cfg *configs.Config) *OtpService {
	logger := loggers.NewLogger(cfg)
	return &OtpService{
		logger: logger,
		cfg: cfg,
	}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	value := OtpDto{Value: otp, Used: false}
	res, err := database.Get[OtpDto](mobileNumber)
	if err == nil && !res.Used {
		return fmt.Errorf("otp exists")
	} else if err == nil && res.Used {
		return fmt.Errorf("otp used")
	}
	err = database.Set(mobileNumber, value, s.cfg.Otp.Expire * time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {
	res, err := database.Get[OtpDto](mobileNumber)
	if err != nil {
		return err
	} else if res.Used {
		return fmt.Errorf("otp used")
	} else if !res.Used && res.Value != otp {
		return fmt.Errorf("invalid otp")
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = database.Set(mobileNumber, res, s.cfg.Otp.Expire * time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}