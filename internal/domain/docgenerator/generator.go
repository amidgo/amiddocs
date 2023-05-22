package docgenerator

import (
	"context"
	"io"

	"github.com/amidgo/amiddocs/internal/models/doctypemodel/doctypefields"
	"github.com/amidgo/amiddocs/internal/models/reqmodel"
	"github.com/amidgo/amiddocs/internal/models/studentmodel"
)

type docxReplacer interface {
	Replace(ctx context.Context, file []byte, wr io.Writer, replaceValues map[string]interface{}) error
}

type studentProvider interface {
	StudentByUserId(ctx context.Context, id uint64) (*studentmodel.StudentDTO, error)
}

type userProvider interface {
}

type reqProvider interface {
	RequestById(ctx context.Context, reqId uint64) (*reqmodel.RequestDTO, error)
}

type docTempProvider interface {
	DocumentTemplate(
		ctx context.Context,
		wr io.Writer,
		depID uint64,
		docType doctypefields.DocumentType,
	) error
}

type docGenerator struct {
	docxReplacer    docxReplacer
	studentProvider studentProvider
	userProvider    userProvider
	reqProvider     reqProvider
	docTempProvider docTempProvider
}

func New(
	docxReplacer docxReplacer,
	stProv studentProvider,
	userProv userProvider,
	reqProv reqProvider,
	docTempProv docTempProvider,
) *docGenerator {
	return &docGenerator{
		docxReplacer:    docxReplacer,
		studentProvider: stProv,
		userProvider:    userProv,
		reqProvider:     reqProv,
		docTempProvider: docTempProv,
	}
}
