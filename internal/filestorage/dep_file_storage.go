package filestorage

import (
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/markbates/pkger"
)

func NewDepFileStorage(app *fiber.App, filepath string) FileStorage {
	depFileStorage := New(path.Join(os.Getenv("PWD"), filepath))
	app.Use(filepath, filesystem.New(filesystem.Config{
		Root: pkger.Dir(filepath),
	}))
	return depFileStorage
}
