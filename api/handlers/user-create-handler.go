package handlers

import (
	"../common"
	"../connect"
	"../model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		common.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	user := model.User{}
	err = json.Unmarshal(body, &user)

	if err != nil {
		common.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("")

	if err != nil {
		common.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	userCreated, err := user.SaveUser(connect.Connection)

	if err != nil {

		formattedError := common.FormatError(err.Error())

		common.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))

	common.JSON(w, http.StatusCreated, userCreated)
}
