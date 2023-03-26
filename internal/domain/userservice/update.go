package userservice

import (
	"context"

	usererrorutils "github.com/amidgo/amiddocs/internal/errorutils/user_error_utils"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (u *UserService) UpdateLogin(ctx context.Context, userId uint64, login userfields.Login) *amiderrors.ErrorResponse {

	usr, _ := u.usrrep.GetUserById(ctx, userId)
	if usr == nil {
		return usererrorutils.USER_NOT_FOUND
	}
	usr, _ = u.usrrep.GetUserByLogin(ctx, login)
	if usr != nil {
		return usererrorutils.LOGIN_ALREADY_EXIST
	}
	return u.usrrep.UpdateLogin(ctx, userId, login)
}

func (u *UserService) UpdatePassword(ctx context.Context, userId uint64, password userfields.Password) *amiderrors.ErrorResponse {
	usr, _ := u.usrrep.GetUserById(ctx, userId)
	if usr == nil {
		return usererrorutils.USER_NOT_FOUND
	}
	hashPassword, err := password.Hash()
	if err != nil {
		return err
	}
	return u.usrrep.UpdatePassword(ctx, userId, hashPassword)
}

func (u *UserService) UpdateName(ctx context.Context, userId uint64, name userfields.Name) *amiderrors.ErrorResponse {
	usr, _ := u.usrrep.GetUserById(ctx, userId)
	if usr == nil {
		return usererrorutils.USER_NOT_FOUND
	}
	return u.usrrep.UpdateName(ctx, userId, name)
}

func (u *UserService) UpdateSurname(ctx context.Context, userId uint64, surname userfields.Surname) *amiderrors.ErrorResponse {
	usr, _ := u.usrrep.GetUserById(ctx, userId)
	if usr == nil {
		return usererrorutils.USER_NOT_FOUND
	}
	return u.usrrep.UpdateSurname(ctx, userId, surname)
}

func (u *UserService) UpdateFatherName(ctx context.Context, userId uint64, fatherName userfields.FatherName) *amiderrors.ErrorResponse {
	usr, _ := u.usrrep.GetUserById(ctx, userId)
	if usr == nil {
		return usererrorutils.USER_NOT_FOUND
	}
	return u.usrrep.UpdateFatherName(ctx, userId, fatherName)
}

func (u *UserService) UpdateEmail(ctx context.Context, userId uint64, email userfields.Email) *amiderrors.ErrorResponse {
	usr, _ := u.usrrep.GetUserById(ctx, userId)
	if usr == nil {
		return usererrorutils.USER_NOT_FOUND
	}
	usr, _ = u.usrrep.GetUserByEmail(ctx, email)
	if usr != nil {
		return usererrorutils.EMAIL_ALREADY_EXIST
	}
	return u.usrrep.UpdateEmail(ctx, userId, email)
}
