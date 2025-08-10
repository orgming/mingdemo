package demo

type Service struct {
	repository *Repository
}

func NewService() *Service {
	return &Service{repository: NewRepository()}
}

func (s *Service) GetUsers() []UserModel {
	return s.repository.GetUserByIds(s.repository.GetUserIds())
}
