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
	response.SetType(post.SendingType(sending.Type))
	response.SetStatus(post.SendingStatus(sending.Status))

	stages := []post.SendingStage{}
	for i := range sending.Stages {
		newStage := new(post.SendingStage)
		newStage.SetName(post.SendingStageName(sending.Stages[i].Name))
		newStage.SetDate(sending.Stages[i].Timestamp)
		newStage.SetPostcode(post.AddressPostcode(sending.Stages[i].Postcode))
		newStage.SetSettlement(settlementByPostcode[sending.Stages[i].Postcode])
		stages = append(stages, *newStage)
	}
	response.SetStages(stages)

	return response, nil
}

func (p *postService) ExportClient(client post.PostClient) map[string]string {
	newClient := make(map[string]string)
	newClient["Name"] = client.GetName()
	newClient["Surname"] = client.GetSurname()
	newClient["MiddleName"], _ = client.MiddleName.Get()
	newClient["Postcode"] = string(client.Address.GetPostcode())
	newClient["Reqion"], _ = client.Address.Region.Get()
	newClient["District"], _ = client.Address.District.Get()
	newClient["Settlement"] = client.Address.GetSettlement()
	newClient["Street"] = client.Address.GetStreet()
	newClient["Building"] = client.Address.GetBuilding()
	newClient["Apartment"], _ = client.Address.Apartment.Get()

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
	response.SetOrderID(post.SendingOrderID(orderID))

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

	if OrderID, ok := params.OrderID.Get(); ok {
		sendingFilter["order_id"] = OrderID
	}
	if len(params.Type) > 0 {
		result := []string{}
		for _, data := range params.Type {
			result = append(result, string(data))
		}
		sendingFilter["type"] = result
	}
	if len(params.Status) > 0 {
		result := []string{}
		for _, data := range params.Status {
			result = append(result, string(data))
		}
		sendingFilter["status"] = result
	}
	if Date, ok := params.Date.Get(); ok {
		if DateStart, ok := Date.GetDateStart().Get(); ok {
			sendingFilter["date_start"] = DateStart
		}
		if DateFinish, ok := Date.GetDateFinish().Get(); ok {
			sendingFilter["date_finish"] = DateFinish
		}
	}
	if Settlements, ok := params.Settlements.Get(); ok {
		if SenderSettlement, ok := Settlements.GetSenderSettlement().Get(); ok {
			sendingFilter["sender_settlement"] = SenderSettlement
		}
		if ReceiverSettlement, ok := Settlements.GetReceiverSettlement().Get(); ok {
			sendingFilter["receiver_settlement"] = ReceiverSettlement
		}
	}
	if Length, ok := params.Length.Get(); ok {
		if LengthMin, ok := Length.GetLengthMin().Get(); ok {
			sendingFilter["length_min"] = LengthMin
		}
		if LengthMax, ok := Length.GetLengthMax().Get(); ok {
			sendingFilter["length_max"] = LengthMax
		}
	}
	if Width, ok := params.Width.Get(); ok {
		if WidthMin, ok := Width.GetWidthMin().Get(); ok {
			sendingFilter["width_min"] = WidthMin
		}
		if WidthMax, ok := Width.GetWidthMax().Get(); ok {
			sendingFilter["width_max"] = WidthMax
		}
	}
	if Height, ok := params.Height.Get(); ok {
		if HeightMin, ok := Height.GetHeightMin().Get(); ok {
			sendingFilter["height_min"] = HeightMin
		}
		if HeightMax, ok := Height.GetHeightMax().Get(); ok {
			sendingFilter["height_max"] = HeightMax
		}
	}
	if Weight, ok := params.Weight.Get(); ok {
		if WeightMin, ok := Weight.GetWeightMin().Get(); ok {
			sendingFilter["weight_min"] = int64(WeightMin)
		}
		if WeightMax, ok := Weight.GetWeightMax().Get(); ok {
			sendingFilter["weight_max"] = int64(WeightMax)
		}
	}
	if Sort, ok := params.Sort.Get(); ok {
		if SortType, ok := Sort.GetSortType().Get(); ok {
			sendingFilter["sort_type"] = string(SortType)
		}
		if SortField, ok := Sort.GetSortField().Get(); ok {
			sendingFilter["sort_field"] = string(SortField)
		}
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

	response.SetTotal(total)
	response.SetResult(result)

	return &response, nil
}
