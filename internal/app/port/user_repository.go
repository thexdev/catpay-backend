package port

type UserRepository interface {
	Create(email, password, role string) error

	GetHashedPasswordByEmail(email string) (string, error)

	Exist(email string) error
}
