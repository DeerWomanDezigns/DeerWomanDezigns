package daos

import (
	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/config"
	"github.com/deer-woman-dezigns/deer-woman-dezigns/code/models"
)

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	return &UserDAO{}
}

func (dao *UserDAO) Get(id string) (*models.User, error) {
	var user models.User

	table := config.Config.DB.Table("Users")
	err := table.Get("user_id", id).One(&user)

	return &user, err
}

func (dao *UserDAO) GetAll() (*[]models.User, error) {
	var users []models.User

	table := config.Config.DB.Table("Users")
	err := table.Scan().All(&users)

	return &users, err
}

func (dao *UserDAO) Add(newUser models.User) (*models.User, error) {
	var user *models.User

	table := config.Config.DB.Table("Users")
	err := table.Put(newUser).Run()
	if err == nil {
		user, err = dao.Get(newUser.ID)
	}

	return user, err
}
