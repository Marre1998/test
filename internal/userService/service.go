package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(user User) (User, error) {
	return u.repo.CreateUser(user)
}
func (u *UserService) GetAllUsers() ([]User, error) {
	return u.repo.GetAllUsers()
}

func (u *UserService) UpdateUserByID(id uint, user User) (User, error) {
	return u.repo.UpdateUserByID(id, user)
}
func (u *UserService) DeleteUserByID(id uint) error {
	return u.repo.DeleteUserByID(id)
}
