package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
)

const (
	Merchant_Category_Code = "52"
	Merchant_Country_Code  = "58"
	Merchant_Name          = "59"
)

func (emv *EMVCo) parseAdditionalMerchant(objectsMap map[string]*objects.DataObject) error {

	// Merchant_Category_Code is mandatory field
	if objectsMap[Merchant_Category_Code] == nil {
		return errors.New("missing merchant category code")
	}

	if objectsMap[Merchant_Category_Code].Length != 4 {
		return errors.New("invalid length of merchant category code")
	}

	emv.mCategoryCode = &objects.MerchantCategoryCode{
		DataObject: objects.DataObject{
			ID:     Merchant_Category_Code,
			Length: objectsMap[Merchant_Category_Code].Length,
			Value:  objectsMap[Merchant_Category_Code].Value,
		},
	}

	// Merchant_Country_Code is mandatory field
	if objectsMap[Merchant_Country_Code] == nil {
		return errors.New("missing merchant country code")
	}

	if objectsMap[Merchant_Country_Code].Length != 2 {
		return errors.New("invalid length of merchant country code")
	}

	emv.mCountryCode = &objects.MerchantCountryCode{
		DataObject: objects.DataObject{
			ID:     Merchant_Country_Code,
			Length: objectsMap[Merchant_Country_Code].Length,
			Value:  objectsMap[Merchant_Country_Code].Value,
		},
	}

	// Merchant_Name is mandatory field
	if objectsMap[Merchant_Name] == nil {
		return errors.New("missing merchant name")
	}

	if objectsMap[Merchant_Name].Length > 25 {
		return errors.New("invalid length of merchant name")
	}

	emv.mName = &objects.MerchantName{
		DataObject: objects.DataObject{
			ID:     Merchant_Name,
			Length: objectsMap[Merchant_Name].Length,
			Value:  objectsMap[Merchant_Name].Value,
		},
	}

	return nil
}
