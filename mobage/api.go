package mobage

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/patapancakes/mobimu/common"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	action := r.PathValue("action")
	if strings.HasPrefix(action, "G") {
		split := strings.Split(action, "_")
		if len(split) != 2 {
			common.HTTPError(w, r, fmt.Sprintf("invalid action \"%s\"", action), http.StatusBadRequest)
			return
		}

		action = split[1]
	}

	switch action {
	case "Authorize":
	case "GetAvatar":
		// TODO: implement this, apps should work without it though

		w.WriteHeader(http.StatusOK)
		return
	case "GetLobbyList":
	case "GetLobbyInfo":
	case "CreateRoom":
	case "ConnectRoom":
	case "SyncRoom":
	case "DisconnectRoom":
	default:
		common.HTTPError(w, r, fmt.Sprintf("unhandled action \"%s\"", action), http.StatusBadRequest)
		return
	}
}
