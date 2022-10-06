package emvco

import (
	"emvco-parser/m/v2/objects"
	"strconv"
)

// “02”-“03”: Reserved for Visa
// “04”-“05”: Reserved for Mastercard
// “06”-“08”: Reserved by EMVCo
// “09”-“10”: Reserved for Discover
// “11”-“12”: Reserved for Amex
// “13”-“14”: Reserved for JCB
// “15”-“16”: Reserved for UnionPay
// “17”-“25”: Reserved by EMVCo
// “26”-“51”: Templates reserved for any payment networks fulfilling the requirements in section 4.7.11 Merchant Account Information Template (IDs "26" to "51"). For content of this template, please refer to
// Table 4.2.
func (emv *EMVCo) parseMerchant(objectsMap map[string]*objects.DataObject) error {
	// Merchant Account Information("02"- "51"). This is mandatory field and at least one Merchant Account Information data object shall be present.
	merchants := make([]objects.MerchantAccount, 0)
	for i := 26; i <= 51; i++ {
		if objectsMap[strconv.Itoa(i)] != nil {
			// Data Object ID Allocation in Merchant Account Information Template (IDs "26" to "51")
			// "00": Globally Unique Identifier (M)
			// 	An identifier that sets the context of the data that follows.
			// 	The value is one of the following:
			// 		• an Application Identifier (AID);
			// 		• a [UUID] without the hyphen (-) separators;
			// 		• a reverse domain name.
			// "01"- "99": Payment network specific (O)
			// 	Association of data objects to IDs and type of data object is specific to the Globally Unique Identifier.
			dataObjects, err := emv.parseDataObjects(objectsMap[strconv.Itoa(i)].Value.(string), false)
			if err != nil {
				return err
			}

			merchants = append(merchants, objects.MerchantAccount{
				DataObject: objects.DataObject{
					ID:          strconv.Itoa(i),
					Length:      objectsMap[strconv.Itoa(i)].Length,
					Value:       objectsMap[strconv.Itoa(i)].Value,
					TemplateMap: dataObjects,
				},
			})

		}

	}

	emv.mc = merchants
	return nil
}
