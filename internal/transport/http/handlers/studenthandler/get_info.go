package studenthandler

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/gofiber/fiber/v2"
)

// GetStudentById godoc
//
//	@Summary		GetStudentByID
//	@Description	get student by id from query param
//	@Tags			students
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	studentmodel.StudentDTO
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		404	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Security		Bearer
//	@Security		Token
//	@Router			/students/info [get]
func (h *studentHandler) StudentInfo(c *fiber.Ctx) error {
	id, err := h.jwt.UserID(c)
	if err != nil {
		return amiderrors.Wrap(err, amiderrors.NewCause("get user id", "StudentInfo", _PROVIDER))
	}
	student, err := h.studentP.StudentByUserId(c.UserContext(), id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(student)
}
