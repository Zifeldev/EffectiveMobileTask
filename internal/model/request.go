package model

type CreatePersonRequest struct {
	Name       string `json:"name" binding:"required"`
	Surname    string `json:"surname" binding:"required"`
	Patronymic string `json:"patronymic"`
}

type FilterPeopleRequest struct {
	Name      string `form:"name"`
	Surname   string `form:"surname"`
	Gender    string `form:"gender"`
	CountryID string `form:"country_id"`
	MinAge    int    `form:"min_age"`
	MaxAge    int    `form:"max_age"`
	Page      int    `form:"page" binding:"min=1"`
	Limit     int    `form:"limit" binding:"min=1,max=100"`
}
