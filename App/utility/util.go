package utility

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(password string) string {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		panic(err)
	}

	return string(encrypt)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
