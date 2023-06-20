package userservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
)

func (u *userService) UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) error {
	return u.usrrep.UpdateLogin(ctx, userId, login)
}

func (u *userService) UpdatePassword(ctx context.Context, userId uint64, password userfields.Password) error {
	hashPassword, err := u.encrypter.Hash(string(password))
	if err != nil {
		return err
	}
	return u.usrrep.UpdatePassword(ctx, userId, hashPassword)
}

func (u *userService) UpdateName(ctx context.Context, userId uint64, name userfields.Name) error {
	return u.usrrep.UpdateName(ctx, userId, name)
}

func (u *userService) UpdateSurname(ctx context.Context, userId uint64, surname userfields.Surname) error {
	return u.usrrep.UpdateSurname(ctx, userId, surname)
}

func (u *userService) UpdateFatherName(ctx context.Context, userId uint64, fatherName userfields.FatherName) error {
	return u.usrrep.UpdateFatherName(ctx, userId, fatherName)
}

func (u *userService) UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) error {
	return u.userprov.UpdateEmail(ctx, userId, email)
}
