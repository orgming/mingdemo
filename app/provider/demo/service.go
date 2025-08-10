package demo

import "github.com/orgming/mingdemo/framework"

type Service struct {
	container framework.Container
}

func NewService(params ...any) (any, error) {
	container := params[0].(framework.Container)
	return &Service{
		container: container,
	}, nil
}

func (s *Service) GetAllStudent() []Student {
	return []Student{
		{
			ID:   1,
			Name: "Andy",
		},
		{
			ID:   2,
			Name: "Bob",
		},
	}
}
