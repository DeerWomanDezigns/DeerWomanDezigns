package services

import "github.com/deer-woman-dezigns/deer-woman-dezigns/code/models"

type userDAO interface {
	Get(id string) (*models.User, error)
	GetAll() (*[]models.User, error)
	Add(newUser models.User) (*models.User, error)
}

type UserService struct {
	dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (s *UserService) Get(id string) (*models.User, error) {
	return s.dao.Get(id)
}

func (s *UserService) GetAll() (*[]models.User, error) {
	return s.dao.GetAll()
}

func (s *UserService) Add(newUser models.User) (*models.User, error) {
	return s.dao.Add(newUser)
}
