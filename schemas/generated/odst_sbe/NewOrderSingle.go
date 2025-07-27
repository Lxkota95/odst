// Generated SBE (Simple Binary Encoding) message codec

package odst_sbe

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type NewOrderSingle struct {
	ClOrdID      [20]byte
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

func (n *NewOrderSingle) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := n.RangeCheck(n.SbeSchemaVersion(), n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, n.ClOrdID[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Account[:]); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Symbol[:]); err != nil {
		return err
	}
	if err := n.Side.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.OrderQty); err != nil {
		return err
	}
	if err := n.OrdType.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, n.Price); err != nil {
		return err
	}
	if err := n.TimeInForce.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, n.TransactTime); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, n.Text[:]); err != nil {
		return err
	}
	return nil
}

func (n *NewOrderSingle) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !n.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			n.ClOrdID[idx] = n.ClOrdIDNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.ClOrdID[:]); err != nil {
			return err
		}
	}
	if !n.AccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			n.Account[idx] = n.AccountNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Account[:]); err != nil {
			return err
		}
	}
	if !n.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 8; idx++ {
			n.Symbol[idx] = n.SymbolNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Symbol[:]); err != nil {
			return err
		}
	}
	if n.SideInActingVersion(actingVersion) {
		if err := n.Side.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.OrderQtyInActingVersion(actingVersion) {
		n.OrderQty = n.OrderQtyNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.OrderQty); err != nil {
			return err
		}
	}
	if n.OrdTypeInActingVersion(actingVersion) {
		if err := n.OrdType.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.PriceInActingVersion(actingVersion) {
		n.Price = n.PriceNullValue()
	} else {
		if err := _m.ReadInt64(_r, &n.Price); err != nil {
			return err
		}
	}
	if n.TimeInForceInActingVersion(actingVersion) {
		if err := n.TimeInForce.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !n.TransactTimeInActingVersion(actingVersion) {
		n.TransactTime = n.TransactTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &n.TransactTime); err != nil {
			return err
		}
	}
	if !n.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 128; idx++ {
			n.Text[idx] = n.TextNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, n.Text[:]); err != nil {
			return err
		}
	}
	if actingVersion > n.SbeSchemaVersion() && blockLength > n.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-n.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := n.RangeCheck(actingVersion, n.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (n *NewOrderSingle) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if n.ClOrdIDInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if n.ClOrdID[idx] < n.ClOrdIDMinValue() || n.ClOrdID[idx] > n.ClOrdIDMaxValue() {
				return fmt.Errorf("Range check failed on n.ClOrdID[%d] (%v < %v > %v)", idx, n.ClOrdIDMinValue(), n.ClOrdID[idx], n.ClOrdIDMaxValue())
			}
		}
	}
	for idx, ch := range n.ClOrdID {
		if ch > 127 {
			return fmt.Errorf("n.ClOrdID[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if n.AccountInActingVersion(actingVersion) {
		for idx := 0; idx < 12; idx++ {
			if n.Account[idx] < n.AccountMinValue() || n.Account[idx] > n.AccountMaxValue() {
				return fmt.Errorf("Range check failed on n.Account[%d] (%v < %v > %v)", idx, n.AccountMinValue(), n.Account[idx], n.AccountMaxValue())
			}
		}
	}
	for idx, ch := range n.Account {
		if ch > 127 {
			return fmt.Errorf("n.Account[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if n.SymbolInActingVersion(actingVersion) {
		for idx := 0; idx < 8; idx++ {
			if n.Symbol[idx] < n.SymbolMinValue() || n.Symbol[idx] > n.SymbolMaxValue() {
				return fmt.Errorf("Range check failed on n.Symbol[%d] (%v < %v > %v)", idx, n.SymbolMinValue(), n.Symbol[idx], n.SymbolMaxValue())
			}
		}
	}
	for idx, ch := range n.Symbol {
		if ch > 127 {
			return fmt.Errorf("n.Symbol[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if err := n.Side.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.OrderQtyInActingVersion(actingVersion) {
		if n.OrderQty < n.OrderQtyMinValue() || n.OrderQty > n.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on n.OrderQty (%v < %v > %v)", n.OrderQtyMinValue(), n.OrderQty, n.OrderQtyMaxValue())
		}
	}
	if err := n.OrdType.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.PriceInActingVersion(actingVersion) {
		if n.Price < n.PriceMinValue() || n.Price > n.PriceMaxValue() {
			return fmt.Errorf("Range check failed on n.Price (%v < %v > %v)", n.PriceMinValue(), n.Price, n.PriceMaxValue())
		}
	}
	if err := n.TimeInForce.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if n.TransactTimeInActingVersion(actingVersion) {
		if n.TransactTime < n.TransactTimeMinValue() || n.TransactTime > n.TransactTimeMaxValue() {
			return fmt.Errorf("Range check failed on n.TransactTime (%v < %v > %v)", n.TransactTimeMinValue(), n.TransactTime, n.TransactTimeMaxValue())
		}
	}
	if n.TextInActingVersion(actingVersion) {
		for idx := 0; idx < 128; idx++ {
			if n.Text[idx] < n.TextMinValue() || n.Text[idx] > n.TextMaxValue() {
				return fmt.Errorf("Range check failed on n.Text[%d] (%v < %v > %v)", idx, n.TextMinValue(), n.Text[idx], n.TextMaxValue())
			}
		}
	}
	for idx, ch := range n.Text {
		if ch > 127 {
			return fmt.Errorf("n.Text[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	return nil
}

func NewOrderSingleInit(n *NewOrderSingle) {
	return
}

func (*NewOrderSingle) SbeBlockLength() (blockLength uint16) {
	return 195
}

func (*NewOrderSingle) SbeTemplateId() (templateId uint16) {
	return 100
}

func (*NewOrderSingle) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*NewOrderSingle) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*NewOrderSingle) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*NewOrderSingle) SbeSemanticVersion() (semanticVersion string) {
	return "1.0"
}

func (*NewOrderSingle) ClOrdIDId() uint16 {
	return 11
}

func (*NewOrderSingle) ClOrdIDSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) ClOrdIDInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.ClOrdIDSinceVersion()
}

func (*NewOrderSingle) ClOrdIDDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) ClOrdIDMetaAttribute(meta int) string {
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

func (*NewOrderSingle) ClOrdIDMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle) ClOrdIDMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle) ClOrdIDNullValue() byte {
	return 0
}

