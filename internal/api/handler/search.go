package handler

import (
	"fmt"
	"net/http"

	"github.com/DanielTitkov/crud-api-service-sample/internal/domain"

	"github.com/labstack/echo"
	"github.com/DanielTitkov/crud-api-service-sample/internal/api/model"
)

func (h *Handler) SearchHandler(c echo.Context) error {
	request := new(model.SearchRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	if h.cfg.Env == "dev" { // FIXME
		fmt.Printf("%+v\n", request)
		for _, t := range request.Tags {
			fmt.Printf("%+v\n", t)
		}
	}

	if len(request.Tags) > h.cfg.App.SearchTagsLimit {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "search failed",
			Error:   fmt.Sprintf("max amount of tags is %d", h.cfg.App.SearchTagsLimit),
		})
	}

	if len(request.Exclude) > h.cfg.App.SearchTagsLimit {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "search failed",
			Error:   fmt.Sprintf("max amount of exclude tags is %d", h.cfg.App.SearchTagsLimit),
		})
	}

	if request.Page < 0 {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Message: "search failed",
			Error:   fmt.Sprintf("invalid page number %d", request.Page),
		})
	}
	page := request.Page
	if request.Page == 0 {
		page = 1
	}

	result, err := h.app.Search(&domain.SearchArgs{
		Tags:    request.Tags,
		Exclude: request.Exclude,
		Goals:   request.Goals,
		Strict:  request.Strict,
		Limit:   h.cfg.App.DefaultResultsOnPage, // TODO: add max results
		Offset:  (page - 1) * h.cfg.App.DefaultResultsOnPage,
	})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorResponse{
			Message: "search failed",
			Error:   err.Error(),
		})
	}

	var users []model.SearchResponseUser
	for _, u := range result.Result {
		goals := &domain.UserGoals{}
		if u.Goals != nil {
			goals = u.Goals
		}

		users = append(users, model.SearchResponseUser{
			ID:          u.ID,
			Relevance:   u.Relevance,
			Tags:        u.Tags,
			Goals:       *goals,
			PublicName:  u.PublicName,
			Gender:      u.Gender,
			Age:         u.Age,
			Country:     u.Country,
			City:        u.City,
			Contact:     u.Contact,
			Description: u.Description,
		})
	}

	return c.JSON(http.StatusOK, model.SearchResponse{
		Elapsed:       result.Elapsed,
		Total:         result.Total,
		ResultsOnPage: result.ResultsOnPage,
		Result:        users,
	})
}
