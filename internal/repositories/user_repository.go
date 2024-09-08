package repositories

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser() {
	// Create user
}