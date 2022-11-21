package main

import (
	post "backend/api"
	"backend/models"
	"context"
	"github.com/google/uuid"
	"sync"
)

type postService struct {
	mux sync.Mutex
}

func (p *postService) SendingGet(ctx context.Context, params post.SendingGetParams) (post.SendingGetRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	s := new(models.Sending)
	sending, err := s.GetSendingByOrderID(uuid.UUID(params.OrderID))
	if err != nil {
		return &post.SendingGetApplicationJSONNotFound{ErrorMessage: err.Error()}, nil
	}

	po := new(models.PostOffice)
	settlementByPostcode, err := po.GetSettlementByPostcode()
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
		new_stage.Settlement = settlementByPostcode[sending.Stages[i].Postcode]
		stages = append(stages, *new_stage)
	}
	response.Stages = stages

	return response, nil
}

func (p *postService) ExportClient(client post.PostClient) map[string]string {
	newClient := make(map[string]string)
	newClient["Name"] = client.Name
	newClient["Surname"] = client.Surname
	newClient["MiddleName"] = client.MiddleName.Value
	newClient["Postcode"] = string(client.Address.Postcode)
	newClient["Reqion"] = client.Address.Region.Value
	newClient["District"] = client.Address.District.Value
	newClient["Settlement"] = client.Address.Settlement
	newClient["Street"] = client.Address.Street
	newClient["Building"] = client.Address.Building
	newClient["Apartment"] = client.Address.Apartment.Value

	return newClient
}

func (p *postService) SendingPost(ctx context.Context, req post.SendingPostReq) (post.SendingPostRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	newSending := make(map[string]interface{})
	newSending["Type"] = string(req.Type)
	newSending["Sender"] = p.ExportClient(req.Sender)
	newSending["Receiver"] = p.ExportClient(req.Receiver)
	newSending["Length"] = req.Size.Length
	newSending["Width"] = req.Size.Width
	newSending["Height"] = req.Size.Height
	newSending["Weight"] = req.Weight

	s := new(models.Sending)
	orderID, err := s.InsertNewSending(newSending)
	if err != nil {
		return nil, err
	}

	response := new(post.SendingPostResponse)
	response.OrderID = post.SendingOrderID(orderID)

	return response, nil

}

func (p *postService) PostcodesBySettlementGet(ctx context.Context) (post.PostcodesBySettlementGetResponse, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	po := new(models.PostOffice)
	postcodesBySettlement, err := po.GetPostcodesBySettlement()
	if err != nil {
		return nil, err
	}

	response := make(map[string][]post.AddressPostcode)

	for key := range postcodesBySettlement {
		postcodesArray := []post.AddressPostcode{}
		for postcode := range postcodesBySettlement[key] {
			postcodesArray = append(postcodesArray, post.AddressPostcode(postcodesBySettlement[key][postcode]))
		}
		response[key] = postcodesArray
	}

	return response, nil
}
