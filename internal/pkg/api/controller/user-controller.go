package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/grootkng/clean-arch-golang/internal/domain/entity"
	"github.com/grootkng/clean-arch-golang/internal/domain/usecase"
	"github.com/grootkng/clean-arch-golang/internal/pkg/api/dto"
)

type UserController struct {
	Implementation usecase.IUserUsecase
}

func NewUserController(imp usecase.IUserUsecase) *UserController {
	return &UserController{
		Implementation: imp,
	}
}

// CreateUser godoc
// @Summary      Create user
// @Description  create a user
// @Tags         users
// @Produce      json
// @Param        user body      dto.UserDto true  "Payload"
// @Success      201  {string}  string
// @Failure      500  {string}  dto.HTTPInternalServerErrorDTO
// @Router       /users [post]
func (uc *UserController) Create(c *gin.Context) {
	var payload dto.UserDto

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, setMessage("Invalid payload. Error: "+err.Error()))
		return
	}

	u := &entity.User{
		Name:   payload.Name,
		Age:    int8(payload.Age),
		Gender: payload.Gender,
	}

	if err := uc.Implementation.Create(u); err != nil {
		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	c.JSON(http.StatusCreated, setMessage("Success"))
}

// FindAllUsers godoc
// @Summary      Find all users
// @Description  retrieve all users
// @Tags         users
// @Produce      json
// @Param        page   		query      	int  true  "page"
// @Param        pageSize   query      	int  true  "pageSize"
// @Success      200  			{array}  		entity.User
// @Success      204  			{array}   	entity.User
// @Failure      428  			{string}  	dto.HTTPInternalServerErrorDTO
// @Failure      500  			{string}  	dto.HTTPInternalServerErrorDTO
// @Router       /users [get]
func (uc *UserController) FindAll(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusPreconditionRequired, setMessage("Send page correctly"))
		return
	}

	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		c.JSON(http.StatusPreconditionRequired, setMessage("Send pageSize correctly"))
		return
	}

	users, err := uc.Implementation.FindAll(page, pageSize)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, setMessage(err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	if len(users) > 0 {
		c.JSON(http.StatusOK, users)
		return
	}

	c.JSON(http.StatusNoContent, users)
}

// FindByIDUser godoc
// @Summary      Find user by ID
// @Description  retrieve a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  entity.User
// @Failure      500  {string}  dto.HTTPInternalServerErrorDTO
// @Router       /users/{id} [get]
func (uc *UserController) FindBy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	user, err := uc.Implementation.FindBy(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	if user == nil {
		c.JSON(http.StatusNoContent, &entity.User{})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateByIDUser godoc
// @Summary      Update user by ID
// @Description  update a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Param        user body      dto.UserDto true  "Payload"
// @Success      200  {string}  string
// @Failure      500  {string}  dto.HTTPInternalServerErrorDTO
// @Router       /users/{id} [put]
func (uc *UserController) UpdateBy(c *gin.Context) {
	var payload dto.UserDto

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, setMessage("Invalid payload. Error: "+err.Error()))
		return
	}

	u := &entity.User{
		Id:     id,
		Name:   payload.Name,
		Age:    int8(payload.Age),
		Gender: payload.Gender,
	}

	if err := uc.Implementation.UpdateBy(u); err != nil {
		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, setMessage("Success"))
}

// DeleteByIDUser godoc
// @Summary      Delete user by ID
// @Description  delete a user by ID
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {string}  string
// @Failure      500  {string}  dto.HTTPInternalServerErrorDTO
// @Router       /users/{id} [delete]
func (uc *UserController) DeleteBy(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	err = uc.Implementation.DeleteBy(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, setMessage("Sucess"))
}

func setMessage(message string) map[string]any {
	r := make(map[string]any)
	r["message"] = message

	return r
}
