package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
	"strconv"
)

const (
	Additional_Data_Field_Template   = "62"
	Bill_Number                      = "01"
	Mobile_Number                    = "02"
	Store_Label                      = "03"
	Loyalty_Number                   = "04"
	Reference_Label                  = "05"
	Customer_Label                   = "06"
	Terminal_Label                   = "07"
	Purpose_Of_Transaction           = "08"
	Additional_Consumer_Data_Request = "09"
	Merchant_Tax_ID                  = "10"
	Merchant_Channel                 = "11"
)

func (emv *EMVCo) parseAdditionalDataFieldTemplate(objectsMap map[string]*objects.DataObject) error {
	emv.dataFieldTemplate = &objects.AdditionalDataFieldTemplate{}

	if objectsMap[Additional_Data_Field_Template] != nil {
		if objectsMap[Additional_Data_Field_Template].Length > 99 {
			return errors.New("invalid length of Additional Data Field Template")
		}

		dataObjects, err := emv.parseDataObjects(objectsMap[Additional_Data_Field_Template].Value.(string), false)
		if err != nil {
			return err
		}

		emv.dataFieldTemplate.DataObject = objects.DataObject{
			ID:          Additional_Data_Field_Template,
			Length:      objectsMap[Additional_Data_Field_Template].Length,
			Value:       objectsMap[Additional_Data_Field_Template].Value,
			TemplateMap: dataObjects,
		}

		err = emv.parseBillNumber(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseMobileNumber(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseStoreLabel(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseLoyaltyNumber(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseReferenceLabel(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseCustomerLabel(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseTerminalLabel(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parsePurposeOfTransaction(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseAdditionalConsumerDataRequest(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseMerchantTaxID(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseMerchantChannel(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseRFU(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parsePaymentSystemSpecificTemplates(dataObjects)
		if err != nil {
			return err
		}

	}
	return nil
}

func (emv *EMVCo) parseBillNumber(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Bill_Number] != nil {
		if objectsMap[Bill_Number].Length > 25 {
			return errors.New("invalid length of Bill Number")
		}

		emv.dataFieldTemplate.BillNumber = &objects.AdditionalObjectBillNumber{
			DataObject: objects.DataObject{
				ID:     Bill_Number,
				Length: objectsMap[Bill_Number].Length,
				Value:  objectsMap[Bill_Number].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseMobileNumber(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Mobile_Number] != nil {
		if objectsMap[Mobile_Number].Length > 25 {
			return errors.New("invalid length of Mobile Number")
		}

		emv.dataFieldTemplate.MobileNumber = &objects.AdditionalObjectMobileNumber{
			DataObject: objects.DataObject{
				ID:     Mobile_Number,
				Length: objectsMap[Mobile_Number].Length,
				Value:  objectsMap[Mobile_Number].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseStoreLabel(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Store_Label] != nil {
		if objectsMap[Store_Label].Length > 25 {
			return errors.New("invalid length of Store Label")
		}

		emv.dataFieldTemplate.StoreLabel = &objects.AdditionalObjectStoreLabel{
			DataObject: objects.DataObject{
				ID:     Store_Label,
				Length: objectsMap[Store_Label].Length,
				Value:  objectsMap[Store_Label].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseLoyaltyNumber(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Loyalty_Number] != nil {
		if objectsMap[Loyalty_Number].Length > 25 {
			return errors.New("invalid length of Loyalty Number")
		}

		emv.dataFieldTemplate.LoyaltyNumber = &objects.AdditionalObjectLoyaltyNumber{
			DataObject: objects.DataObject{
				ID:     Loyalty_Number,
				Length: objectsMap[Loyalty_Number].Length,
				Value:  objectsMap[Loyalty_Number].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseReferenceLabel(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Reference_Label] != nil {
		if objectsMap[Reference_Label].Length > 25 {
			return errors.New("invalid length of Reference Label")
		}

		emv.dataFieldTemplate.ReferenceLabel = &objects.AdditionalObjectReferenceLabel{
			DataObject: objects.DataObject{
				ID:     Reference_Label,
				Length: objectsMap[Reference_Label].Length,
				Value:  objectsMap[Reference_Label].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseCustomerLabel(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Customer_Label] != nil {
		if objectsMap[Customer_Label].Length > 25 {
			return errors.New("invalid length of Customer Label")
		}

		emv.dataFieldTemplate.CustomerLabel = &objects.AdditionalObjectCustomerLabel{
			DataObject: objects.DataObject{
				ID:     Customer_Label,
				Length: objectsMap[Customer_Label].Length,
				Value:  objectsMap[Customer_Label].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseTerminalLabel(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Terminal_Label] != nil {
		if objectsMap[Terminal_Label].Length > 25 {
			return errors.New("invalid length of Terminal Label")
		}

		emv.dataFieldTemplate.TerminalLabel = &objects.AdditionalObjectTerminalLabel{
			DataObject: objects.DataObject{
				ID:     Terminal_Label,
				Length: objectsMap[Terminal_Label].Length,
				Value:  objectsMap[Terminal_Label].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parsePurposeOfTransaction(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Purpose_Of_Transaction] != nil {
		if objectsMap[Purpose_Of_Transaction].Length > 25 {
			return errors.New("invalid length of Purpose of Transaction")
		}

		emv.dataFieldTemplate.PurposeOfTnx = &objects.AdditionalObjectPurposeOfTransaction{
			DataObject: objects.DataObject{
				ID:     Purpose_Of_Transaction,
				Length: objectsMap[Purpose_Of_Transaction].Length,
				Value:  objectsMap[Purpose_Of_Transaction].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseAdditionalConsumerDataRequest(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Additional_Consumer_Data_Request] != nil {
		if objectsMap[Additional_Consumer_Data_Request].Length > 3 {
			return errors.New("invalid length of Additional Consumer Data Request")
		}

		emv.dataFieldTemplate.AdditionalConsumerDataRequest = &objects.AdditionalObjectConsumerDataRequest{
			DataObject: objects.DataObject{
				ID:     Additional_Consumer_Data_Request,
				Length: objectsMap[Additional_Consumer_Data_Request].Length,
				Value:  objectsMap[Additional_Consumer_Data_Request].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseMerchantTaxID(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Merchant_Tax_ID] != nil {
		if objectsMap[Merchant_Tax_ID].Length > 20 {
			return errors.New("invalid length of Merchant Tax ID")
		}

		emv.dataFieldTemplate.MerchantTaxID = &objects.AdditionalObjectMerchantTaxID{
			DataObject: objects.DataObject{
				ID:     Merchant_Tax_ID,
				Length: objectsMap[Merchant_Tax_ID].Length,
				Value:  objectsMap[Merchant_Tax_ID].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseMerchantChannel(objectsMap map[string]*objects.DataObject) error {

	if objectsMap[Merchant_Channel] != nil {
		if objectsMap[Merchant_Channel].Length > 3 {
			return errors.New("invalid length of Merchant Channel")
		}

		emv.dataFieldTemplate.MerchantChannel = &objects.AdditionalObjectMerchantChannel{
			DataObject: objects.DataObject{
				ID:     Merchant_Channel,
				Length: objectsMap[Merchant_Channel].Length,
				Value:  objectsMap[Merchant_Channel].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseRFU(objectsMap map[string]*objects.DataObject) error {

	rfus := make([]*objects.AdditionalObjectRFU, 0)
	for i := 12; i <= 49; i++ {
		if objectsMap[strconv.Itoa(i)] != nil {

			dataObjects, err := emv.parseDataObjects(objectsMap[strconv.Itoa(i)].Value.(string), false)
			if err != nil {
				return err
			}

			rfus = append(rfus, &objects.AdditionalObjectRFU{
				DataObject: objects.DataObject{
					ID:          strconv.Itoa(i),
					Length:      objectsMap[strconv.Itoa(i)].Length,
					Value:       objectsMap[strconv.Itoa(i)].Value,
					TemplateMap: dataObjects,
				},
			})

		}

	}

	emv.dataFieldTemplate.RFU = rfus
	return nil
}

func (emv *EMVCo) parsePaymentSystemSpecificTemplates(objectsMap map[string]*objects.DataObject) error {

	templates := make([]*objects.AdditionalObjectPaymentSystemSpecificTemplate, 0)
	for i := 50; i <= 99; i++ {
		if objectsMap[strconv.Itoa(i)] != nil {

			dataObjects, err := emv.parseDataObjects(objectsMap[strconv.Itoa(i)].Value.(string), false)
			if err != nil {
				return err
			}

			templates = append(templates, &objects.AdditionalObjectPaymentSystemSpecificTemplate{
				DataObject: objects.DataObject{
					ID:          strconv.Itoa(i),
					Length:      objectsMap[strconv.Itoa(i)].Length,
					Value:       objectsMap[strconv.Itoa(i)].Value,
					TemplateMap: dataObjects,
				},
			})

		}

	}

	emv.dataFieldTemplate.PaymentSystemSpecificTemplates = templates
	return nil
}
