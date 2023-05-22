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
	"golang.org/x/sync/errgroup"
)

func (s *studentService) CreateStudent(
	ctx context.Context,
	student *studentmodel.CreateStudentDTO,
) (*studentmodel.StudentDTO, error) {

	errgr, errctx := errgroup.WithContext(ctx)

	errgr.Go(func() error {
		return s.checkOrderNumber(errctx, student.Document.OrderNumber)
	})
	errgr.Go(func() error {
		return s.checkDocNumber(errctx, student.Document.DocNumber)
	})
	errgr.Go(func() error {
		return s.checkEmail(errctx, student.User.Email)
	})
	err := errgr.Wait()
	group, err := s.groupProv.GroupByName(ctx, student.GroupName)
	if err != nil {
		return nil, err
	}
	department, err := s.depProv.StudyDepartment(ctx, group.StudyDepartmentId)
	if err != nil {
		return nil, err
	}
	user := usermodel.NewCreateUserDTO(student.User.Name, student.User.Surname, student.User.FatherName, student.User.Email, make([]userfields.Role, 0))
	login, password, err := s.generateLoginAndPassword(ctx, user)
	if err != nil {
		return nil, err
	}
	studentDTO := student.StudentDTO(login, password, group, department)
	studentDTO, err = s.studentRep.InsertStudent(ctx, studentDTO)
	if err != nil {
		return nil, err
	}
	return studentDTO, nil
}

func (s *studentService) checkOrderNumber(ctx context.Context, orderNumber stdocfields.OrderNumber) error {
	_, err := s.studentDocProv.DocumentByOrderNumber(ctx, orderNumber)
	if !amiderrors.Is(err, stdocerror.DOC_NOT_FOUND) {
		return stdocerror.ORDER_NUMBER_EXIST
	}
	return nil
}

func (s *studentService) checkDocNumber(ctx context.Context, docNumber stdocfields.DocNumber) error {
	_, err := s.studentDocProv.DocumentByDocNumber(ctx, docNumber)
	if !amiderrors.Is(err, stdocerror.DOC_NOT_FOUND) {
		return stdocerror.DOC_NUMBER_EXIST
	}
	return nil
}

func (s *studentService) checkEmail(ctx context.Context, email userfields.Email) error {
	if len(email) == 0 {
		return nil
	}
	_, err := s.userProv.UserByEmail(ctx, email)
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
