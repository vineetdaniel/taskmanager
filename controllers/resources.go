package controllers

import (
	"github.com/vineetdaniel/AiOps/apiv1/models"
)

type (
	UserResource struct {
		Data models.User `json:"data"`
	}

	//For post - user/login

	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	//Response for authorized user post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	//Model for authentication

	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//model for authorized user with access token

	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)

type (
	//for Post/Put
	UrlResource struct {
		Data models.Url `json:"data"`
	}
	//for GET
	UrlsResource struct {
		Data []models.Url `json:"data"`
	}
)
