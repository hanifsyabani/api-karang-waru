package usecase

import (
	"api-karang-waru/internal/domain"
	
	"api-karang-waru/pkg/types"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	GetAllUser(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	)([]domain.User, error)
	GetUserByID(id uint)(*domain.User, error)
	UpdateUser(id uint, req *types.UserRequest)(*domain.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repository UserRepository
	validate *validator.Validate
}

func NewUserService(repository UserRepository) UserService{
	return &userService{
		repository: repository,
		validate: validator.New(),
	}
}


func (s *userService) GetAllUser(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.User, error) {
	return s.repository.FindAll(search, page, limit, sortBy, sortOrder)
}

func (s *userService) GetUserByID(id uint) (*domain.User, error) {
	return s.repository.FindByID(id)
}

func (s *userService) UpdateUser(id uint, req *types.UserRequest)(*domain.User,error){
	if err := s.validate.Struct(req);err !=nil{
		return nil, err
	}

	user,err := s.repository.FindByID(id)
	if err != nil{
		return nil, err
	}

	user.Name = req.Name
	user.Email = req.Email
	user.Role = req.Role
	user.Password = req.Password

	err = s.repository.Update(user)
	return user, err
}

func (s *userService) DeleteUser(id uint) error {
	contact, err := s.repository.FindByID(id)
	if err != nil{
		return err
	}

	return s.repository.Delete(contact)
}
