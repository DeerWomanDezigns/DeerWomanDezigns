package daos

import (
	"github.com/Maybeenaught/deer-woman-dezigns/code/config"
	"github.com/Maybeenaught/deer-woman-dezigns/code/models"
)

// UserDAO persists user data in database
type UserDAO struct{}

// NewUserDAO creates a new UserDAO
func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

// Get does the actual query to database, if user with specified id is not found error is returned
func (dao *UserDAO) Get(id uint) (*models.User, error) {
	var user models.User

	// Query Database here...

	//user = models.User{
	//	Model: models.Model{ID: 1},
	//	FirstName: "Sample",
	//	LastName: "McSampleson",
	//	Address: "55555 Sample Drive, Sample, Alaska",
	//	Email: "Sample.McSampleson@gmail.com"}

	// if using Gorm:
	err := config.Config.DB.Where("id = ?", id).
		First(&user).
		Error

	return &user, err
}
