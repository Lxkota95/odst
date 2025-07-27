// Generated SBE (Simple Binary Encoding) message codec

package odst_sbe

import (
	"fmt"
	"io"
	"reflect"
)

type OrdTypeEnum byte
type OrdTypeValues struct {
	MARKET    OrdTypeEnum
	LIMIT     OrdTypeEnum
	NullValue OrdTypeEnum
}

var OrdType = OrdTypeValues{49, 50, 0}

func (o OrdTypeEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(o)); err != nil {
		return err
	}
	return nil
}

func (o *OrdTypeEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(o)); err != nil {
		return err
	}
	return nil
}

func (o OrdTypeEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(OrdType)
	for idx := 0; idx < value.NumField(); idx++ {
		if o == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on OrdType, unknown enumeration value %d", o)
}

func (*OrdTypeEnum) EncodedLength() int64 {
	return 1
}

func (*OrdTypeEnum) MARKETSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) MARKETInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MARKETSinceVersion()
}

func (*OrdTypeEnum) MARKETDeprecated() uint16 {
	return 0
}

func (*OrdTypeEnum) LIMITSinceVersion() uint16 {
	return 0
}

func (o *OrdTypeEnum) LIMITInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LIMITSinceVersion()
}

func (*OrdTypeEnum) LIMITDeprecated() uint16 {
	return 0
}
