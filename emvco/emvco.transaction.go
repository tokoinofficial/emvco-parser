package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
)

const (
	Transaction_Currency = "53"
	Transaction_Amount   = "54"
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

	return nil
}
