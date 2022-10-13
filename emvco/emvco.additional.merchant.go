package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
)

const (
	Merchant_Category_Code = "52"
	Merchant_Country_Code  = "58"
	Merchant_Name          = "59"
	Merchant_City          = "60"
	Merchant_Postal_Code   = "61"
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

	err := emv.parseMerchantCity(objectsMap)
	if err != nil {
		return err
	}

	err = emv.parseMerchantPostalCode(objectsMap)
	if err != nil {
		return err
	}

	return nil
}

func (emv *EMVCo) parseMerchantCity(objectsMap map[string]*objects.DataObject) error {
	// Merchant_City is mandatory field
	if objectsMap[Merchant_City] == nil {
		return errors.New("missing merchant city")
	}

	if objectsMap[Merchant_City].Length > 15 {
		return errors.New("invalid length of merchant city")
	}

	emv.mCity = &objects.MerchantCity{
		DataObject: objects.DataObject{
			ID:     Merchant_City,
			Length: objectsMap[Merchant_City].Length,
			Value:  objectsMap[Merchant_City].Value,
		},
	}

	return nil
}

func (emv *EMVCo) parseMerchantPostalCode(objectsMap map[string]*objects.DataObject) error {
	// Merchant_Postal_Code is optional
	if objectsMap[Merchant_Postal_Code] != nil {
		if objectsMap[Merchant_Postal_Code].Length > 10 {
			return errors.New("invalid length of merchant postal code")
		}

		emv.mPostalCode = &objects.MerchantPostalCode{
			DataObject: objects.DataObject{
				ID:     Merchant_Postal_Code,
				Length: objectsMap[Merchant_Postal_Code].Length,
				Value:  objectsMap[Merchant_Postal_Code].Value,
			},
		}
	}

	return nil
}
