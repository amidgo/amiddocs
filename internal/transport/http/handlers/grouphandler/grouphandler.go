package grouphandler

import (
	"context"

	"github.com/amidgo/amiddocs/internal/jwttoken"
	"github.com/amidgo/amiddocs/internal/models/groupmodel"
	"github.com/amidgo/amiddocs/internal/models/usermodel/userfields"
	"github.com/gofiber/fiber/v2"
)

const _PROVIDER = "internal/transport/http/handlers/grouphandler"

type groupService interface {
	CreateGroup(ctx context.Context, group *groupmodel.GroupDTO) (*groupmodel.GroupDTO, error)
}

type groupProvider interface {
	GroupById(ctx context.Context, id uint64) (*groupmodel.GroupDTO, error)
}

type roleValidator interface {
	ValidateRole(c *fiber.Ctx, role userfields.Role) error
}

type GroupHandler struct {
	groupS groupService
	roleV  roleValidator
	groupP groupProvider
}

const _GROUP_PATH = "/groups"

const (
	_CREATE_GROUP = "/create"
	_GET_BY_ID    = "/get-by-id"
)

func SetUp(
	app *fiber.App,
	jwt func(c *fiber.Ctx) error,
	groupS groupService,
	groupP groupProvider,
) {
	handler := &GroupHandler{groupS: groupS, groupP: groupP}
	route := app.Group(_GROUP_PATH)

	route.Post(_CREATE_GROUP, jwt, jwttoken.AdminAccess, handler.CreateGroup)
	route.Get(_GET_BY_ID, handler.GetGroupById)
}
