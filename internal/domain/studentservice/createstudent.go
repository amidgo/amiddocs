package studentservice

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
)

func (s *StudentService) CreateStudent(
	ctx context.Context,
	student *studentmodel.CreateStudentDTO,
) (*studentmodel.StudentDTO, error) {

	studentDTO, err := s.ValidateCreateStudentDTO(ctx, student)
	if err != nil {
		return nil, err
	}
	studentDTO, err = s.studentRep.InsertStudent(ctx, studentDTO)
	if err != nil {
		return nil, err
	}
	return studentDTO, nil
}

/*
Validate create student dto with database call

check email, docnumber unique field
check group is exists
check department is exists
map studentmodel.CreateStudentDTO to studentmodel.StudentDTO and return it
*/
func (s *StudentService) ValidateCreateStudentDTO(ctx context.Context, student *studentmodel.CreateStudentDTO) (*studentmodel.StudentDTO, error) {
	group, err := s.groupProv.GroupByName(ctx, student.GroupName)
	if err != nil {
		return nil, err
	}
	fmt.Println(group.StudyDepartmentId)
	department, err := s.depProv.StudyDepartmentById(ctx, group.StudyDepartmentId)
	if err != nil {
		return nil, err
	}
	fmt.Println(department)
	user := usermodel.NewCreateUserDTO(student.Name, student.Surname, student.FatherName, student.Email, make([]userfields.Role, 0))
	login, password, err := s.generateLoginAndPassword(ctx, user)
	if err != nil {
		return nil, err
	}
	fmt.Println(user)
	studentDTO := student.StudentDTO(login, password, group, &department.DepartmentDTO)
	return studentDTO, err
}

func (s *StudentService) generateLoginAndPassword(ctx context.Context, user *usermodel.CreateUserDTO) (userfields.Login, userfields.Password, error) {
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
