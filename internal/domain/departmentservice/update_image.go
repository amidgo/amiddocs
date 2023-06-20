package departmentservice

import (
	"context"
	"path"

	"github.com/amidgo/amiddocs/internal/models/depmodel/depfields"
	"github.com/google/uuid"
)

func (s *departmentService) UpdateImageUrl(ctx context.Context, id uint64, b []byte) error {
	dep, err := s.depProvider.DepartmentById(ctx, id)
	if err != nil {
		return err
	}
	file, err := s.updatePhotoFS(ctx, dep.ImageUrl, b)
	if err != nil {
		return err
	}
	err = s.updatePhotoDB(ctx, id, depfields.ImageUrl(s.depFS.FilePath(file)))
	if err != nil {
		s.revertUpdatePhotoFS(ctx, file)
		return err
	}
	return nil
}

// update photo if exist or insert new and return file name in <file>.svg
func (s *departmentService) updatePhotoFS(ctx context.Context, old depfields.ImageUrl, b []byte) (string, error) {
	if old != "" {
		_, file := path.Split(string(old))
		err := s.depFS.PutFile(ctx, file, b)
		if err != nil {
			return "", err
		}
		return file, nil
	}
	file := uuid.New().String() + ".svg"
	err := s.depFS.PutFile(ctx, file, b)
	if err != nil {
		return "", err
	}
	return file, nil
}

// remove file in fs
func (s *departmentService) revertUpdatePhotoFS(ctx context.Context, fileName string) error {
	err := s.depFS.RemoveFile(ctx, fileName)
	if err != nil {
		return err
	}
	return nil
}

// update photo in db
func (s *departmentService) updatePhotoDB(ctx context.Context, id uint64, imageUrl depfields.ImageUrl) error {
	err := s.depRep.UpdateDepartmentPhoto(ctx, id, imageUrl)
	if err != nil {
		return err
	}
	return nil
}
