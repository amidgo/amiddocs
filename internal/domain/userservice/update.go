package userservice

import (
	"context"

	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (u *userService) UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) error {
	_, err := u.userprov.UserById(ctx, userId)
	if err != nil {
		return err
	}
	_, err = u.userprov.UserByLogin(ctx, login)
	if !amiderrors.Is(err, usererrorutils.NOT_FOUND) {
		return usererrorutils.LOGIN_ALREADY_EXIST
	}
	return u.usrrep.UpdateLogin(ctx, userId, login)
}

func (u *userService) UpdatePassword(ctx context.Context, userId uint64, password userfields.Password) error {
	_, err := u.userprov.UserById(ctx, userId)
	if err == nil {
		return err
	}
	hashPassword, err := u.encrypter.Hash(string(password))
	if err != nil {
		return err
	}
	return u.usrrep.UpdatePassword(ctx, userId, hashPassword)
}

func (u *userService) UpdateName(ctx context.Context, userId uint64, name userfields.Name) error {
	_, err := u.userprov.UserById(ctx, userId)
	if err != nil {
		return err
	}
	return u.usrrep.UpdateName(ctx, userId, name)
}

func (u *userService) UpdateSurname(ctx context.Context, userId uint64, surname userfields.Surname) error {
	_, err := u.userprov.UserById(ctx, userId)
	if err != nil {
		return err
	}
	return u.usrrep.UpdateSurname(ctx, userId, surname)
}

func (u *userService) UpdateFatherName(ctx context.Context, userId uint64, fatherName userfields.FatherName) error {
	_, err := u.userprov.UserById(ctx, userId)
	if err != nil {
		return err
	}
	return u.usrrep.UpdateFatherName(ctx, userId, fatherName)
}

func (u *userService) UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) error {
	_, err := u.userprov.UserById(ctx, userId)
	if err != nil {
		return err
	}
	_, err = u.userprov.UserByEmail(ctx, email)
	if !amiderrors.Is(err, usererrorutils.NOT_FOUND) {
		return usererrorutils.EMAIL_ALREADY_EXIST
	}
	return u.userprov.UpdateEmail(ctx, userId, email)
}
