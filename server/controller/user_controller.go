package controller

import (
	"go-react-todo/server/models"
	"go-react-todo/server/usecase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type UserController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &UserController{uu}
}

func (c *UserController) SignUp(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	resUser, err := c.uu.SignUp(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, resUser)
}

func (c *UserController) LogIn(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := c.uu.LogIn(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
	return ctx.NoContent(http.StatusOK)

}
func (c *UserController) LogOut(ctx echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	ctx.SetCookie(cookie)
	return ctx.NoContent(http.StatusOK)
}

func (c *UserController) CsrfToken(ctx echo.Context) error {
	token := ctx.Get("csrf")
	return ctx.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
