package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vineetdaniel/taskmanager/common"
	"github.com/vineetdaniel/taskmanager/data"
	"github.com/vineetdaniel/taskmanager/models"
)

//Handle for HTTP Post - "/user/register"
//Add a new user document

func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	//Decode the incoming user json
	err := json.Decoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w,
			err, "Invalid User Data",
			500,
		)
		return
	}

	user := *dataResource.Data
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	//Insert user document
	repo.CreateUser(user)
	//clean up hash password from resposne
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
		common.DisplayAppError(w,
			err, "An unexpected error", 500,
		)

		return

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

//Handler for HTTP Post - "/users/login"
//Authenticate with username and password

func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	//decode the incoming login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w,
			err, "Invalid Login data",
			500,
		)
		return
	}
	loginModel := dataResource.Data
	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("users")
	repo := &data.UserRepository{c}
	//authenticate the login user
	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(w,
			err, "Invalid login credentials",
			401,
		)
		return
	} else {
		//Genearate JWT Token
		token, err = common.GenerateJWT(user.Email, "member")
		if err != nil {
			common.DisplayAppError(w,
				err, "error generating token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser = AuthUserModel{
			User:  user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
		if err != nil {
			common.DisplayAppError(w,
				err, "An unexpected error has occurred", 500,
			)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
