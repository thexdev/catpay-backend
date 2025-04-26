package service

type PasswordHasher interface {
	Make(plain string) (string, error)

	Verify(plain string, hashed string) bool
}
