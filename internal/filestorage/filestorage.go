package filestorage

import (
	"context"
	"errors"
	"os"
	"path"
	"sync"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const _PROVIDER = "internal/filestorage"

type fileStorage struct {
	mu            *sync.RWMutex
	storageFolder string
}

func New(storageFolder string) *fileStorage {
	os.Mkdir(storageFolder, os.ModePerm)
	return &fileStorage{storageFolder: storageFolder, mu: &sync.RWMutex{}}
}

func (s *fileStorage) FilePath(filename string) string {
	return path.Join(s.storageFolder, filename)
}

func (s *fileStorage) File(ctx context.Context, filename string) ([]byte, error) {
	select {
	case <-ctx.Done():
		err := errors.New("context timeout with read file " + s.FilePath(filename))
		return nil, amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("get file", "File", _PROVIDER))
	default:
		s.mu.RLock()
		defer s.mu.RUnlock()
		b, err := os.ReadFile(s.FilePath(filename))
		if err != nil {
			return nil, amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("get file", "File", _PROVIDER))
		}
		return b, nil
	}

}

func (s *fileStorage) PutFile(ctx context.Context, filename string, b []byte) error {
	select {
	case <-ctx.Done():
		err := errors.New("context timeout with write file " + s.FilePath(filename))
		return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("put file", "PutFile", _PROVIDER))
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		err := os.WriteFile(s.FilePath(filename), b, os.ModePerm)
		if err != nil {
			return amiderrors.NewInternalErrorResponse(err, amiderrors.NewCause("put file", "PutFile", _PROVIDER))
		}
		return nil
	}

}
