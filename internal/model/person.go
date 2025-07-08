package model

import "time"

type Gender string

const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderUnknown Gender = "unknown"
)

type Person struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	Gender     Gender    `json:"gender"`
	Age        int       `json:"age"`
	CountryID  string    `json:"country_id"`
	CreatedAt  time.Time `json:"created_at"`
}
