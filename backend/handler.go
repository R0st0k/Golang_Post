package main

import (
	post "backend/api"
	"backend/models"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"os"
	"sync"
	"time"
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

	stages := []post.SendingGetResponseStagesItem{}
	for i := range sending.Stages {
		newStage := new(post.SendingGetResponseStagesItem)
		newStage.SetName(post.SendingGetResponseStagesItemName(sending.Stages[i].Name))
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

func (p *postService) ImportPostClient(params map[string]interface{}) (post.PostClient, error) {
	client := new(post.PostClient)
	if name, ok := params["name"]; ok {
		client.SetName(name.(string))
	}
	if surname, ok := params["surname"]; ok {
		client.SetSurname(surname.(string))
	}
	if middleName, ok := params["middle_name"]; ok {
		value := new(post.OptString)
		value.SetTo(middleName.(string))
		client.SetMiddleName(*value)
	}

	if value, ok := params["address"]; ok {
		address, err := p.ImportAddress(value.(map[string]interface{}))
		if err != nil {
			return post.PostClient{}, err
		}
		client.SetAddress(address)
	}

	return *client, nil
}

func (p *postService) ImportAddress(params map[string]interface{}) (post.Address, error) {
	address := new(post.Address)

	if value, ok := params["postcode"]; ok {
		address.SetPostcode(post.AddressPostcode(value.(string)))
	}
	if value, ok := params["region"]; ok {
		region := new(post.OptString)
		region.SetTo(value.(string))
		address.SetRegion(*region)
	}
	if value, ok := params["district"]; ok {
		district := new(post.OptString)
		district.SetTo(value.(string))
		address.SetDistrict(*district)
	}
	if value, ok := params["settlement"]; ok {
		address.SetSettlement(value.(string))
	}
	if value, ok := params["street"]; ok {
		address.SetStreet(value.(string))
	}
	if value, ok := params["building"]; ok {
		address.SetBuilding(value.(string))
	}
	if value, ok := params["apartment"]; ok {
		apartment := new(post.OptString)
		apartment.SetTo(value.(string))
		address.SetApartment(*apartment)
	}

	return *address, nil
}

func (p *postService) ImportStages(params []interface{}) ([]post.SendingStage, error) {
	answer := []post.SendingStage{}

	for _, stageJSON := range params {
		stage := new(post.SendingStage)
		if name, ok := stageJSON.(map[string]interface{})["name"]; ok {
			stage.SetName(post.SendingStageName(name.(string)))
		}
		if value, ok := stageJSON.(map[string]interface{})["timestamp"]; ok {
			if timestamp, ok := value.(map[string]interface{})["$date"]; ok {
				time, err := time.Parse(time.RFC3339, timestamp.(string))
				if err != nil {
					return []post.SendingStage{}, err
				}
				stage.SetTimestamp(post.SendingStageTimestamp{Date: time})
			}
		}
		if postcode, ok := stageJSON.(map[string]interface{})["postcode"]; ok {
			stage.SetPostcode(post.AddressPostcode(postcode.(string)))
		}
		if value, ok := stageJSON.(map[string]interface{})["employee_id"]; ok {
			if id, ok := value.(map[string]interface{})["$oid"]; ok {
				stage.SetEmployeeID(post.ObjectID{Oid: id.(string)})
			}
		}

		answer = append(answer, *stage)
	}

	return answer, nil
}

func (p *postService) ImportSendings(params []map[string]interface{}) ([]post.Sending, error) {
	answer := []post.Sending{}

	for _, json := range params {
		sending := new(post.Sending)

		if value, ok := json["_id"]; ok {
			if id, ok := value.(map[string]interface{})["$oid"]; ok {
				sending.SetID(post.ObjectID{Oid: id.(string)})
			}
		}
		if value, ok := json["order_id"]; ok {
			uuid, err := uuid.Parse(value.(string))
			if err != nil {
				return []post.Sending{}, err
			}
			sending.SetOrderID(post.SendingOrderID(uuid))
		}
		if value, ok := json["registration_date"]; ok {
			if date, ok := value.(map[string]interface{})["$date"]; ok {
				time, err := time.Parse(time.RFC3339, date.(string))
				if err != nil {
					return []post.Sending{}, err
				}
				sending.SetRegistrationDate(post.SendingRegistrationDate{Date: time})
			}
		}
		if jsonSender, ok := json["sender"]; ok {
			sender, err := p.ImportPostClient(jsonSender.(map[string]interface{}))
			if err != nil {
				return []post.Sending{}, err
			}
			sending.SetSender(sender)
		}
		if jsonReceiver, ok := json["receiver"]; ok {
			receiver, err := p.ImportPostClient(jsonReceiver.(map[string]interface{}))
			if err != nil {
				return []post.Sending{}, err
			}
			sending.SetReceiver(receiver)
		}
		if value, ok := json["type"]; ok {
			sending.SetType(post.SendingType(value.(string)))
		}
		if size, ok := json["size"].(map[string]interface{}); ok {
			newSize := new(post.SendingSize)
			if length, ok := size["length"]; ok {
				newSize.SetLength(int64(length.(float64)))
			}
			if width, ok := size["width"]; ok {
				newSize.SetWidth(int64(width.(float64)))
			}
			if height, ok := size["height"]; ok {
				newSize.SetHeight(int64(height.(float64)))
			}
			sending.SetSize(*newSize)
		}
		if value, ok := json["weight"]; ok {
			sending.SetWeight(post.SendingWeight(int64(value.(float64))))
		}
		if value, ok := json["stages"]; ok {
			stages, err := p.ImportStages(value.([]interface{}))
			if err != nil {
				return []post.Sending{}, err
			}
			sending.SetStages(stages)
		}

		answer = append(answer, *sending)
	}

	return answer, nil
}

func (p *postService) DataExportSendingGet(ctx context.Context) ([]post.Sending, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	s := new(models.Sending)
	fileName, err := s.Export()
	if err != nil {
		return nil, err
	}
	defer os.Remove(fileName) // clean up

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		return nil, err
	}

	answer, err := p.ImportSendings(result)
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (p *postService) ExportAddress(params post.Address) (map[string]interface{}, error) {
	address := map[string]interface{}{}
	address["postcode"] = params.GetPostcode()
	if region, ok := params.GetRegion().Get(); ok {
		address["region"] = region
	}
	if district, ok := params.GetDistrict().Get(); ok {
		address["district"] = district
	}
	address["settlement"] = params.GetSettlement()
	address["street"] = params.GetStreet()
	address["building"] = params.GetBuilding()
	if apartment, ok := params.GetApartment().Get(); ok {
		address["apartment"] = apartment
	}

	return address, nil
}

func (p *postService) ExportPostClient(params post.PostClient) (map[string]interface{}, error) {
	client := map[string]interface{}{}
	client["name"] = params.GetName()
	client["surname"] = params.GetSurname()
	if middleName, ok := params.GetMiddleName().Get(); ok {
		client["middle_name"] = middleName
	}
	var err error
	client["address"], err = p.ExportAddress(params.Address)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (p *postService) ExportStages(params []post.SendingStage) ([]map[string]interface{}, error) {
	answer := []map[string]interface{}{}

	for _, stage := range params {
		json := map[string]interface{}{}
		json["name"] = string(stage.GetName())
		json["timestamp"] = map[string]string{"$date": stage.GetTimestamp().GetDate().Format(time.RFC3339)}
		json["postcode"] = string(stage.GetPostcode())
		json["employee_id"] = map[string]string{"$oid": stage.GetEmployeeID().GetOid()}
		answer = append(answer, json)
	}

	return answer, nil
}

func (p *postService) ExportSending(params []post.Sending) ([]map[string]interface{}, error) {
	answer := []map[string]interface{}{}
	var err error

	for _, sending := range params {
		json := map[string]interface{}{}
		json["_id"] = map[string]string{"$oid": sending.GetID().GetOid()}
		json["order_id"] = uuid.UUID(sending.GetOrderID()).String()
		json["registration_date"] = map[string]string{"$date": sending.GetRegistrationDate().GetDate().Format(time.RFC3339)}
		json["sender"], err = p.ExportPostClient(sending.GetSender())
		if err != nil {
			return nil, err
		}
		json["receiver"], err = p.ExportPostClient(sending.GetReceiver())
		if err != nil {
			return nil, err
		}
		json["type"] = string(sending.GetType())
		{
			size := map[string]interface{}{}
			size["length"] = sending.GetSize().GetLength()
			size["width"] = sending.GetSize().GetWidth()
			size["height"] = sending.GetSize().GetHeight()
			json["size"] = size
		}
		json["weight"] = int64(sending.GetWeight())
		json["stages"], err = p.ExportStages(sending.GetStages())
		if err != nil {
			return nil, err
		}

		answer = append(answer, json)
	}

	return answer, nil
}

func (p *postService) DataImportSendingPost(ctx context.Context, req []post.Sending) (post.DataImportSendingPostRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	data, err := p.ExportSending(req)
	if err != nil {
		return nil, err
	}

	f, err := os.CreateTemp("", "sendings")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	asJSON, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return nil, err
	}

	_, err = f.Write(asJSON)
	if err != nil {
		return nil, err
	}

	err = f.Close()
	if err != nil {
		return nil, err
	}

	s := new(models.Sending)
	err = s.Import(f.Name())
	if err != nil {
		return nil, err
	}

	return new(post.DataImportSendingPostOK), nil
}
