package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
)

const (
	Merchant_Category_Code = "52"
)

func (emv *EMVCo) parseAdditionalMerchant(objectsMap map[string]*objects.DataObject) error {

	// Merchant_Category_Code is mandatory field
	if objectsMap[Merchant_Category_Code] == nil {
		return errors.New("missing merchant category code")
	}

	if objectsMap[Merchant_Category_Code].Length != 4 {
		return errors.New("invalid length of merchant category code")
	}

	emv.mcc = &objects.MerchantCategoryCode{
		DataObject: objects.DataObject{
			ID:     Merchant_Category_Code,
			Length: objectsMap[Merchant_Category_Code].Length,
			Value:  objectsMap[Merchant_Category_Code].Value,
		},
	}

	return nil
}
