package repositories

import (
	"api-karang-waru/models"
	"strings"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindAll(
		search string,
		page int,
		limit int,
		sortBy string,
		sortOrder string,
	) ([]models.User, error)
	FindByEmail(email string) (*models.User, error)
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
func (r *userRepository) FindAll(
	search string,
	page int,
	limit int,
	sortBy string,
	sortOrder string,
) ([]models.User, error) {
	var users []models.User

	offset := (page - 1) * limit
	query := r.db.Model(&models.User{})

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

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// *models.User, error mengembalikan satu data User dalam bentuk pointer, plus error
func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// tidak butuh id karena langsung method Save() cari priamry key di struct models.User
func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(user *models.User) error {
	return r.db.Delete(user).Error
}
