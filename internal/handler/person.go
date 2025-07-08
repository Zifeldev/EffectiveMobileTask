package handler

import (
	"em/internal/model"
	"em/internal/repository"
	"em/internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PersonHandler struct {
	repo *repository.PersonRepository
	enricher *service.Enricher
}

func NewPersonHandler(repo *repository.PersonRepository, enricher *service.Enricher) *PersonHandler {
	return &PersonHandler{
		repo: repo,
		enricher: enricher,
	}
}

// GetPeople godoc
// @Summary Get list of people
// @Description With filter and pagination
// @Tags people
// @Accept json
// @Produce json
// @Param name query string false "Name"
// @Param surname query string false "Surname"
// @Param gender query string false "Gender" Enums(male, female, unknown)
// @Param country_id query string false "country ID"
// @Param min_age query int false "Min age"
// @Param max_age query int false "Max age"
// @Param page query int false "page number" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {array} model.Person
// @Failure 500 {object} model.ErrorResponse
// @Router /people [get]
func (h *PersonHandler) GetPeople(c *gin.Context) {
	var filter model.FilterPeopleRequest

	filter.Name = c.Query("name")
	filter.Surname = c.Query("surname")
	filter.Gender = c.Query("gender")
	filter.CountryID = c.Query("country_id")

	if v := c.Query("min_age"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			filter.MinAge = i
		}
	}
	if v := c.Query("max_age"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			filter.MaxAge = i
		}
	}
	if v := c.Query("page"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			filter.Page = i
		}
	}
	if v := c.Query("limit"); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			filter.Limit = i
		}
	}

	people, err := h.repo.GetPeople(c.Request.Context(), filter)
	if err != nil {
		log.Printf("error getting people: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get people"})
		return
	}

	c.JSON(http.StatusOK, people)
}

// CreatePerson godoc
// @Summary Create new person
// @Description Save and enrich
// @Tags people
// @Accept json
// @Produce json
// @Param person body model.CreatePersonRequest true "User data"
// @Success 201 {object} model.Person
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /people [post]
func (h *PersonHandler) CreatePerson(c *gin.Context){
	var req model.CreatePersonRequest
	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid/bad requset"})
		return
	}

	person := model.Person{
		Name: req.Name,
		Surname: req.Surname,
		Patronymic: req.Patronymic,
	}

	if err := h.enricher.EnrichPeople(&person); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enrich data"})
		return
	}

	err := h.repo.CreatePerson(c.Request.Context(), &person)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to save person"})
		return
	}
	
	c.JSON(http.StatusCreated, &person)
}