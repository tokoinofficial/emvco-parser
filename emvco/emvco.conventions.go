package emvco

import (
	"emvco-parser/m/v2/objects"
	"errors"
)

const (
	Payload_Format_Indicator   = "00"
	Point_of_Initiation_Method = "01"
	CRC                        = "63"
)

func (emv *EMVCo) parseConventions(objectsMap map[string]*objects.DataObject) error {

	err := emv.parsePayloadFormatIndicator(objectsMap)
	if err != nil {
		return err
	}

	err = emv.parsePointOfInitiationMethod(objectsMap)
	if err != nil {
		return err
	}

	err = emv.parseCRC(objectsMap)
	if err != nil {
		return err
	}

	return nil
}

func (emv *EMVCo) parsePayloadFormatIndicator(objectsMap map[string]*objects.DataObject) error {
	// Payload_Format_Indicator is mandatory field
	if objectsMap[Payload_Format_Indicator] == nil {
		return errors.New("missing Payload Format Indicator")
	}

	if objectsMap[Payload_Format_Indicator].Length != 2 {
		return errors.New("invalid length of Payload Format Indicator")
	}

	emv.pfi = &objects.PayloadFormatIndicator{
		DataObject: objects.DataObject{
			ID:     Payload_Format_Indicator,
			Length: objectsMap[Payload_Format_Indicator].Length,
			Value:  objectsMap[Payload_Format_Indicator].Value,
		},
	}

	return nil
}

func (emv *EMVCo) parsePointOfInitiationMethod(objectsMap map[string]*objects.DataObject) error {
	// Point_of_Initiation_Method is optional
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

func (emv *EMVCo) parseCRC(objectsMap map[string]*objects.DataObject) error {
	// CRC is mandatory field
	if objectsMap[CRC] == nil {
		return errors.New("missing CRC")
	}

	if objectsMap[CRC].Length != 4 {
		return errors.New("invalid length of CRC")
	}

	emv.CRC = &objects.CRC{
		DataObject: objects.DataObject{
			ID:     CRC,
			Length: objectsMap[CRC].Length,
			Value:  objectsMap[CRC].Value,
		},
	}

	return nil
}
