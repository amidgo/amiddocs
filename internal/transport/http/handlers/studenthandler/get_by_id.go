package studenthandler

import (
	"net/http"
	"strconv"

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
//	@Param			id	query		uint64	true	"student id"
//	@Success		200	{object}	studentmodel.StudentDTO
//	@Failure		400	{object}	amiderrors.ErrorResponse
//	@Failure		404	{object}	amiderrors.ErrorResponse
//	@Failure		500	{object}	amiderrors.ErrorResponse
//	@Router			/students/get-by-id [get]
func (h *studentHandler) GetStudentById(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("parse body", "GetDocumentById", _PROVIDER))
	}
	st, err := h.studentP.StudentById(c.UserContext(), id)
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(st)
}
