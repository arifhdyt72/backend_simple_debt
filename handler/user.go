package handler

import (
	"backend_test_debt/helper"
	"backend_test_debt/initializer"
	"backend_test_debt/middleware"
	"backend_test_debt/models"
	"backend_test_debt/user"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	userRepository user.Repository
	jwtService     middleware.JWTService
}

func NewUserHandler(userRepository user.Repository, jwtService middleware.JWTService) *userHandler {
	return &userHandler{
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var input user.UserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("create user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var user models.User
	truest := true
	user.Name = input.Name
	user.Username = input.Username
	user.Password = input.Password
	user.Email = input.Email
	user.Status = &truest

	result, err := h.userRepository.Create(user)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("create user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully created user", http.StatusCreated, "success", result)
	c.JSON(http.StatusCreated, response)
}

func (h *userHandler) GetAllData(c *gin.Context) {
	input := c.Request.URL.Query()
	var data user.DatatableInput

	page, _ := strconv.Atoi(input.Get("page"))
	first, _ := strconv.Atoi(input.Get("first"))
	rows, _ := strconv.Atoi(input.Get("rows"))
	pageCount, _ := strconv.Atoi(input.Get("pageCount"))

	data.Filters = input.Get("filters")
	data.Page = page
	data.First = first
	data.Rows = rows
	data.PageCount = pageCount
	data.SortField = input.Get("sortField")
	data.SortOrder = input.Get("sortOrder")

	listUser, err := h.userRepository.GetAllData(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get all user failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	count, err := h.userRepository.GetCountData(data)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("Get all user failed", http.StatusBadRequest, "failed", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	result := gin.H{"data": listUser, "countData": count}
	response := helper.ApiResponse("Successfully to get users", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginHandler(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("login user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user, err := h.userRepository.FindUserByUsernameOrEmail(input.Username)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("login user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	if user.ID == 0 {
		errorMessage := gin.H{"errors": errors.New("username or email not found")}
		response := helper.ApiResponse("login user failed not found", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(input.Password))
	if err != nil {
		errorMessage := gin.H{"errors": errors.New("invalid password")}
		response := helper.ApiResponse("login user failed password", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.jwtService.GenerateToken(user)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("login user failed token", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result := gin.H{"token": token, "user": user}
	response := helper.ApiResponse("successfully login user", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var input user.UserUpdate

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("update user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	var user models.User
	initializer.DB.First(&user, int(input.ID))

	user.Name = input.Name
	user.Username = input.Username
	user.Password = input.Password
	user.Email = input.Email

	result, err := h.userRepository.Update(user)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("update user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully updated user", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	var input user.InputId

	err := c.ShouldBindUri(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("delete user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	result, err := h.userRepository.DeleteData(input.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.ApiResponse("delete user failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.ApiResponse("Successfully deleted user", http.StatusOK, "success", result)
	c.JSON(http.StatusOK, response)
}
