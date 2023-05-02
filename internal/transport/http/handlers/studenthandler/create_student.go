package studenthandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/amidgo/amiddocs/pkg/validate"
	"github.com/gofiber/fiber/v2"
)

// CreateStudent godoc
//
//	@Summary		Create Student
//	@Description	create student
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Param			student	body		studentmodel.CreateStudentDTO	true	"create student dto"
//	@Success		200		{object}	studentmodel.StudentDTO
//	@Failure		400		{object}	amiderrors.ErrorResponse
//	@Failure		401		{object}	amiderrors.ErrorResponse
//	@Failure		403		{object}	amiderrors.ErrorResponse
//	@Failure		404		{object}	amiderrors.ErrorResponse
//	@Failure		500		{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Router			/students/create [post]
func (h *studentHandler) CreateStudent(c *fiber.Ctx) error {

	createStudentDTO := new(studentmodel.CreateStudentDTO)
	err := c.BodyParser(createStudentDTO)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("parse body", "CreateStudent", _PROVIDER))
	}

	err = validate.ValidateFields(createStudentDTO)
	if err != nil {
		return err
	}
	st, err := h.studentS.CreateStudent(c.UserContext(), createStudentDTO)
	if err != nil {
		return err
	}
	return c.Status(http.StatusCreated).JSON(st)
}
