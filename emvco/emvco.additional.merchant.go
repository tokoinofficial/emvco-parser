package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
	"strconv"
)

const (
	Merchant_Category_Code      = "52"
	Merchant_Country_Code       = "58"
	Merchant_Name               = "59"
	Merchant_City               = "60"
	Merchant_Postal_Code        = "61"
	Merchant_Alternate_Language = "64"
)

func (emv *EMVCo) parseAdditionalMerchant(objectsMap map[string]*objects.DataObject) error {

	// Merchant_Category_Code is mandatory field
	if objectsMap[Merchant_Category_Code] == nil {
		return errors.New("missing Merchant Category Code")
	}

	if objectsMap[Merchant_Category_Code].Length != 4 {
		return errors.New("invalid length of Merchant Category Code")
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

	err = emv.parseMerchantAlternateLanguage(objectsMap)
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
			return errors.New("invalid length of Merchant Postal Code")
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

// The Merchant Information— Language Template includes merchant information in an alternate language and may use a character set different from the Common Character Set. It provides an alternative to the merchant information under the root.
// For the list of data objects that can be included in this template, please refer to Table 3.8.
const (
	Merchant_Language_Preference     = "00"
	Merchant_Name_Alternate_Language = "01"
	Merchant_City_Alternate_Language = "02"
)

func (emv *EMVCo) parseMerchantAlternateLanguage(objectsMap map[string]*objects.DataObject) error {
	// Merchant_Alternate_Language is optional
	if objectsMap[Merchant_Alternate_Language] != nil {

		if objectsMap[Merchant_Alternate_Language].Length > 99 {
			return errors.New("invalid length of Merchant Alternate Language")
		}

		emv.mAlternateLanguage = &objects.MerchantAlternateLanguage{}
		dataObjects, err := emv.parseDataObjects(objectsMap[Merchant_Alternate_Language].Value.(string), false)
		if err != nil {
			return err
		}

		err = emv.parseMerchantLanguagePreference(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseMerchantNameAlternateLanguage(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseMerchantCityAlternateLanguage(dataObjects)
		if err != nil {
			return err
		}

		err = emv.parseRFUAlternateLanguage(dataObjects)
		if err != nil {
			return err
		}

		emv.mAlternateLanguage.TemplateMap = dataObjects

	}

	return nil
}

func (emv *EMVCo) parseMerchantLanguagePreference(objectsMap map[string]*objects.DataObject) error {
	// Language Preference is mandatory
	if objectsMap[Merchant_Language_Preference] == nil {
		return errors.New("missing Language Preference")
	}

	if objectsMap[Merchant_Language_Preference].Length != 2 {
		return errors.New("invalid length of Language Preference")
	}

	emv.mAlternateLanguage.LanguagePreference = &objects.MerchantLanguagePreference{
		DataObject: objects.DataObject{
			ID:     Merchant_Language_Preference,
			Length: objectsMap[Merchant_Language_Preference].Length,
			Value:  objectsMap[Merchant_Language_Preference].Value,
		},
	}

	return nil
}

func (emv *EMVCo) parseMerchantNameAlternateLanguage(objectsMap map[string]*objects.DataObject) error {
	// Merchant Name Alternate Language is mandatory
	if objectsMap[Merchant_Name_Alternate_Language] == nil {
		return errors.New("missing Merchant Name Alternate Language")
	}

	if objectsMap[Merchant_Name_Alternate_Language].Length > 25 {
		return errors.New("invalid length of Merchant Name Alternate Language")
	}

	emv.mAlternateLanguage.MerchantName = &objects.MerchantNameAlternateLanguage{
		DataObject: objects.DataObject{
			ID:     Merchant_Name_Alternate_Language,
			Length: objectsMap[Merchant_Name_Alternate_Language].Length,
			Value:  objectsMap[Merchant_Name_Alternate_Language].Value,
		},
	}

	return nil
}

func (emv *EMVCo) parseMerchantCityAlternateLanguage(objectsMap map[string]*objects.DataObject) error {
	// Merchant City Alternate Language is optional
	if objectsMap[Merchant_City_Alternate_Language] != nil {
		if objectsMap[Merchant_City_Alternate_Language].Length > 15 {
			return errors.New("invalid length of Merchant Name Alternate Language")
		}

		emv.mAlternateLanguage.MerchantCity = &objects.MerchantCityAlternateLanguage{
			DataObject: objects.DataObject{
				ID:     Merchant_City_Alternate_Language,
				Length: objectsMap[Merchant_City_Alternate_Language].Length,
				Value:  objectsMap[Merchant_City_Alternate_Language].Value,
			},
		}
	}

	return nil
}

func (emv *EMVCo) parseRFUAlternateLanguage(objectsMap map[string]*objects.DataObject) error {
	// Merchant City Alternate Language is optional
	rfus := make([]objects.MerchantRFUAlternateLanguage, 0)
	for i := 3; i <= 99; i++ {
		if objectsMap[strconv.Itoa(i)] != nil {
			dataObjects, err := emv.parseDataObjects(objectsMap[strconv.Itoa(i)].Value.(string), false)
			if err != nil {
				return err
			}

			rfus = append(rfus, objects.MerchantRFUAlternateLanguage{
				DataObject: objects.DataObject{
					ID:          strconv.Itoa(i),
					Length:      objectsMap[strconv.Itoa(i)].Length,
					Value:       objectsMap[strconv.Itoa(i)].Value,
					TemplateMap: dataObjects,
				},
			})

		}

	}

	emv.mAlternateLanguage.RFUs = rfus
	return nil
}
