package studentservice

import (
	"context"

	"github.com/amidgo/amiddocs/internal/errorutils/stdocerror"
	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/stdocmodel/stdocfields"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (s *studentService) CreateStudent(
	ctx context.Context,
	student *studentmodel.CreateStudentDTO,
) (*studentmodel.StudentDTO, error) {
	group, err := s.groupRep.GroupByName(ctx, student.GroupName)
	if err != nil {
		return nil, err
	}

	err = s.checkOrderNumber(ctx, student.Document.OrderNumber)
	if err != nil {
		return nil, err
	}

	err = s.checkDocNumber(ctx, student.Document.DocNumber)
	if err != nil {
		return nil, err
	}
	err = s.checkEmail(ctx, student.User.Email)
	if err != nil {
		return nil, err
	}
	login, password, err := s.generateLoginAndPassword(ctx, student.User)
	if err != nil {
		return nil, err
	}
	studentDTO := studentmodel.NewStudentDTO(
		0,
		usermodel.NewUserDTO(
			0,
			login,
			password,
			student.User.Name,
			student.User.Surname,
			student.User.FatherName,
			student.User.Email,
			[]userfields.Role{userfields.STUDENT},
		),
		student.Document,
		group,
	)
	studentDTO, err = s.studentRep.InsertStudent(ctx, studentDTO)
	if err != nil {
		return nil, err
	}
	return studentDTO, nil
}

func (s *studentService) checkOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) error {
	_, err := s.studentDocRep.DocumentByOrderNumber(ctx, orderNumber)
	if !amiderrors.Is(err, stdocerror.DOC_NOT_FOUND) {
		return stdocerror.ORDER_NUMBER_EXIST
	}
	return nil
}

func (s *studentService) checkDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) error {
	_, err := s.studentDocRep.DocumentByDocNumber(ctx, docNumber)
	if !amiderrors.Is(err, stdocerror.DOC_NOT_FOUND) {
		return stdocerror.DOC_NUMBER_EXIST
	}
	return nil
}

func (s *studentService) checkEmail(ctx context.Context, email userfields.Email) error {
	if len(email) == 0 {
		return nil
	}
	_, err := s.userRep.UserByEmail(ctx, email)
	if !amiderrors.Is(err, usererror.NOT_FOUND) {
		return usererror.EMAIL_ALREADY_EXIST
	}
	return nil
}

func (s *studentService) generateLoginAndPassword(ctx context.Context, user *usermodel.CreateUserDTO) (userfields.Login, userfields.Password, error) {
	login, password, err := user.GenerateLoginAndPassword()
	if err != nil {
		return "", "", err
	}
	hash, err := s.encrypter.Hash(string(password))
	if err != nil {
		return "", "", err
	}
	return login, userfields.Password(hash), nil
}
