package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
	"fmt"
	"strconv"
)

type EMVCo struct {
	pfi                                *objects.PayloadFormatIndicator
	poim                               *objects.PointOfInitiationMethod
	mAccounts                          []objects.MerchantAccount
	mCategoryCode                      *objects.MerchantCategoryCode
	mCountryCode                       *objects.MerchantCountryCode
	mName                              *objects.MerchantName
	mCity                              *objects.MerchantCity
	mPostalCode                        *objects.MerchantPostalCode
	mAlternateLanguage                 *objects.MerchantAlternateLanguage
	tnxCurrency                        *objects.TransactionCurrency
	tnxAmount                          *objects.TransactionAmount
	tnxTip                             *objects.TransactionTip
	tnxValueOfConvenienceFeeFixed      *objects.TransactionValueOfConvenienceFeeFixed
	tnxValueOfConvenienceFeePercentage *objects.TransactionValueOfConvenienceFeePercentage
	dataFieldTemplate                  *objects.AdditionalDataFieldTemplate
	CRC                                *objects.CRC
}

func NewEMVCo() *EMVCo {
	return &EMVCo{}
}

func (emv *EMVCo) ToString() string {
	return fmt.Sprintf(`
		PayloadFormatIndicator: %v | PointOfInitiationMethod: %v
		MerchantAccount: %v
		MerchantCategoryCode: %v
		MerchantCountryCode: %v
		MerchantName: %v
		MerchantCity: %v
		MerchantPostalCode: %v
		MerchantAlternateLanguage: %v
		TransactionCurrency: %v
		TransactionAmount: %v
		TransactionTip: %v
		TransactionValueOfConvenienceFeeFixed: %v
		TransactionValueOfConvenienceFeePercentage: %v
		AdditionalDataFieldTemplate: %v
		CRC: %v
	`,
		emv.pfi, emv.poim,
		emv.mAccounts,
		emv.mCategoryCode,
		emv.mCountryCode,
		emv.mName,
		emv.mCity,
		emv.mPostalCode,
		emv.mAlternateLanguage,
		emv.tnxCurrency,
		emv.tnxAmount,
		emv.tnxTip,
		emv.tnxValueOfConvenienceFeeFixed,
		emv.tnxValueOfConvenienceFeePercentage,
		emv.dataFieldTemplate,
		emv.CRC,
	)
}

func (emv *EMVCo) Parse(str string) error {

	objectsMap, err := emv.parseDataObjects(str, true)
	if err != nil {
		return err
	}

	err = emv.parseConventions(objectsMap)
	if err != nil {
		return err
	}

	err = emv.parseMerchant(objectsMap)
	if err != nil {
		return err
	}

	err = emv.parseAdditionalMerchant(objectsMap)
	if err != nil {
		return err
	}

	err = emv.parseTransaction(objectsMap)
	if err != nil {
		return err
	}

	err = emv.parseAdditionalDataFieldTemplate(objectsMap)
	if err != nil {
		return err
	}

	err = emv.parseMerchantAlternateLanguage(objectsMap)
	if err != nil {
		return err
	}

	return nil
}

func (emv *EMVCo) parseDataObjects(str string, isRoot bool) (map[string]*objects.DataObject, error) {
	var (
		length  int
		firstID int
		lastID  int
	)

	firstID = 1
	objectsMap := map[string]*objects.DataObject{}
	for i := 0; i < len(str); i = i + 4 + length {
		idStr := str[i : i+2]
		ID, err := emv.validateEMVCoID(idStr)
		if err != nil {
			return nil, err
		}

		if firstID > ID {
			firstID = ID
		}

		if lastID < ID {
			lastID = ID
		}

		length, _ = strconv.Atoi(str[i+2 : i+4])
		if i+4+length > len(str) {
			return nil, errors.New("invalid qris value")
		}
		value := str[i+4 : i+4+length]
		objectsMap[idStr] = &objects.DataObject{
			ID:     idStr,
			Length: length,
			Value:  value,
		}
	}

	if isRoot {
		if firstID != 0 || lastID != 63 {
			return nil, errors.New("invalid qris")
		}
	}

	return objectsMap, nil
}

func (emv *EMVCo) validateEMVCoID(ID string) (int, error) {
	qrisID, error := strconv.Atoi(ID)
	if error != nil {
		return -1, error
	}

	if qrisID < 0 || qrisID > 99 {
		return qrisID, errors.New("invalid qris id")
	}
	return qrisID, nil
}

// TODO: in order to get which information, please create a public function here
