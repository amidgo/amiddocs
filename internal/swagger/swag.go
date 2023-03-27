package swagger

import (
	_ "github.com/amidgo/amiddocs/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

//	@title						Fiber Example API
//	@version					1.0
//	@description				This is a sample swagger for Fiber
//	@termsOfService				http://swagger.io/terms/
//	@license.name				Apache 2.0
//	@license.url				http://www.apache.org/licenses/LICENSE-2.0.html
//
//	@securityDefinitions.apikey	Bearer
//
//	@in							header
//	@name						Authorization
//	@description				Bearer Token Auth
//
//	@host						localhost:10101
//	@BasePath					/
func SetUp(app *fiber.App) {
	app.Get("/swagger/*", swagger.HandlerDefault)
}
