package handlers

import (
	"../common"
	"../connect"
	"../model"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func SignIn(email, password string) (string, error) {
	var err error

	var user model.User
	db := connect.Connection

	err = db.Debug().First(&user, "email = ?", email).Take(&user).Error

	//err = connect.Connection.First(model.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = model.VerifyPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return common.CreateToken(user.ID)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	var user model.User

	if err != nil {
		common.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user, err = model.UnmarshalUser(body)

	if err != nil {
		common.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")

	if err != nil {
		common.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := SignIn(user.Email, user.Password)

	if err != nil {
		formattedError := common.FormatError(err.Error())
		common.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}

	common.JSON(w, http.StatusOK, token)
}
