// Generated SBE (Simple Binary Encoding) message codec

package odst_sbe

import (
	"fmt"
	"io"
	"reflect"
)

type SideEnum byte
type SideValues struct {
	BUY       SideEnum
	SELL      SideEnum
	NullValue SideEnum
}

var Side = SideValues{49, 50, 0}

func (s SideEnum) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, byte(s)); err != nil {
		return err
	}
	return nil
}

func (s *SideEnum) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16) error {
	if err := _m.ReadUint8(_r, (*byte)(s)); err != nil {
		return err
	}
	return nil
}

func (s SideEnum) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if actingVersion > schemaVersion {
		return nil
	}
	value := reflect.ValueOf(Side)
	for idx := 0; idx < value.NumField(); idx++ {
		if s == value.Field(idx).Interface() {
			return nil
		}
	}
	return fmt.Errorf("Range check failed on Side, unknown enumeration value %d", s)
}

func (*SideEnum) EncodedLength() int64 {
	return 1
}

func (*SideEnum) BUYSinceVersion() uint16 {
	return 0
}

func (s *SideEnum) BUYInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BUYSinceVersion()
}

func (*SideEnum) BUYDeprecated() uint16 {
	return 0
}

func (*SideEnum) SELLSinceVersion() uint16 {
	return 0
}

func (s *SideEnum) SELLInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SELLSinceVersion()
}

func (*SideEnum) SELLDeprecated() uint16 {
	return 0
}
