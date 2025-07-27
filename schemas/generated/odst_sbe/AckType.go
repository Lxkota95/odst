// Generated SBE (Simple Binary Encoding) message codec

package odst_sbe

import (
	"fmt"
	"io"
	"reflect"
)

type AckTypeEnum byte
type AckTypeValues struct {
	ORDER_RECEIVED  AckTypeEnum
	ORDER_ACCEPTED  AckTypeEnum
	ORDER_REJECTED  AckTypeEnum
	CANCEL_ACCEPTED AckTypeEnum
	CANCEL_REJECTED AckTypeEnum
	NullValue       AckTypeEnum
}

var AckType = AckTypeValues{48, 49, 56, 52, 57, 0}

func (a AckTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(a)); err != nil {
		return err
	}
	return nil
}

func (a *AckTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(a)); err != nil {
		return err
	}
	return nil
}

func (a AckTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(AckType)
	for idx := 0; idx < value.NumField(); idx++ {
		if a == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on AckType, unknown enumeration value %d", a)
}

func (*AckTypeEnum) EncodedLength() int64 {
	return 1
}

func (*AckTypeEnum) ORDER_RECEIVEDSinceVersion() uint16 {
	return 0
}

func (a *AckTypeEnum) ORDER_RECEIVEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.ORDER_RECEIVEDSinceVersion()
}

func (*AckTypeEnum) ORDER_RECEIVEDDeprecated() uint16 {
	return 0
}

func (*AckTypeEnum) ORDER_ACCEPTEDSinceVersion() uint16 {
	return 0
}

func (a *AckTypeEnum) ORDER_ACCEPTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.ORDER_ACCEPTEDSinceVersion()
}

func (*AckTypeEnum) ORDER_ACCEPTEDDeprecated() uint16 {
	return 0
}

func (*AckTypeEnum) ORDER_REJECTEDSinceVersion() uint16 {
	return 0
}

func (a *AckTypeEnum) ORDER_REJECTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.ORDER_REJECTEDSinceVersion()
}

func (*AckTypeEnum) ORDER_REJECTEDDeprecated() uint16 {
	return 0
}

func (*AckTypeEnum) CANCEL_ACCEPTEDSinceVersion() uint16 {
	return 0
}

func (a *AckTypeEnum) CANCEL_ACCEPTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.CANCEL_ACCEPTEDSinceVersion()
}

func (*AckTypeEnum) CANCEL_ACCEPTEDDeprecated() uint16 {
	return 0
}

func (*AckTypeEnum) CANCEL_REJECTEDSinceVersion() uint16 {
	return 0
}

func (a *AckTypeEnum) CANCEL_REJECTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.CANCEL_REJECTEDSinceVersion()
}

func (*AckTypeEnum) CANCEL_REJECTEDDeprecated() uint16 {
	return 0
}
