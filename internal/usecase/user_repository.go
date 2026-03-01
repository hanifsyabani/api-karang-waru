package usecase

import (
	"api-karang-waru/internal/domain"
	"strings"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	FindAll(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindByID(id uint) (*domain.User, error)
	Update(user *domain.User) error
	Delete(user *domain.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// []domain.User artinya fungsi tersebut mengembalikan sebuah slice (mirip array tapi lebih fleksibel di Go) yang berisi banyak objek domain.User.
func (r *userRepository) FindAll(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]domain.User, error) {
	var users []domain.User

	offset := (page - 1) * limit
	query := r.db.Model(&domain.User{})

	// WHERE username LIKE '%search%' OR email LIKE '%search%'
	// ? → placeholder (aman dari SQL injection)
	// %search% → pencarian partial (mengandung kata)
	if search != "" {
		search = strings.ToLower(search)
		query = query.Where(
			"LOWER(full_name) LIKE ? OR LOWER(email) LIKE ?",
			"%"+search+"%",
			"%"+search+"%",
		)
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "desc"
	}

	// sortBy → nama kolom
	// sortOrder → arah sorting
	query = query.Order(sortBy + " " + sortOrder)

	err := query.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// *domain.User, error mengembalikan satu data User dalam bentuk pointer, plus error
func (r *userRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct domain.User
func (r *userRepository) Update(user *domain.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(user *domain.User) error {
	return r.db.Delete(user).Error
}

