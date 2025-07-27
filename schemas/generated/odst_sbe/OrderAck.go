// Generated SBE (Simple Binary Encoding) message codec

package odst_sbe

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type OrderAck struct {
	OrderID      [16]byte
	ClOrdID      [20]byte
	AckType      AckTypeEnum
	OrdStatus    OrdStatusEnum
	Account      [12]byte
	Symbol       [8]byte
	Side         SideEnum
	OrderQty     uint64
	OrdType      OrdTypeEnum
	Price        int64
	TimeInForce  TimeInForceEnum
	TransactTime uint64
	Text         [128]byte
}

func (o *OrderAck) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, o.OrderID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.ClOrdID[:]); err != nil {
		return err
	}
	if err := o.AckType.Encode(_m, _w); err != nil {
		return err
	}
	if err := o.OrdStatus.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Account[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Symbol[:]); err != nil {
		return err
	}
	if err := o.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.OrderQty); err != nil {
		return err
	}
	if err := o.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, o.Price); err != nil {
		return err
	}
	if err := o.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Text[:]); err != nil {
		return err
	}
	return nil
}

func (o *OrderAck) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.OrderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 16; idx++ {
			o.OrderID[idx] = o.OrderIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.OrderID[:]); err != nil {
			return err
		}
	}
	if !o.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			o.ClOrdID[idx] = o.ClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.ClOrdID[:]); err != nil {
			return err
		}
	}
	if o.AckTypeInActingVersion(actingVersion) {
		if err := o.AckType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if o.OrdStatusInActingVersion(actingVersion) {
		if err := o.OrdStatus.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.AccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			o.Account[idx] = o.AccountNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Account[:]); err != nil {
			return err
		}
	}
	if !o.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 8; idx++ {
			o.Symbol[idx] = o.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Symbol[:]); err != nil {
			return err
		}
	}
	if o.SideInActingVersion(actingVersion) {
		if err := o.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.OrderQtyInActingVersion(actingVersion) {
		o.OrderQty = o.OrderQtyNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.OrderQty); err != nil {
			return err
		}
	}
	if o.OrdTypeInActingVersion(actingVersion) {
		if err := o.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.PriceInActingVersion(actingVersion) {
		o.Price = o.PriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.Price); err != nil {
			return err
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if err := o.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !o.TransactTimeInActingVersion(actingVersion) {
		o.TransactTime = o.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.TransactTime); err != nil {
			return err
		}
	}
	if !o.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 128; idx++ {
			o.Text[idx] = o.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, o.Text[:]); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *OrderAck) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderIDInActingVersion(actingVersion) {
		for idx := 0; idx < 16; idx++ {
			if o.OrderID[idx] < o.OrderIDMinValue() || o.OrderID[idx] > o.OrderIDMaxValue() {
				return fmt.Errorf("Range check failed on o.OrderID[%d] (%v < %v > %v)", idx, o.OrderIDMinValue(), o.OrderID[idx], o.OrderIDMaxValue())
			}
		}
	}
	for idx, ch := range o.OrderID {
		if ch > 127 {
			return fmt.Errorf("o.OrderID[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if o.ClOrdID[idx] < o.ClOrdIDMinValue() || o.ClOrdID[idx] > o.ClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on o.ClOrdID[%d] (%v < %v > %v)", idx, o.ClOrdIDMinValue(), o.ClOrdID[idx], o.ClOrdIDMaxValue())
			}
		}
	}
	for idx, ch := range o.ClOrdID {
		if ch > 127 {
			return fmt.Errorf("o.ClOrdID[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if err := o.AckType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if err := o.OrdStatus.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.AccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			if o.Account[idx] < o.AccountMinValue() || o.Account[idx] > o.AccountMaxValue() {
				return fmt.Errorf("Range check failed on o.Account[%d] (%v < %v > %v)", idx, o.AccountMinValue(), o.Account[idx], o.AccountMaxValue())
			}
		}
	}
	for idx, ch := range o.Account {
		if ch > 127 {
			return fmt.Errorf("o.Account[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if o.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 8; idx++ {
			if o.Symbol[idx] < o.SymbolMinValue() || o.Symbol[idx] > o.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on o.Symbol[%d] (%v < %v > %v)", idx, o.SymbolMinValue(), o.Symbol[idx], o.SymbolMaxValue())
			}
		}
	}
	for idx, ch := range o.Symbol {
		if ch > 127 {
			return fmt.Errorf("o.Symbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if err := o.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.OrderQtyInActingVersion(actingVersion) {
		if o.OrderQty < o.OrderQtyMinValue() || o.OrderQty > o.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderQty (%v < %v > %v)", o.OrderQtyMinValue(), o.OrderQty, o.OrderQtyMaxValue())
		}
	}
	if err := o.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.PriceInActingVersion(actingVersion) {
		if o.Price < o.PriceMinValue() || o.Price > o.PriceMaxValue() {
			return fmt.Errorf("Range check failed on o.Price (%v < %v > %v)", o.PriceMinValue(), o.Price, o.PriceMaxValue())
		}
	}
	if err := o.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if o.TransactTimeInActingVersion(actingVersion) {
		if o.TransactTime < o.TransactTimeMinValue() || o.TransactTime > o.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.TransactTime (%v < %v > %v)", o.TransactTimeMinValue(), o.TransactTime, o.TransactTimeMaxValue())
		}
	}
	if o.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 128; idx++ {
			if o.Text[idx] < o.TextMinValue() || o.Text[idx] > o.TextMaxValue() {
				return fmt.Errorf("Range check failed on o.Text[%d] (%v < %v > %v)", idx, o.TextMinValue(), o.Text[idx], o.TextMaxValue())
			}
		}
	}
	for idx, ch := range o.Text {
		if ch > 127 {
			return fmt.Errorf("o.Text[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func OrderAckInit(o *OrderAck) {
	return
}

func (*OrderAck) SbeBlockLength() (blockLength uint16) {
	return 213
}

func (*OrderAck) SbeTemplateId() (templateId uint16) {
	return 101
}

func (*OrderAck) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderAck) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OrderAck) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OrderAck) SbeSemanticVersion() (semanticVersion string) {
	return "1.0"
}

