package docxreplacer

import (
	"bytes"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/lukasjarosch/go-docx"
)

type docxReplacer struct{}

const _PROVIDER = "internal/docxreplacer"

func (r *docxReplacer) Replace(file []byte, replaceValues map[string]interface{}) ([]byte, error) {
	doc, err := docx.OpenBytes(file)
	if err != nil {
		return nil, amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("open bytes", "Replace", _PROVIDER))
	}
	defer doc.Close()
	err = doc.ReplaceAll(replaceValues)
	if err != nil {
		return nil, amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("replace values", "Replace", _PROVIDER))
	}
	wr := &bytes.Buffer{}
	err = doc.Write(wr)
	if err != nil {
		return nil, amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("write to buffer", "Replace", _PROVIDER))
	}
	return wr.Bytes(), nil
}

func New() *docxReplacer {
	return &docxReplacer{}
}
