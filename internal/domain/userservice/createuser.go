package userservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *userService) CreateUser(ctx context.Context, u *usermodel.CreateUserDTO) (*usermodel.UserDTO, error) {
	err := s.checkEmail(ctx, u.Email)
	if err != nil {
		return nil, err
	}
	login, password, err := s.generateLoginAndPassword(u)
	if err != nil {
		return nil, err
	}
	user := usermodel.NewUserDTO(0, login, password, u.Name, u.Surname, u.FatherName, u.Email, u.Roles)
	user, err = s.usrrep.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) checkEmail(ctx context.Context, email userfields.Email) error {
	if len(email) == 0 {
		return nil
	}
	_, err := s.userprov.UserByEmail(ctx, email)
	if !amiderrors.Is(err, usererror.NOT_FOUND) {
		return usererror.EMAIL_ALREADY_EXIST
	}
	return nil
}

func (s *userService) generateLoginAndPassword(user *usermodel.CreateUserDTO) (userfields.Login, userfields.Password, error) {
	l, p, err := user.GenerateLoginAndPassword()
	if err != nil {
		return "", "", err
	}
	hash, err := s.encrypter.Hash(string(p))
	if err != nil {
		return "", "", err
	}
	login, password := l, userfields.Password(hash)
	return login, password, nil
}
