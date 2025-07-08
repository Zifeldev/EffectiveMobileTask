package service

import (
	"em/internal/model"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type Enricher struct {
	client *resty.Client
}

func NewEnricher() *Enricher {
	return &Enricher{
		client: resty.New(),
	}
}

func (e *Enricher) EnrichPeople(p *model.Person) error {
	name := p.Name

	var genderResp struct {
		Gender string `json:"gender"`
	}
	_, err := e.client.R().SetQueryParam("name", name).SetResult(&genderResp).Get("https://api.genderize.io")
	if err != nil {
		logrus.Printf("gender API error: %v", err)
	}
	p.Gender = model.Gender(genderResp.Gender)

	var ageResp struct {
		Age int `json:"age"`
	}
	_, err = e.client.R().SetQueryParam("name", name).SetResult(&ageResp).Get("https://api.agify.io")
	if err != nil {
		logrus.Printf("Age API error: %v", err)
	}
	p.Age = ageResp.Age

	var countryResp struct {
		Country []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}
	_, err = e.client.R().SetQueryParam("name", name).SetResult(&countryResp).Get("https://api.nationalize.io")
	if err != nil {
		logrus.Printf("country API error: %v", err)
	}
	if len(countryResp.Country) > 0 {
		p.CountryID = countryResp.Country[0].CountryID
	}

	return nil
}
