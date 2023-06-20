package restapierror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

const REST_TYPE = "client_key"

var (
	WRONG_CLIENT_KEY = amiderrors.NewException(http.StatusBadRequest, REST_TYPE, "wrong")
)
