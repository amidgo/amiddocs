package docgenerator

import (
	"bytes"
	"context"
	"io"

	"github.com/amidgo/amiddocs/internal/errorutils/doctypeerror"
	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
)

func (d *docGenerator) GenerateDocument(ctx context.Context, wr io.Writer, reqId uint64) error {
	req, err := d.reqProvider.RequestById(ctx, reqId)
	if err != nil {
		return err
	}
	buf := &bytes.Buffer{}
	err = d.docTempProvider.DocumentTemplate(ctx, buf, req.DepartmentID, req.DocumentType.Type)
	if err != nil {
		return err
	}
	switch req.DocumentType.Type {
	case doctypefields.STUDY_DOCUMENT_BUDGET:
		return d.replaceStudyDocument(ctx, buf.Bytes(), wr, req.UserID, reqId)
	case doctypefields.STUDY_DOCUMENT_NO_BUDGET:
		return d.replaceStudyDocument(ctx, buf.Bytes(), wr, req.UserID, reqId)
	default:
		return doctypeerror.DOC_TYPE_NOT_FOUND
	}
}
