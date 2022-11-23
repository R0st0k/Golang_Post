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
	newSending["Weight"] = int64(req.Weight)

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

func (p *postService) ExportSendingFilter(params post.SendingFilterGetParams) map[string]interface{} {
	sendingFilter := make(map[string]interface{})
	sendingFilter["page"] = int64(params.Page)
	sendingFilter["elems"] = int64(params.ElemsOnPage)

	if OrderID, ok := params.Filter.OrderID.Get(); ok {
		sendingFilter["order_id"] = OrderID
	}
	if Type, ok := params.Filter.Type.Get(); ok {
		sendingFilter["type"] = string(Type)
	}
	if Status, ok := params.Filter.Status.Get(); ok {
		sendingFilter["status"] = string(Status)
	}
	if DateStart, ok := params.Filter.DateStart.Get(); ok {
		sendingFilter["date_start"] = DateStart
	}
	if DateFinish, ok := params.Filter.DateFinish.Get(); ok {
		sendingFilter["date_finish"] = DateFinish
	}
	if SenderSettlement, ok := params.Filter.SenderSettlement.Get(); ok {
		sendingFilter["sender_settlement"] = SenderSettlement
	}
	if ReceiverSettlement, ok := params.Filter.ReceiverSettlement.Get(); ok {
		sendingFilter["receiver_settlement"] = ReceiverSettlement
	}
	if Length, ok := params.Filter.Length.Get(); ok {
		sendingFilter["length"] = Length
	}
	if Width, ok := params.Filter.Width.Get(); ok {
		sendingFilter["width"] = Width
	}
	if Height, ok := params.Filter.Height.Get(); ok {
		sendingFilter["height"] = Height
	}
	if Weight, ok := params.Filter.Weight.Get(); ok {
		sendingFilter["weight"] = int64(Weight)
	}
	if SortType, ok := params.Sort.SortType.Get(); ok {
		sendingFilter["sort_type"] = string(SortType)
	}
	if SortField, ok := params.Sort.SortField.Get(); ok {
		sendingFilter["sort_field"] = string(SortField)
	}

	return sendingFilter
}

func (p *postService) ModelToSendingFilterGetResponseResultItem(sending models.Sending) (post.SendingFilterGetResponseResultItem, error) {
	var sendingItem post.SendingFilterGetResponseResultItem

	uuid, err := uuid.Parse(sending.OrderID)
	if err != nil {
		return post.SendingFilterGetResponseResultItem{}, err
	}

	sendingItem.SetOrderID(post.SendingOrderID(uuid))
	sendingItem.SetType(post.SendingType(sending.Type))
	sendingItem.SetDate(sending.RegistrationDate)
	sendingItem.Settlement.SetSender(sending.Sender.Address.Settlement)
	sendingItem.Settlement.SetReceiver(sending.Receiver.Address.Settlement)
	sendingItem.SetWeight(post.SendingWeight(sending.Weight))
	sendingItem.Size.SetLength(sending.Size.Length)
	sendingItem.Size.SetWidth(sending.Size.Width)
	sendingItem.Size.SetHeight(sending.Size.Height)
	sendingItem.SetStatus(post.SendingStatus(sending.Status))

	return sendingItem, nil
}

func (p *postService) SendingFilterGet(ctx context.Context, params post.SendingFilterGetParams) (post.SendingFilterGetRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	sendingFilter := p.ExportSendingFilter(params)

	s := new(models.Sending)
	total, resultSending, err := s.FilterSending(sendingFilter)
	if err != nil {
		return nil, err
	}

	var response post.SendingFilterGetResponse
	result := []post.SendingFilterGetResponseResultItem{}

	for i := range resultSending {
		item, err := p.ModelToSendingFilterGetResponseResultItem(resultSending[i])
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}

	response.Result = result
	response.Total = total

	return &response, nil
}
