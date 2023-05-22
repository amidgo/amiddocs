package restapierror

import (
	"net/http"

	"github.com/amidgo/amiddocs/pkg/amiderrors"
)

var (
	WRONG_CLIENT_KEY = amiderrors.NewErrorResponse("wrong client key", http.StatusBadRequest, "wrong_client_key")
)
