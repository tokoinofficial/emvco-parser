package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
)

const (
	Payload_Format_Indicator   = "00"
	Point_of_Initiation_Method = "01"
)

func (emv *EMVCo) parseConventions(objectsMap map[string]*objects.DataObject) error {
	// Payload_Format_Indicator is mandatory field
	if objectsMap[Payload_Format_Indicator] == nil {
		return errors.New("missing payload format indicator")
	}

	if objectsMap[Payload_Format_Indicator].Length != 2 {
		return errors.New("invalid length of payload format indicator")
	}

	emv.pfi = &objects.PayloadFormatIndicator{
		DataObject: objects.DataObject{
			ID:     Payload_Format_Indicator,
			Length: objectsMap[Payload_Format_Indicator].Length,
			Value:  objectsMap[Payload_Format_Indicator].Value,
		},
	}

	if objectsMap[Point_of_Initiation_Method] != nil {
		if objectsMap[Point_of_Initiation_Method].Length != 2 {
			return errors.New("invalid length of point of initiation method")
		}

		emv.poim = &objects.PointOfInitiationMethod{
			DataObject: objects.DataObject{
				ID:     Point_of_Initiation_Method,
				Length: objectsMap[Point_of_Initiation_Method].Length,
				Value:  objectsMap[Point_of_Initiation_Method].Value,
			},
		}
	}

	return nil
}
