package encrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type Encrypter struct {
	cost int
}

func New(cost int) *Encrypter {
	return &Encrypter{cost: cost}
}

func (e *Encrypter) Hash(input string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(input), e.cost)
	return string(b), err
}

func (e *Encrypter) Verify(hashPassword string, password string) bool {
	hash := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
