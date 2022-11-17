package main

import (
	post "backend/api"
	"backend/models"
	"context"
	"sync"
)

type postService struct {
	mux sync.Mutex
}

func (p *postService) SendingGet(ctx context.Context, params post.SendingGetParams) (post.SendingGetRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	s := new(models.Sending)
	sending, err := s.GetSendingByOrderID(params.OrderID)
	if err != nil {
		return &post.SendingGetApplicationJSONNotFound{ErrorMessage: err.Error()}, nil
	}

	po := new(models.PostOffice)
	cityByPostcode, err := po.FindCityByPostcode()
	if err != nil {
		return nil, err
	}

	response := new(post.SendingGetResponse)
	response.Type = post.SendingType(sending.Type)
	response.Status = post.SendingStatus(sending.Status)

	stages := []post.SendingStage{}
	for i := range sending.Stages {
		new_stage := new(post.SendingStage)
		new_stage.Name = post.SendingStageName(sending.Stages[i].Name)
		new_stage.Date = sending.Stages[i].Date
		new_stage.Postcode = post.AddressPostcode(sending.Stages[i].Postcode)
		new_stage.City = cityByPostcode[sending.Stages[i].Postcode]
		stages = append(stages, *new_stage)
	}
	response.Stages = stages

	return response, nil
}

func (p *postService) PostcodesByCityGet(ctx context.Context) (post.PostcodesByCityGetResponse, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	po := new(models.PostOffice)
	postcodesByCity, err := po.FindPostcodesByCity()
	if err != nil {
		return nil, err
	}

	response := make(map[string][]post.AddressPostcode)

	for key := range postcodesByCity {
		postcodesArray := []post.AddressPostcode{}
		for postcode := range postcodesByCity[key] {
			postcodesArray = append(postcodesArray, post.AddressPostcode(postcodesByCity[key][postcode]))
		}
		response[key] = postcodesArray
	}

	return response, nil
}
