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

type FileStorage interface {
	FilePath(filename string) string
	File(ctx context.Context, filename string) ([]byte, error)
	PutFile(ctx context.Context, filename string, b []byte) error
	RemoveFile(ctx context.Context, filename string) error
}

func New(storageFolder string) FileStorage {
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
		return nil, amiderrors.Wrap(err, amiderrors.NewCause("get file", "File", _PROVIDER))
	default:
		s.mu.RLock()
		defer s.mu.RUnlock()
		b, err := os.ReadFile(s.FilePath(filename))
		if err != nil {
			return nil, amiderrors.Wrap(err, amiderrors.NewCause("get file", "File", _PROVIDER))
		}
		return b, nil
	}

}

func (s *fileStorage) PutFile(ctx context.Context, filename string, b []byte) error {
	select {
	case <-ctx.Done():
		err := errors.New("context timeout with write file " + s.FilePath(filename))
		return amiderrors.Wrap(err, amiderrors.NewCause("put file", "PutFile", _PROVIDER))
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		err := os.WriteFile(s.FilePath(filename), b, os.ModePerm)
		if err != nil {
			return amiderrors.Wrap(err, amiderrors.NewCause("put file", "PutFile", _PROVIDER))
		}
		return nil
	}

}

func (s *fileStorage) RemoveFile(ctx context.Context, filename string) error {
	select {
	case <-ctx.Done():
		err := errors.New("context timeout with delete file " + s.FilePath(filename))
		return amiderrors.Wrap(err, amiderrors.NewCause("drop file", "RemoveFile", _PROVIDER))
	default:
		s.mu.Lock()
		defer s.mu.Unlock()
		err := os.Remove(s.FilePath(filename))
		if err != nil {
			return amiderrors.Wrap(err, amiderrors.NewCause("drop file", "RemoveFile", _PROVIDER))
		}
		return nil
	}
}
