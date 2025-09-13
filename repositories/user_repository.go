package repositories

import (
	"api-karang-waru/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(user *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// []models.User artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek models.User.
func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err	
}

// *models.User, error mengembalikan satu data User dalam bentuk pointer, plus error
func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return  &user, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.User
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}


func (r *userRepository) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}