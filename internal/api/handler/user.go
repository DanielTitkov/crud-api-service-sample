package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/DanielTitkov/crud-api-service-sample/internal/api/model"
	"github.com/DanielTitkov/crud-api-service-sample/internal/api/util"
	"github.com/DanielTitkov/crud-api-service-sample/internal/domain"
)

func (h *Handler) CreateUserHandler(c echo.Context) error {
	request := new(model.CreateUserRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	err := h.app.CreateUser(&domain.User{
		Password:    request.Password,
		Email:       request.Email,
		Goals:       &request.Goals,
		Tags:        request.Tags,
		PublicName:  request.PublicName,
		Gender:      request.Gender,
		Age:         request.Age,
		Country:     request.Country,
		City:        request.City,
		Contact:     request.Contact,
		Description: request.Description,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to create user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "user created",
	})
}

func (h *Handler) UpdateUserHandler(c echo.Context) error {
	request := new(model.UpdateUserRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	userID, err := util.UserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(),
		})
	}

	u, err := h.app.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get user",
			Error:   err.Error(),
		})
	}

	err = h.app.UpdateUser(&domain.User{
		ID:          u.ID,
		Email:       u.Email,
		Goals:       &request.Goals,
		Tags:        request.Tags,
		PublicName:  request.PublicName,
		Gender:      request.Gender,
		Age:         request.Age,
		Country:     request.Country,
		City:        request.City,
		Contact:     request.Contact,
		Description: request.Description,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to update user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "user updated",
	})
}

func (h *Handler) GetPublicUserHandler(c echo.Context) error {
	request := new(model.GetPublicUserRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	u, err := h.app.GetUserByID(request.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.GetUserResponse{
		ID:          u.ID,
		Tags:        u.Tags,
		Goals:       *u.Goals,
		PublicName:  u.PublicName,
		Gender:      u.Gender,
		Age:         u.Age,
		Country:     u.Country,
		City:        u.City,
		Contact:     u.Contact,
		Description: u.Description,
	})
}

func (h *Handler) GetUserHandler(c echo.Context) error {
	userID, err := util.UserIDFromToken(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "token is invalid",
			Error:   err.Error(),
		})
	}

	u, err := h.app.GetUserByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get user",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.GetUserResponse{
		ID:          u.ID,
		Email:       u.Email,
		Tags:        u.Tags,
		Goals:       *u.Goals,
		PublicName:  u.PublicName,
		Gender:      u.Gender,
		Age:         u.Age,
		Country:     u.Country,
		City:        u.City,
		Contact:     u.Contact,
		Description: u.Description,
	})
}

func (h *Handler) GetTokenHandler(c echo.Context) error {
	request := new(model.GetTokenRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	user := &domain.User{
		Email:    request.Email,
		Password: request.Password,
	}

	valid, err := h.app.ValidateUserPassword(user)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Message: "failed to authorize",
			Error:   err.Error(),
		})
	}
	if !valid {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Message: "password is invalid",
		})
	}

	token, err := h.app.GetUserToken(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorResponse{
			Message: "failed to get token",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.GetTokenResponse{
		Token: token,
	})
}

func (h *Handler) GetMatchesHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, model.OKResponse{
		Status:  "ok",
		Message: "match",
	})
}
