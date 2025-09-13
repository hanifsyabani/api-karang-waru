package services

import (
	"api-karang-waru/models"
	"api-karang-waru/repositories"
	"api-karang-waru/requests"
	"golang.org/x/crypto/bcrypt"
	"github.com/go-playground/validator/v10"
)

type UserService interface {
	CreateUser(req *requests.UserRequest)(*models.User, error)
	GetAllUser()([]models.User, error)
	GetUserByID(id uint)(*models.User, error)
	UpdateUser(id uint, req *requests.UserRequest)(*models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repository repositories.UserRepository
	validate *validator.Validate
}

func NewUserService(repository repositories.UserRepository) UserService{
	return &userService{
		repository: repository,
		validate: validator.New(),
	}
}

func (s *userService) CreateUser(req *requests.UserRequest) (*models.User,error) {
	if err := s.validate.Struct(req);err !=nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err !=nil {
		return nil, err
	}

	user := models.User{
		Name: req.Name,
		Email: req.Email,
		Password: string(hashedPassword),
	}

	err = s.repository.Create(&user)
	return &user, err
}

func (s *userService) GetAllUser() ([]models.User, error) {
	return s.repository.FindAll()
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repository.FindByID(id)
}

func (s *userService) UpdateUser(id uint, req *requests.UserRequest)(*models.User,error){
	if err := s.validate.Struct(req);err !=nil{
		return nil, err
	}

	user,err := s.repository.FindByID(id)
	if err != nil{
		return nil, err
	}

	user.Name = req.Name
	user.Email = req.Email
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