package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/matheusgcoppi/barber-finance-api/database"
	"github.com/matheusgcoppi/barber-finance-api/database/model"
	"github.com/matheusgcoppi/barber-finance-api/repository"
	"github.com/matheusgcoppi/barber-finance-api/utils"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"regexp"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Sub       uint `json:"sub"`
	AccountID uint `json:"account_id"`
}

type APIServer struct {
	store            *database.CustomDB
	repositoryServer *repository.DbRepository
}

func NewAPIServer(store *database.CustomDB, repository *repository.DbRepository) *APIServer {
	return &APIServer{
		store:            store,
		repositoryServer: repository,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (a *APIServer) HandleLogin(c echo.Context) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	user, err := a.repositoryServer.LoginUser(body.Email, body.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	account, err := a.repositoryServer.GetAccount(user.ID)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
		},
		Sub:       user.ID,
		AccountID: account.ID,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid to create token",
		})
	}

	cookie := new(http.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 24 * 30)
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Secure = true
	c.SetCookie(cookie)

	type LoginResponse struct {
		User      *model.User `json:"user"`
		AccountID uint        `json:"accountId"`
	}

	return c.JSON(http.StatusOK, LoginResponse{
		User:      user,
		AccountID: account.ID,
	})
}

func (a *APIServer) HandleIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, "hey")
}

func (a *APIServer) HandleCreateUser(c echo.Context) error {
	userDTO := new(model.UserDTO)
	if err := c.Bind(userDTO); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if userDTO.Type == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Type is required."})
	}

	if (userDTO.Type != model.System) && (userDTO.Type != model.Support) && (userDTO.Type != model.Admin) {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "This Role does not exist."})
	}

	if userDTO.Username == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username is required."})
	}

	if userDTO.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email is required."})
	}

	if userDTO.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password is required."})
	}

	password, err := HashPassword(userDTO.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newUser := repository.NewUser(
		true,
		int(userDTO.Type),
		userDTO.Username,
		userDTO.Email,
		password,
	)

	user, account, err := a.repositoryServer.CreateUser(newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if account != nil {
		user.Account = account
	}

	return c.JSON(http.StatusOK, user)
}

func (a *APIServer) IndexHandler(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/plain")
	returnStatus := http.StatusOK
	c.Response().WriteHeader(returnStatus)
	message := fmt.Sprintf("Hello %s!", c.Request().UserAgent())
	_, err := c.Response().Write([]byte(message))
	if err != nil {
		return err
	}
	return nil
}

func (a *APIServer) HandleGetUser(c echo.Context) error {
	err, result := a.repositoryServer.GetUser()
	if err != nil {
		err := c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		if err != nil {
			return err
		}
	}
	return c.JSON(http.StatusOK, result)
}

func (a *APIServer) HandleGetUserByID(c echo.Context) error {
	id := c.Param("id")
	result, err := a.repositoryServer.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func (a *APIServer) HandleDeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := a.repositoryServer.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"result": "User with id = " + id + " was deleted successfully"})
}

func (a *APIServer) HandleUpdateUser(c echo.Context) error {
	id := c.Param("id")
	updatedUser := new(model.UserDTO)
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if updatedUser.Type != 0 {
		if (updatedUser.Type != model.System) && (updatedUser.Type != model.Support) && (updatedUser.Type != model.Admin) {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "This Role does not exist."})
		}
	}

	if updatedUser.Email != "" {
		match, _ := regexp.MatchString(utils.EmailPattern, updatedUser.Email)
		if match == false {
			println(match)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Email Address"})
		}
	}

	if updatedUser.Password != "" {
		password, err := HashPassword(updatedUser.Password)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]error{"error": err})
		}
		updatedUser.Password = password
	}

	user, err := a.repositoryServer.UpdateUser(updatedUser, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, user)
}

func (a *APIServer) Validate(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]bool{"message": true})
}

func (a *APIServer) HandleRequestForgotPassword(c echo.Context) error {
	type emailRequest struct {
		Email string `json:"email"`
	}
	input := new(emailRequest)

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if input.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Field E-mail cannot be null"})
	}

	err := a.repositoryServer.ForgotPassword(input.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	cookie := new(http.Cookie)
	cookie.Name = "email"
	cookie.Value = input.Email
	cookie.Expires = time.Now().Add(time.Minute * 20)
	cookie.SameSite = http.SameSiteStrictMode
	cookie.Secure = true
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Email with Token was sent to %s", input.Email)})
}

func (a *APIServer) HandleResetPassword(c echo.Context) error {
	type tokenRequest struct {
		Token string `json:"token"`
	}
	input := new(tokenRequest)
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	cookie, err := c.Cookie("email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": "Cookie does not exist or could not be read"})
	}

	err = a.repositoryServer.ChangePassword(cookie.Value, input.Token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{"result": "Your password was changed successfully"})
}

func (a *APIServer) HandleResetPasswordd(c echo.Context) error {
	token := c.Param("token")
	type passwordInput struct {
		Password string `json:"password"`
	}
	input := new(passwordInput)

	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := a.repositoryServer.TestResetPassword(token, input.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{"result": "Your password was changed successfully"})

	//err = a.repositoryServer.ChangePassword(cookie.Value, input.Token)
	//if err != nil {
	//	return c.JSON(http.StatusBadRequest, map[string]string{"Error": err.Error()})
	//}

	//return c.JSON(http.StatusBadRequest, map[string]string{"result": "Your password was changed successfully"})
}
