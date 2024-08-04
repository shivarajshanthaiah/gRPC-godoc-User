package services

import "golang.org/x/crypto/bcrypt"

func (u *UserService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	Password := string(bytes)
	return Password, nil
}

func (u *UserService) CheckPassword(providedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
