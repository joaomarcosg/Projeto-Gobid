package api

import (
	"net/http"

	"github.com/joaomarcosg/Projeto-Gobid/internal/jsonutils"
	"github.com/joaomarcosg/Projeto-Gobid/internal/usecase/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {

	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"user_name": data.UserName,
		"email":     data.Email,
		"password":  data.Password,
		"bio":       data.Bio,
	})

}
