package docxreplacer

import (
	"context"
	"errors"
	"io"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/lukasjarosch/go-docx"
)

type docxReplacer struct{}

const _PROVIDER = "internal/docxreplacer"

func (r *docxReplacer) Replace(ctx context.Context, file []byte, wr io.Writer, replaceValues map[string]interface{}) error {
	select {
	case <-ctx.Done():
		return errors.New("context timeout")
	default:
		return replace(file, wr, replaceValues)
	}
}

func replace(file []byte, wr io.Writer, replaceValues map[string]interface{}) error {
	doc, err := docx.OpenBytes(file)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("open bytes", "Replace", _PROVIDER))
	}
	defer doc.Close()
	err = doc.ReplaceAll(replaceValues)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("replace values", "Replace", _PROVIDER))
	}
	err = doc.Write(wr)
	if err != nil {
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("write to buffer", "Replace", _PROVIDER))
	}

	return nil
}

func New() *docxReplacer {
	return &docxReplacer{}
}
