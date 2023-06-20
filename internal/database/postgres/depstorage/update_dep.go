package depstorage

import (
	"context"
	"fmt"

	"github.com/amidgo/amiddocs/internal/models/depmodel"
	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/amidgo/amiddocs/pkg/amiderrors"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	updatePhotoUrlQuery = fmt.Sprintf(
		`UPDATE %s SET %s = $1 WHERE %s = $2`,
		depmodel.DepartmentTable,

		depmodel.SQL.ImageUrl,
		depmodel.SQL.ID,
	)
)

func (s *departmentStorage) UpdateDepartmentPhoto(ctx context.Context, id uint64, imageUrl depfields.ImageUrl) error {
	_, err := s.p.Pool.Exec(ctx, updatePhotoUrlQuery,
		id,
		pgtype.Text{String: string(imageUrl), Valid: imageUrl != ""},
	)
	if err != nil {
		return departmentError(err, amiderrors.NewCause("update department photo", "UpdateDepartmentPhoto", _PROVIDER))
	}
	return nil
}
