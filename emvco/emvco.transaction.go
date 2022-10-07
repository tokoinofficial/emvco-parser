package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
)

const (
	Transaction_Currency                = "53"
	Transaction_Amount                  = "54"
	Transaction_Tip                     = "55"
	Value_Of_Convenience_Fee_Fixed      = "56"
	Value_Of_Convenience_Fee_Percentage = "57"
)

func (emv *EMVCo) parseTransaction(objectsMap map[string]*objects.DataObject) error {
	// Transaction_Currency is mandatory field
	if objectsMap[Transaction_Currency] == nil {
		return errors.New("missing transaction currency")
	}

	if objectsMap[Transaction_Currency].Length != 3 {
		return errors.New("invalid length of transaction currency")
	}

	emv.tnxCurrency = &objects.TransactionCurrency{
		DataObject: objects.DataObject{
			ID:     Transaction_Currency,
			Length: objectsMap[Transaction_Currency].Length,
			Value:  objectsMap[Transaction_Currency].Value,
		},
	}

	// Transaction_Amount is custom
	if objectsMap[Transaction_Amount] != nil {
		if objectsMap[Transaction_Amount].Length > 13 {
			return errors.New("invalid length of transaction amount")
		}

		emv.tnxAmount = &objects.TransactionAmount{
			DataObject: objects.DataObject{
				ID:     Transaction_Amount,
				Length: objectsMap[Transaction_Amount].Length,
				Value:  objectsMap[Transaction_Amount].Value,
			},
		}
	}

	// Transaction_Tip is optional
	if objectsMap[Transaction_Tip] != nil {
		if objectsMap[Transaction_Tip].Length != 2 {
			return errors.New("invalid length of transaction tip")
		}

		emv.tnxTip = &objects.TransactionTip{
			DataObject: objects.DataObject{
				ID:     Transaction_Tip,
				Length: objectsMap[Transaction_Tip].Length,
				Value:  objectsMap[Transaction_Tip].Value,
			},
		}
	}

	// Transaction_Tip is optional
	if objectsMap[Transaction_Tip] != nil {
		if objectsMap[Transaction_Tip].Length != 2 {
			return errors.New("invalid length of transaction tip")
		}

		emv.tnxTip = &objects.TransactionTip{
			DataObject: objects.DataObject{
				ID:     Transaction_Tip,
				Length: objectsMap[Transaction_Tip].Length,
				Value:  objectsMap[Transaction_Tip].Value,
			},
		}
	}

	// Value_Of_Convenience_Fee_Fixed is custom
	if objectsMap[Value_Of_Convenience_Fee_Fixed] != nil {
		if objectsMap[Value_Of_Convenience_Fee_Fixed].Length > 13 {
			return errors.New("invalid length of value of convenience fee fixed")
		}

		emv.tnxValueOfConvenienceFeeFixed = &objects.TransactionValueOfConvenienceFeeFixed{
			DataObject: objects.DataObject{
				ID:     Value_Of_Convenience_Fee_Fixed,
				Length: objectsMap[Value_Of_Convenience_Fee_Fixed].Length,
				Value:  objectsMap[Value_Of_Convenience_Fee_Fixed].Value,
			},
		}
	}

	// Value_Of_Convenience_Fee_Percentage is custom
	if objectsMap[Value_Of_Convenience_Fee_Percentage] != nil {
		if objectsMap[Value_Of_Convenience_Fee_Percentage].Length > 5 {
			return errors.New("invalid length of value of convenience fee percentage")
		}

		emv.tnxValueOfConvenienceFeePercentage = &objects.TransactionValueOfConvenienceFeePercentage{
			DataObject: objects.DataObject{
				ID:     Value_Of_Convenience_Fee_Percentage,
				Length: objectsMap[Value_Of_Convenience_Fee_Percentage].Length,
				Value:  objectsMap[Value_Of_Convenience_Fee_Percentage].Value,
			},
		}
	}

	return nil
}
