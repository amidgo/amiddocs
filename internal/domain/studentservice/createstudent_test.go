package studentservice_test

import (
	"context"
	"testing"

	"github.com/amidgo/amiddocs/internal/domain/studentservice"
	"github.com/amidgo/amiddocs/internal/encrypt"
	"github.com/amidgo/amiddocs/internal/errorutils/departmenterror"
	"github.com/amidgo/amiddocs/internal/errorutils/grouperror"
	"github.com/amidgo/amiddocs/internal/errorutils/stdocerror"
	"github.com/amidgo/amiddocs/internal/errorutils/usererror"
	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/groupmodel/groupfields"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/pkg/amidtime"
	"github.com/amidgo/amiddocs/pkg/assert"
	"github.com/golang/mock/gomock"
)

var (
	validCreateStudentDTO = studentmodel.NewCreateStudentDTO(
		"test",
		"test",
		"test",
		"test@mail.ru",
		"test",
		"test",
		amidtime.NewDate(2020, 1, 1),
		amidtime.NewDate(2020, 1, 1),
		"test",
	)
	validGroupDTO = groupmodel.NewGroupDTO(1,
		"test",
		true,
		groupfields.FULL_TIME,
		amidtime.NewDate(2020, 1, 1),
		amidtime.NewDate(2020, 1, 1),
		1,
		1,
	)
	validDepartmentDTO = depmodel.NewDepartmentDTO(
		1,
		"test",
		"test",
	)
)

/*
данный тест тестирует проверку на наличие:

	студентов с такой почтой (если почта пустая то проверка не выполняется)
	документов с аналогичным номером
*/
func TestCheckDocNumberAndEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	groupProv := studentservice.NewMockGroupProvider(ctrl)
	studentDocProvider := studentservice.NewMockStudentDocProvider(ctrl)
	userProvider := studentservice.NewMockUserProvider(ctrl)
	studentProvider := studentservice.NewMockStudentProvider(ctrl)
	studentRepository := studentservice.NewMockStudentRepository(ctrl)
	departmentProvider := studentservice.NewMockDepProvider(ctrl)
	encrypter := encrypt.New(10)

	studentService := studentservice.New(groupProv, studentDocProvider, userProvider, studentProvider, studentRepository, departmentProvider, encrypter)

	ctx := context.Background()
	student := validCreateStudentDTO

	group := validGroupDTO
	dep := validDepartmentDTO

	// базовый ответ для успешного завершение функции
	groupProv.EXPECT().
		GroupByName(ctx, student.GroupName).
		Return(group, nil).AnyTimes()
	departmentProvider.EXPECT().
		StudyDepartmentById(ctx, group.StudyDepartmentId).
		Return(dep, nil).AnyTimes()
	studentRepository.EXPECT().
		InsertStudent(ctx, gomock.Any()).
		Return(nil, nil).
		AnyTimes()

	var exp error
	exp = stdocerror.DOC_NUMBER_EXIST

	// если провайдер не возвращает ошибку это означает что он нашел студенческий билет с таким номером
	// в таком случае функция должна выбросить ошибку что в базе данных такой билет уже существует
	studentDocProvider.EXPECT().
		DocumentByDocNumber(ctx, student.DocNumber).
		Return(nil, nil).
		Times(1)
	// если почта пустая то проверка не выполняется и по большому счёту неважно что вернёт функция
	// в данном случае почта не пустая и если провайдер найдёт студента по заданному эмейлу то
	// функция должна будет вернуть ошибку о том что почта уже существует
	userProvider.EXPECT().
		UserByEmail(ctx, student.Email).
		Return(nil, usererror.NOT_FOUND).
		Times(1)

	// в данном случае мы ожидаем ошибки о том что студенческий билет с таким номером уже существует
	// поскольку из studentDocProvider мы возвращаем пустую ошибку
	_, act := studentService.CreateStudent(ctx, student)
	assert.ErrorEqual(t, exp, act, "doc number exist test failed")

	// кейс когда мы не нашли документ с подобным номером
	studentDocProvider.EXPECT().
		DocumentByDocNumber(ctx, student.DocNumber).
		Return(nil, stdocerror.DOC_NOT_FOUND).
		Times(1)
	// в данном случае мы нашли пользователя с такой-же почтой
	userProvider.EXPECT().
		UserByEmail(ctx, student.Email).
		Return(nil, nil).
		Times(1)

	// в данном кейсе мы должны вернуть ошибку: пользователь с такой почтой уже существует
	exp = usererror.EMAIL_ALREADY_EXIST
	_, act = studentService.CreateStudent(ctx, student)
	assert.ErrorEqual(t, exp, act, "user email (not empty) test failed")

	// в данном случае проверка почты пользователя не происходит, потому что она пустая
	student.Email = ""
	exp = nil
	// здесь мы говорим что такого номера в базе ещё не существует
	studentDocProvider.EXPECT().
		DocumentByDocNumber(ctx, student.DocNumber).
		Return(nil, stdocerror.DOC_NOT_FOUND).
		Times(1)
	// почту мы не проверяли, проверка номера студенческого вернула положительный результат, ошибка должна равняться nil
	_, act = studentService.CreateStudent(ctx, student)
	assert.ErrorEqual(t, exp, act, "right input user email (empty) test failed")

	// тот же самый кейс как и в предыдущем только почта не пустая и мы включаем проверку эмейла
	// которая возвращает нам отсутствие пользователя с такой почтой в базе данных
	student.Email = "test@mail.ru"
	exp = nil
	studentDocProvider.EXPECT().
		DocumentByDocNumber(ctx, student.DocNumber).
		Return(nil, stdocerror.DOC_NOT_FOUND).
		Times(1)
	userProvider.EXPECT().
		UserByEmail(ctx, student.Email).
		Return(nil, usererror.NOT_FOUND).
		Times(1)

	_, act = studentService.CreateStudent(ctx, student)
	assert.ErrorEqual(t, exp, act, "right input test failed")

}

