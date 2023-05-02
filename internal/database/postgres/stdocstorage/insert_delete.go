package stdocstorage

import (
	"context"

	"github.com/amidgo/amiddocs/internal/models/stdocmodel"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

func (st *studentDocStorage) InsertDocument(ctx context.Context, doc *stdocmodel.StudentDocumentDTO) (*stdocmodel.StudentDocumentDTO, error) {
	rows, err := st.p.DB.NamedQueryContext(ctx,
		`INSERT INTO student_documents (doc_number,order_number,order_date,study_start_date)
		VALUES (:doc_number,:order_number,:order_date,:study_start_date)
		RETURNING id`,
		doc,
	)
	if err != nil {
		return nil, studentDocumentError(err, amiderrors.NewCause("insert student query", "InsertDocument", _PROVIDER))
	}
	for rows.Next() {
		err = rows.Scan(&doc.ID)
		if err != nil {
			return nil, studentDocumentError(err, amiderrors.NewCause("scan student id", "InsertDocument", _PROVIDER))
		}
	}
	return doc, nil
}

// create table if not exists student_documents(
//     id bigserial primary key,
//     doc_number varchar(60) not null unique,
//     order_number varchar(60) not null,
//     order_date date not null,
//     study_start_date date not null
// );
