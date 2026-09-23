package service

import (
	"context"

	"github.com/kodacampmain/koda-b9-gin/internal/dto"
	"github.com/kodacampmain/koda-b9-gin/internal/repo"
)

type PersonService struct {
	pr *repo.PersonRepo
}

func NewPersonService(pr *repo.PersonRepo) *PersonService {
	return &PersonService{
		pr: pr,
	}
}

func (p *PersonService) GetAllPerson(ctx context.Context) ([]dto.Person, error) {
	result, err := p.pr.GetAllPerson(ctx)
	data := make([]dto.Person, 0, len(result))
	for _, v := range result {
		data = append(data, dto.Person{
			Name: v.Name,
			Age:  v.Age,
		})
	}
	return data, err
}
