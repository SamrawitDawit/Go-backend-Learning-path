package infrastructure

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashedPassword, nil
}

func CheckPassword(hashedPassword, plain_text string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plain_text))
}
