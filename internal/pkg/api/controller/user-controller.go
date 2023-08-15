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

func (uc *UserController) FindAll(c *gin.Context) {
	users, err := uc.Implementation.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, setMessage("Something went wrong. Error: "+err.Error()))
		return
	}

	if len(users) > 0 {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNoContent, users)
	}
}

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
