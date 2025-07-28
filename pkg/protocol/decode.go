package protocol

import (
	"time"

	"github.com/lxkota95/odst/schemas/generated/odst_sbe"
)

func encodeNewOrder() []byte {
	// Create buffer
	buffer := make([]byte, 1024)
	offset := 0

	// Encode SBE header
	header := odst_sbe.MessageHeader{}
	header.WrapForEncode(buffer, uint64(offset), uint64(len(buffer)))
	header.SetBlockLength(195) // NewOrderSingle size
	header.SetTemplateId(100)  // NewOrderSingle template
	header.SetSchemaId(1)
	header.SetVersion(0)
	offset += int(header.Size())

	// Encode message body
	order := odst_sbe.NewOrderSingle{}
	order.WrapForEncode(buffer, uint64(offset), uint64(len(buffer)))
	order.SetClOrdID("CLIENT001")
	order.SetAccount("ACCT123")
	order.SetSymbol("AAPL")
	order.SetSide(odst_sbe.Side.BUY)
	order.SetOrderQty(100)
	order.SetOrdType(odst_sbe.OrdType.LIMIT)
	order.SetPrice(15000) // $150.00 in ticks
	order.SetTimeInForce(odst_sbe.TimeInForce.DAY)
	order.SetTransactTime(uint64(time.Now().UnixNano()))
	order.SetText("Test limit order")

	totalSize := int(header.Size()) + int(order.Size())
	return buffer[:totalSize]
}
