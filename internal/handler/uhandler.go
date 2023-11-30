package handler

import (
	"errors"
	"github.com/Sokol111/educational-go/internal/model"
	"github.com/Sokol111/educational-go/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type createUserRequest struct {
	Name    string `json:"name" binding:"required"`
	Enabled bool   `json:"enabled"`
}

func (r *createUserRequest) validate() error {
	return nil
}

type updateUserRequest struct {
	Id      int64  `json:"id" binding:"required"`
	Version int32  `json:"version" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Enabled bool   `json:"enabled"`
}

type userResponse struct {
	Id      int64  `json:"id"`
	Version int32  `json:"version"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) BindRoutes(engine *gin.Engine) {
	group := engine.Group("/user")
	group.POST("/create", h.createUser)
	group.PUT("/update", h.updateUser)
	group.GET("/list", h.getAllUsers)
	group.GET("/:id", h.getUserById)
}

func (h *UserHandler) createUser(c *gin.Context) {
	var request createUserRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	if u, err := h.service.Create(c, model.User{Name: request.Name, Enabled: request.Enabled}); err == nil {
		c.JSON(200, toJson(&u))
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (h *UserHandler) getUserById(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		c.AbortWithError(http.StatusInternalServerError, errors.New("id parameter is empty"))
		return
	}
	intId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("wrong id parameter"))
		return
	}

	if u, err := h.service.GetById(c, intId); err == nil {
		c.JSON(200, toJson(&u))
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (h *UserHandler) getAllUsers(c *gin.Context) {
	if users, err := h.service.GetUsers(c); err == nil {
		response := make([]*userResponse, 0, len(users))
		for _, u := range users {
			response = append(response, toJson(&u))
		}
		c.JSON(200, response)
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func (h *UserHandler) updateUser(c *gin.Context) {
	var request updateUserRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	if u, err := h.service.Update(c, model.User{ID: request.Id, Version: request.Version, Name: request.Name, Enabled: request.Enabled}); err == nil {
		c.JSON(200, toJson(&u))
	} else {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
}

func toJson(u *model.User) *userResponse {
	return &userResponse{Id: u.ID, Version: u.Version, Name: u.Name, Enabled: u.Enabled}
}