func (n *NewOrderSingle) ClOrdIDCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderSingle) AccountId() uint16 {
	return 1
}

func (*NewOrderSingle) AccountSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) AccountInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.AccountSinceVersion()
}

func (*NewOrderSingle) AccountDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) AccountMetaAttribute(meta int) string {
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

func (*NewOrderSingle) AccountMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle) AccountMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle) AccountNullValue() byte {
	return 0
}

func (n *NewOrderSingle) AccountCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderSingle) SymbolId() uint16 {
	return 55
}

func (*NewOrderSingle) SymbolSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SymbolInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SymbolSinceVersion()
}

func (*NewOrderSingle) SymbolDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SymbolMetaAttribute(meta int) string {
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

func (*NewOrderSingle) SymbolMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle) SymbolMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle) SymbolNullValue() byte {
	return 0
}

func (n *NewOrderSingle) SymbolCharacterEncoding() string {
	return "US-ASCII"
}

func (*NewOrderSingle) SideId() uint16 {
	return 54
}

func (*NewOrderSingle) SideSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) SideInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.SideSinceVersion()
}

func (*NewOrderSingle) SideDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) SideMetaAttribute(meta int) string {
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

func (*NewOrderSingle) OrderQtyId() uint16 {
	return 38
}

func (*NewOrderSingle) OrderQtySinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrderQtySinceVersion()
}

func (*NewOrderSingle) OrderQtyDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) OrderQtyMetaAttribute(meta int) string {
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

func (*NewOrderSingle) OrderQtyMinValue() uint64 {
	return 0
}

func (*NewOrderSingle) OrderQtyMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderSingle) OrderQtyNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderSingle) OrdTypeId() uint16 {
	return 40
}

func (*NewOrderSingle) OrdTypeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) OrdTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.OrdTypeSinceVersion()
}

func (*NewOrderSingle) OrdTypeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) OrdTypeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) PriceId() uint16 {
	return 44
}

func (*NewOrderSingle) PriceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.PriceSinceVersion()
}

func (*NewOrderSingle) PriceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) PriceMetaAttribute(meta int) string {
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

func (*NewOrderSingle) PriceMinValue() int64 {
	return math.MinInt64 + 1
}

func (*NewOrderSingle) PriceMaxValue() int64 {
	return math.MaxInt64
}

func (*NewOrderSingle) PriceNullValue() int64 {
	return math.MinInt64
}

func (*NewOrderSingle) TimeInForceId() uint16 {
	return 59
}

func (*NewOrderSingle) TimeInForceSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TimeInForceSinceVersion()
}

func (*NewOrderSingle) TimeInForceDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TimeInForceMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TransactTimeId() uint16 {
	return 60
}

func (*NewOrderSingle) TransactTimeSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TransactTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TransactTimeSinceVersion()
}

func (*NewOrderSingle) TransactTimeDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TransactTimeMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TransactTimeMinValue() uint64 {
	return 0
}

func (*NewOrderSingle) TransactTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*NewOrderSingle) TransactTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*NewOrderSingle) TextId() uint16 {
	return 58
}

func (*NewOrderSingle) TextSinceVersion() uint16 {
	return 0
}

func (n *NewOrderSingle) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= n.TextSinceVersion()
}

func (*NewOrderSingle) TextDeprecated() uint16 {
	return 0
}

func (*NewOrderSingle) TextMetaAttribute(meta int) string {
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

func (*NewOrderSingle) TextMinValue() byte {
	return byte(32)
}

func (*NewOrderSingle) TextMaxValue() byte {
	return byte(126)
}

func (*NewOrderSingle) TextNullValue() byte {
	return 0
}

func (n *NewOrderSingle) TextCharacterEncoding() string {
	return "US-ASCII"
}
