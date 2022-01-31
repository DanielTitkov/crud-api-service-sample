package main

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"

	"github.com/DanielTitkov/crud-api-service-sample/cmd/app/prepare"
	"github.com/DanielTitkov/crud-api-service-sample/internal/app"
	"github.com/DanielTitkov/crud-api-service-sample/internal/configs"
	"github.com/DanielTitkov/crud-api-service-sample/internal/logger"
	"github.com/DanielTitkov/crud-api-service-sample/internal/repository/entgo"
	"github.com/DanielTitkov/crud-api-service-sample/internal/repository/entgo/ent"
	"github.com/DanielTitkov/crud-api-service-sample/internal/repository/entgo/ent/enttest"

	_ "github.com/mattn/go-sqlite3"
)

func setupServer(t *testing.T) (*ent.Client, *gin.Engine) {
	db := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	logger := logger.NewLogger("dev")
	repo := entgo.NewEntgoRepository(db, logger)
	cfg := configs.Config{Env: "dev"}
	app, _ := app.NewApp(cfg, logger, repo)
	server := prepare.NewServer(cfg, logger, app)
	return db, server
}

func TestCreateAndGet(t *testing.T) {
	db, server := setupServer(t)
	defer db.Close()

	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "Fiery",
			"price": 77700,
			"dougn": "thick",
			"description": "as spicy as expensive"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	apitest.New().
		Handler(server).
		Post("/api/v1/getPizzaByID").
		Body(`{
			"id": 1
		}`).
		Expect(t).
		Assert(jsonpath.Equal(`$.title`, "Fiery")).
		Assert(jsonpath.Equal(`$.price`, float64(77700))).
		Assert(jsonpath.Equal(`$.dough`, "thick")).
		Assert(jsonpath.Equal(`$.description`, "as spicy as expensive")).
		Status(http.StatusOK).
		End()
}

func TestFailedCreate(t *testing.T) {
	db, server := setupServer(t)
	defer db.Close()

	// not enough data
	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "P123"
		}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// bad dough
	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "P345",
			"price": 345,
			"dough": "bold"
		}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

	// setup for duplication test
	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "Fiery",
			"price": 77700,
			"dougn": "thick"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	// duplication
	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "Fiery",
			"price": 77700,
			"dougn": "thick"
		}`).
		Expect(t).
		Status(http.StatusBadRequest).
		End()
}

func TestCreateAndUpdate(t *testing.T) {
	db, server := setupServer(t)
	defer db.Close()

	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "Fiery",
			"price": 77700,
			"dougn": "thick",
			"description": "as spicy as expensive"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	apitest.New().
		Handler(server).
		Post("/api/v1/updatePizza").
		Body(`{
			"id": 1,
			"price": 777,
			"dough": "thin",
			"description": "whoa!"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	apitest.New().
		Handler(server).
		Post("/api/v1/getPizzaByID").
		Body(`{
			"id": 1
		}`).
		Expect(t).
		Assert(jsonpath.Equal(`$.price`, float64(777))).
		Assert(jsonpath.Equal(`$.dough`, "thin")).
		Assert(jsonpath.Equal(`$.description`, "whoa!")).
		Status(http.StatusOK).
		End()
}

func TestCreateAndGetAll(t *testing.T) {
	db, server := setupServer(t)
	defer db.Close()

	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "P1",
			"price": 111,
			"dough": "thick"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "P2",
			"price": 222,
			"dough": "thin"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	apitest.New().
		Handler(server).
		Post("/api/v1/getPizzas").
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len("$.pizzas", 2)).
		End()
}

func TestCreateAndDelete(t *testing.T) {
	db, server := setupServer(t)
	defer db.Close()

	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "P1",
			"price": 111,
			"dough": "thick"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	apitest.New().
		Handler(server).
		Post("/api/v1/createPizza").
		Body(`{
			"title": "P2",
			"price": 222,
			"dough": "thin"
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	apitest.New().
		Handler(server).
		Post("/api/v1/deletePizzaByID").
		Body(`{
			"id": 1
		}`).
		Expect(t).
		Status(http.StatusOK).
		End()

	apitest.New().
		Handler(server).
		Post("/api/v1/getPizzas").
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Len("$.pizzas", 1)).
		End()
}
