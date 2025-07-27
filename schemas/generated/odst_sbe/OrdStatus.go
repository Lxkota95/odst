// Generated SBE (Simple Binary Encoding) message codec

package odst_sbe

import (
	"fmt"
	"io"
	"reflect"
)

type OrdStatusEnum byte
type OrdStatusValues struct {
	NEW       OrdStatusEnum
	CANCELED  OrdStatusEnum
	REJECTED  OrdStatusEnum
	NullValue OrdStatusEnum
}

var OrdStatus = OrdStatusValues{48, 52, 56, 0}

func (o OrdStatusEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrdStatusEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrdStatusEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrdStatus)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrdStatus, unknown enumeration value %d", o)
}

func (*OrdStatusEnum) EncodedLength() int64 {
	return 1
}

func (*OrdStatusEnum) NEWSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) NEWInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.NEWSinceVersion()
}

func (*OrdStatusEnum) NEWDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) CANCELEDSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) CANCELEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CANCELEDSinceVersion()
}

func (*OrdStatusEnum) CANCELEDDeprecated() uint16 {
	return 0
}

func (*OrdStatusEnum) REJECTEDSinceVersion() uint16 {
	return 0
}

func (o *OrdStatusEnum) REJECTEDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.REJECTEDSinceVersion()
}

func (*OrdStatusEnum) REJECTEDDeprecated() uint16 {
	return 0
}
