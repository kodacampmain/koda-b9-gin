package service

import (
	"errors"

	"github.com/kodacampmain/koda-b9-gin/internal/dto"
)

var Users = []dto.Account{}

type PingService struct{}

func NewPingService() *PingService {
	return &PingService{}
}

func (p *PingService) EmptyValidation(data dto.User) error {
	if data.Name == "" && data.Age == 0 {
		return errors.New("empty body")
	}
	return nil
}
