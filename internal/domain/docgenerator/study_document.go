package docgenerator

import (
	"context"
	"io"

	"github.com/amidgo/amiddocs/internal/models/studentmodel"
	"github.com/amidgo/amiddocs/pkg/amidtime"
)

func studyDocumentBudgetReplaceValues(
	reqId uint64,
	student *studentmodel.StudentDTO,
) map[string]interface{} {
	return map[string]interface{}{
		"{DATE}":             amidtime.Now().Date().Human(),
		"{DEP_SN}":           student.Department.ShortName,
		"{ID}":               reqId,
		"{STUDENT_FIO}":      student.User.Fio(),
		"{STUDY_START_DATE}": student.Document.EducationStartDate.Human(),
		"{STUDY_END_DATE}":   student.Group.EducationFinishDate.Human(),
		"{ORDER_NUMBER}":     student.Document.OrderNumber,
		"{ORDER_DATE}":       student.Document.OrderDate.Human(),
		"{DOP_INFO}":         "",
	}
}

func (d *docGenerator) replaceStudyDocument(ctx context.Context, file []byte, wr io.Writer, userId uint64, reqId uint64) error {
	student, err := d.studentProvider.StudentByUserId(ctx, userId)
	if err != nil {
		return err
	}
	err = d.docxReplacer.Replace(ctx, file, wr, studyDocumentBudgetReplaceValues(reqId, student))
	if err != nil {
		return err
	}
	return nil
}
