package repo

import (
	"errors"
	"shape-api/auth"
	"shape-api/db"
	"shape-api/form"
	"shape-api/model"

	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct{}

func (r UserRepo) Register(registerForm form.RegisterForm) (user model.User, err error) {
	var countUser int64
	if err := db.PgClient.Table("user").Model(&model.User{}).Where("user_name", registerForm.UserName).Count(&countUser).Error; err != nil {
		return user, err
	}
	if countUser > 0 {
		return user, errors.New("username already exists")
	}
	//Compare the password form and database if match
	bytePassword := []byte(registerForm.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	user = model.User{
		UserName:  registerForm.UserName,
		Password:  string(hashedPassword),
		FullName:  registerForm.FullName,
		Email:     registerForm.Email,
		Address:   registerForm.Address,
		LastLogin: time.Now().UnixMilli(),
		CreatedOn: time.Now().UnixMilli(),
	}
	result := db.PgClient.Table("user").Create(&user)
	if result.Error != nil {
		return user, errors.New("something went wrong, please try again later: " + result.Error.Error())
	}
	return user, nil
}

func (r UserRepo) Login(loginForm form.LoginForm) (user model.User, token model.Token, err error) {

	if err := db.PgClient.Table("user").Model(&model.User{}).Where("user_name", loginForm.UserName).First(&user).Error; err != nil {
		return user, token, errors.New("something went wrong, please try again later: " + err.Error())
	}

	//Compare the password form and database if match
	bytePassword := []byte(loginForm.Password)
	byteHashedPassword := []byte(user.Password)

	err = bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
	if err != nil {
		return user, token, err
	}

	//Generate the JWT auth token
	authController := new(auth.Auth)
	tokenDetails, err := authController.CreateToken(user.UserName)
	if err != nil {
		return user, token, err
	}

	token.AccessToken = tokenDetails.AccessToken
	token.RefreshToken = tokenDetails.RefreshToken
	return user, token, nil
}