func TestCreateStudent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	groupProv := studentservice.NewMockGroupProvider(ctrl)
	studentDocProvider := studentservice.NewMockStudentDocProvider(ctrl)
	userProvider := studentservice.NewMockUserProvider(ctrl)
	studentProvider := studentservice.NewMockStudentProvider(ctrl)
	studentRepository := studentservice.NewMockStudentRepository(ctrl)
	departmentProvider := studentservice.NewMockDepProvider(ctrl)
	encrypter := encrypt.New(10)

	studentService := studentservice.New(groupProv, studentDocProvider, userProvider, studentProvider, studentRepository, departmentProvider, encrypter)

	ctx := context.Background()
	student := validCreateStudentDTO

	// ставим зашлушки чтобы проверки которые мы здесь не тестируем проходили без проблем
	studentDocProvider.EXPECT().
		DocumentByDocNumber(ctx, student.DocNumber).
		Return(nil, stdocerror.DOC_NOT_FOUND).
		AnyTimes()
	userProvider.EXPECT().
		UserByEmail(ctx, student.Email).
		Return(nil, usererror.NOT_FOUND).
		AnyTimes()
	studentRepository.EXPECT().
		InsertStudent(ctx, gomock.Any()).
		Return(nil, nil).
		AnyTimes()

	var exp error
	// проверяем кейс когда мы не нашли группу по имени
	// функция в таком случае должна вернуть ошибку которую ей передал провайдер
	exp = grouperror.GROUP_NOT_FOUND
	groupProv.EXPECT().
		GroupByName(ctx, student.GroupName).
		Return(nil, exp).
		Times(1)
	_, act := studentService.CreateStudent(ctx, student)
	assert.ErrorEqual(t, exp, act, "group error test failed")

	group := validGroupDTO
	// проверяем кейс при попытке достать из базы учебное отделение нам вернуло ошибку
	// функция в таком случае должна вернуть ошибку которую ей передал провайдер
	exp = departmenterror.DEPARTMENT_NOT_FOUND
	groupProv.EXPECT().
		GroupByName(ctx, student.GroupName).
		Return(group, nil).
		Times(1)
	departmentProvider.EXPECT().
		StudyDepartmentById(ctx, validGroupDTO.StudyDepartmentId).
		Return(nil, exp).
		Times(1)

	_, act = studentService.CreateStudent(ctx, student)
	assert.ErrorEqual(t, exp, act, "department error test failed")

	department := validDepartmentDTO
	// в данном кейсе у все вызовы из репозиториев проходят без ошибок
	exp = nil
	groupProv.EXPECT().
		GroupByName(ctx, student.GroupName).
		Return(group, nil).
		Times(1)
	departmentProvider.EXPECT().
		StudyDepartmentById(ctx, group.StudyDepartmentId).
		Return(department, nil).
		Times(1)

	_, act = studentService.CreateStudent(ctx, student)
	assert.ErrorEqual(t, exp, act, "department error test failed")
}
