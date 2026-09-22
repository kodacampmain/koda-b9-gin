package service

import (
	"errors"

	"github.com/kodacampmain/koda-b9-gin/internal/dto"
)

type PingServiceNew struct{}

func NewPingServiceNew() *PingServiceNew {
	return &PingServiceNew{}
}

func (p *PingServiceNew) EmptyValidation(data dto.User) error {
	if data.Name == "" && data.Age == 0 {
		return errors.New("body kosong")
	}
	return nil
}
