package usecase

import (
	"go-react-todo/server/models"
	"go-react-todo/server/repository"
	"go-react-todo/server/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user models.User) (models.UserResponse, error)
	LogIn(user models.User) (string, error)
}

type UserUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &UserUsecase{ur, uv}
}

func (u *UserUsecase) SignUp(user models.User) (models.UserResponse, error) {
	if err := u.uv.ValidateUser(&user); err != nil {
		return models.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.UserResponse{}, err
	}

	newUser := models.User{
		Email:    user.Email,
		Password: string(hash),
	}

	if err = u.ur.CreateUser(&newUser); err != nil {
		return models.UserResponse{}, err
	}

	resUser := models.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email,
	}

	return resUser, nil
}

func (u *UserUsecase) LogIn(user models.User) (string, error) {
	dbUser := models.User{}

	if err := u.uv.ValidateUser(&user); err != nil {
		return "", err
	}

	if err := u.ur.GetUserByEmail(&dbUser, user.Email); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": dbUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