func (*OrderAck) OrderIDId() uint16 {
	return 37
}

func (*OrderAck) OrderIDSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) OrderIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIDSinceVersion()
}

func (*OrderAck) OrderIDDeprecated() uint16 {
	return 0
}

func (*OrderAck) OrderIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) OrderIDMinValue() byte {
	return byte(32)
}

func (*OrderAck) OrderIDMaxValue() byte {
	return byte(126)
}

func (*OrderAck) OrderIDNullValue() byte {
	return 0
}

func (o *OrderAck) OrderIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderAck) ClOrdIDId() uint16 {
	return 11
}

func (*OrderAck) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ClOrdIDSinceVersion()
}

func (*OrderAck) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*OrderAck) ClOrdIDMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*OrderAck) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*OrderAck) ClOrdIDNullValue() byte {
	return 0
}

func (o *OrderAck) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderAck) AckTypeId() uint16 {
	return 150
}

func (*OrderAck) AckTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) AckTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AckTypeSinceVersion()
}

func (*OrderAck) AckTypeDeprecated() uint16 {
	return 0
}

func (*OrderAck) AckTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) OrdStatusId() uint16 {
	return 39
}

func (*OrderAck) OrdStatusSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) OrdStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdStatusSinceVersion()
}

func (*OrderAck) OrdStatusDeprecated() uint16 {
	return 0
}

func (*OrderAck) OrdStatusMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) AccountId() uint16 {
	return 1
}

func (*OrderAck) AccountSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) AccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AccountSinceVersion()
}

func (*OrderAck) AccountDeprecated() uint16 {
	return 0
}

func (*OrderAck) AccountMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) AccountMinValue() byte {
	return byte(32)
}

func (*OrderAck) AccountMaxValue() byte {
	return byte(126)
}

func (*OrderAck) AccountNullValue() byte {
	return 0
}

func (o *OrderAck) AccountCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderAck) SymbolId() uint16 {
	return 55
}

func (*OrderAck) SymbolSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SymbolSinceVersion()
}

func (*OrderAck) SymbolDeprecated() uint16 {
	return 0
}

func (*OrderAck) SymbolMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) SymbolMinValue() byte {
	return byte(32)
}

func (*OrderAck) SymbolMaxValue() byte {
	return byte(126)
}

func (*OrderAck) SymbolNullValue() byte {
	return 0
}

func (o *OrderAck) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*OrderAck) SideId() uint16 {
	return 54
}

func (*OrderAck) SideSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.SideSinceVersion()
}

func (*OrderAck) SideDeprecated() uint16 {
	return 0
}

func (*OrderAck) SideMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) OrderQtyId() uint16 {
	return 38
}

func (*OrderAck) OrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderAck) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderQtySinceVersion()
}

func (*OrderAck) OrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderAck) OrderQtyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) OrderQtyMinValue() uint64 {
	return 0
}

func (*OrderAck) OrderQtyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderAck) OrderQtyNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderAck) OrdTypeId() uint16 {
	return 40
}

func (*OrderAck) OrdTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrdTypeSinceVersion()
}

func (*OrderAck) OrdTypeDeprecated() uint16 {
	return 0
}

func (*OrderAck) OrdTypeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) PriceId() uint16 {
	return 44
}

func (*OrderAck) PriceSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceSinceVersion()
}

func (*OrderAck) PriceDeprecated() uint16 {
	return 0
}

func (*OrderAck) PriceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderAck) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderAck) PriceNullValue() int64 {
	return math.MinInt64
}

func (*OrderAck) TimeInForceId() uint16 {
	return 59
}

func (*OrderAck) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderAck) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderAck) TimeInForceMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) TransactTimeId() uint16 {
	return 60
}

func (*OrderAck) TransactTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TransactTimeSinceVersion()
}

func (*OrderAck) TransactTimeDeprecated() uint16 {
	return 0
}

func (*OrderAck) TransactTimeMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) TransactTimeMinValue() uint64 {
	return 0
}

func (*OrderAck) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderAck) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderAck) TextId() uint16 {
	return 58
}

func (*OrderAck) TextSinceVersion() uint16 {
	return 0
}

func (o *OrderAck) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TextSinceVersion()
}

func (*OrderAck) TextDeprecated() uint16 {
	return 0
}

func (*OrderAck) TextMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*OrderAck) TextMinValue() byte {
	return byte(32)
}

func (*OrderAck) TextMaxValue() byte {
	return byte(126)
}

func (*OrderAck) TextNullValue() byte {
	return 0
}

func (o *OrderAck) TextCharacterEncoding() string {
	return "US-ASCII"
}
